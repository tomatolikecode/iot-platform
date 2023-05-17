package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/iot-platform/define"
	"github.com/iot-platform/helper"
)

var userServiceAddr = "http://43.139.116.74:10000"

func TestUserLogin(t *testing.T) {
	m := define.M{
		"username": "1",
		"password": "2",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}
