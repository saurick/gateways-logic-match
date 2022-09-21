package api

import (
	"context"
	"fmt"
	clientPb "gateways-logic-match/services-gen-code"
)

//	加入匹配
func JoinMatch(userId int32) error {
	conn, client := ConnectToServicesMatch()
	defer conn.Close()
	//	加入匹配
	_, err := client.Join(context.Background(), &clientPb.C2SJoin{UserId: userId})
	if err != nil {
		return fmt.Errorf("player join match failed: %v", err)
	}
	return nil
}
