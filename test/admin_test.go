package test

import (
	"fmt"
	"testing"

	"github.com/iot-platform/helper"
)

var adminServiceAddr = "http://43.139.116.74:10002"

func TestDeviceList(t *testing.T) {
	rep, err := helper.HttpGet(adminServiceAddr + "/device/list?page=1&size=20&name=一")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}
