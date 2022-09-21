package api

import (
	"context"
	"fmt"
	clientPb "gateways-logic-match/services-gen-code"
)

//	退出匹配
func QuitMatch(userId int32) error {
	conn, client := ConnectToServicesMatch()
	defer conn.Close()
	//	退出匹配
	_, err := client.Quit(context.Background(), &clientPb.C2SQuit{UserId: userId})
	if err != nil {
		return fmt.Errorf("player quit match failed: %v", err)
	}
	return nil
}
