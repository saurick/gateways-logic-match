package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	log "github.com/sirupsen/logrus"

	"gateways-logic-match/api"
	serverPb "gateways-logic-match/gateway-gen-code"

	rdx "gateways-logic-match/pkg/db/rdx"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type MatchServer struct {
	serverPb.UnimplementedMatchServer
}

//	获取玩家id
func getUserId(token string, ctx context.Context) (*int32, error) {
	authRedisClient := rdx.AuthRedisClient()
	defer authRedisClient.Close()
	userIdString, err := authRedisClient.Get(ctx, token).Result()
	if err != nil {
		return nil, fmt.Errorf("get user id failed: %v", err)
	}
	userIdInt, err := strconv.Atoi(userIdString)
	if err != nil {
		return nil, fmt.Errorf("convert int32 to string failed:  %v", err)
	}
	id := int32(userIdInt)
	return &id, nil
}

//	检查token是否存在
func checkTokenExist(token string) bool {
	authRedisClient := rdx.AuthRedisClient()
	defer authRedisClient.Close()
	_, err := authRedisClient.Get(context.Background(), token).Result()
	if err == redis.Nil {
		return false
	}
	if err != redis.Nil && err != nil {
		log.Errorf("check exist failed: %v", err)
		return false
	}
	return true
}

//	玩家进入匹配
func (m *MatchServer) Match(C2SMatch *serverPb.C2SMatch, stream serverPb.Match_MatchServer) error {
	//	metadata校验
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Error(codes.PermissionDenied, "permission denied")
	}
	if md["token"] == nil {
		return status.Error(codes.PermissionDenied, "permission denied")
	}

	//	token鉴权
	tokenFromClient := md["token"][0]
	if !checkTokenExist(tokenFromClient) {
		return status.Error(codes.PermissionDenied, "permission denied")
	}

	//	获取玩家id
	userId, err := getUserId(tokenFromClient, stream.Context())
	if err != nil {
		return status.Error(codes.Canceled, "match failed")
	}

	//	加入匹配
	if err := api.JoinMatch(*userId); err != nil {
		return status.Error(codes.Canceled, "match failed")
	}

	eventBusRedisClient := rdx.EventBusRedisClient()
	defer eventBusRedisClient.Close()
	//	订阅玩家匹配事件
	pubsub := eventBusRedisClient.PSubscribe(context.Background(), "Match_*?_u"+fmt.Sprint(*userId))
	ch := pubsub.Channel()
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			log.WithError(err).Warn("defer close pubsub failed")
		}
	}(pubsub)

	for {
		select {
		case channel := <-ch:
			var msg serverPb.S2CMatch_Msg            //	grpc消息
			var payload api.EventMatchSuccessPayload //	事件payload
			err := json.Unmarshal([]byte(channel.Payload), &payload)
			if err != nil {
				fmt.Println("unmarshal payload failed:", err)
				return status.Error(codes.Canceled, "match failed")
			}

			switch channel.Channel {
			case "Match_QuitSuccess_u" + fmt.Sprint(*userId):
				msg = serverPb.S2CMatch_CancelSuccess
				stream.Send(&serverPb.S2CMatch{Msg: msg})
				return nil
			case "Match_MatchSuccess_m" + fmt.Sprint(payload.MatchId) + "_u" + fmt.Sprint(payload.PlayerId):
				serverInfo := &serverPb.ServerInfo{
					Ip:    payload.SandboxIp,
					Port:  payload.SandboxPort,
					Token: payload.SandboxToken,
				}
				msg = serverPb.S2CMatch_MatchSuccess
				stream.Send(&serverPb.S2CMatch{Msg: msg, ServerInfo: serverInfo})
				return nil
			}
		case <-stream.Context().Done():
			//	stream结束时移除玩家信息
			matchPoolRedisClient := rdx.MatchPoolRedisClient()
			defer matchPoolRedisClient.Close()
			err := matchPoolRedisClient.Del(context.Background(), "match#"+fmt.Sprint(*userId)).Err()
			if err != nil {
				log.Errorf("remove player from match queue failed: %v", err)
			}
			return nil
		}
	}
}

//	玩家取消匹配
func (m *MatchServer) CancelMatch(ctx context.Context, in *serverPb.C2SCancelMatch) (*serverPb.S2CCancelMatch, error) {
	//	metadata校验
	md, ok := metadata.FromIncomingContext(ctx)
	var tokenFromClient string
	if !ok {
		return &serverPb.S2CCancelMatch{}, nil
	}

	//	token鉴权
	tokenFromClient = md["token"][0]
	if !checkTokenExist(tokenFromClient) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	//	获取玩家id
	userId, err := getUserId(tokenFromClient, context.Background())
	if err != nil {
		return nil, status.Error(codes.Canceled, "cancel match failed")
	}

	//	退出匹配
	if err = api.QuitMatch(*userId); err != nil {
		return nil, status.Error(codes.Canceled, "cancel match failed")
	}
	return &serverPb.S2CCancelMatch{}, nil
}

//	启动网关匹配服务
func LaunchGatewaysLogicMatch(port *int) {
	//	监听端口
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//	注册grpc服务器
	grpcServer := grpc.NewServer()
	serverPb.RegisterMatchServer(grpcServer, &MatchServer{})
	log.Printf("server listening at %v", lis.Addr())

	//	启动grpc服务
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
