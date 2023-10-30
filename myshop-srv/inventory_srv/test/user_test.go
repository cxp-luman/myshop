package test_test

import (
	"context"
	"fmt"
	"myshop/user_srvs/proto"
	"testing"

	"google.golang.org/grpc"
)

func TestCheckPassWord(t *testing.T) {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	uc := proto.NewUserClient(conn)
	cr, err := uc.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{Password: "15109787862", EncryptedPassword: "$pbkdf2-sha512$TkvQYHWMXERSpzDu$781be243e063c2460be1ed0be74728744ab47d2a782d6a13d6ad1150ba9a0635"})
	if err != nil {
		fmt.Print(err)
	}
	if cr.Success{
		t.Error(`password is right`)
	}
}