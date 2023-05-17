package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/iot-platform/helper"
)

var adminServiceAddr = "http://43.139.116.74:10002"

// 设备列表
func TestDeviceList(t *testing.T) {
	m := map[string]string{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWRlbnRpdHkiOiIxMjMiLCJuYW1lIjoiMSIsImV4cCI6MTY4NjgxMzcxMn0.NSFPaP4Ndataetlx7be0zUWZnqugX_T3gIMXBMc5d4A",
	}
	header, _ := json.Marshal(m)
	rep, err := helper.HttpGet(adminServiceAddr+"/device/list?page=1size=60&name=", header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

// 产品列表
func TestProductList(t *testing.T) {
	m := map[string]string{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWRlbnRpdHkiOiIxMjMiLCJuYW1lIjoiMSIsImV4cCI6MTY4NjgxMzcxMn0.NSFPaP4Ndataetlx7be0zUWZnqugX_T3gIMXBMc5d4A",
	}
	header, _ := json.Marshal(m)
	rep, err := helper.HttpGet(adminServiceAddr+"/product/list?page=1&size=60&name=22", header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}
