package api

import (
	clientPb "gateways-logic-match/services-gen-code"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EventMatchSuccessPayload struct {
	PlayerId     int32
	MatchId      int32 //	对局匹配id
	SandboxIp    string
	SandboxPort  int32
	SandboxToken string
}

//	连接到匹配服务器
func ConnectToServicesMatch() (*grpc.ClientConn, clientPb.MatchClient) {
	conn, err := grpc.Dial("localhost:30053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := clientPb.NewMatchClient(conn)
	return conn, client
}
