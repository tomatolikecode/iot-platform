package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/iot-platform/helper"
)

var adminServiceAddr = "http://43.139.116.74:10002"
var token = map[string]string{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWRlbnRpdHkiOiIxMjMiLCJuYW1lIjoiMSIsImV4cCI6MTY4NjgxMzcxMn0.NSFPaP4Ndataetlx7be0zUWZnqugX_T3gIMXBMc5d4A",
}

var header []byte

func init() {
	header, _ = json.Marshal(token)
}

/****************
	设备模块
****************/
// 设备列表
func TestDeviceList(t *testing.T) {
	rep, err := helper.HttpGet(adminServiceAddr+"/device/list?page=1size=60&name=", header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

// 设备创建
func TestDeviceCreate(t *testing.T) {
	deviceCreate := map[string]string{
		"name":             "测试设备2",
		"product_identity": "1",
	}
	data, _ := json.Marshal(deviceCreate)

	rep, err := helper.HttpPost(adminServiceAddr+"/device/create", data, header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

// 设备更新
func TestDeviceModify(t *testing.T) {
	deviceModity := map[string]string{
		"identity":         "23faa62c-c336-4320-ad5c-f5508a35faf6",
		"name":             "测试设备_更新_1",
		"product_identity": "1",
	}
	data, _ := json.Marshal(deviceModity)
	rep, err := helper.HttpPut(adminServiceAddr+"/device/modify", data, header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

// 设备删除
func TestDeviceDelete(t *testing.T) {
	deviceDelete := map[string]string{
		"identity": "7339b199-2754-490e-9667-5f7e59dc34fc",
	}
	data, _ := json.Marshal(deviceDelete)

	rep, err := helper.HttpDelete(adminServiceAddr+"/device/delete", data, header...)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(rep))
}

/****************
	产品模块
****************/
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

// 产品创建
func TestProductCreate(t *testing.T) {

}

// 产品更新
func TestProductModify(t *testing.T) {

}

// 产品删除
func TestProductDelete(t *testing.T) {

}
