package test_test

import (
	"fmt"
	"testing"
)


func TestConn(t *testing.T) {
	// initialize.InitConfig()
	s := fmt.Sprintf(`Service == "%s"`, "user_srv")
	fmt.Println(s)
}