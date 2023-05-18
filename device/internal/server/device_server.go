package server

import (
	"context"

	"github.com/iot-platform/device/internal/logic"
	"github.com/iot-platform/device/internal/svc"
	"github.com/iot-platform/device/types/device"
)

type DeviceServer struct {
	svcCtx *svc.ServiceContext
	device.UnimplementedDeviceServer
}

func NewDeviceServer(svcCtx *svc.ServiceContext) *DeviceServer {
	return &DeviceServer{
		svcCtx: svcCtx,
	}
}

func (s *DeviceServer) SendMessage(ctx context.Context, in *device.SendMessageRequest) (*device.SendMessageReply, error) {
	l := logic.NewSendMessageLogic(ctx, s.svcCtx)
	return l.SendMessage(in)
}
