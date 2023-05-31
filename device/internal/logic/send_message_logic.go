package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/iot-platform/device/internal/mqtt"
	"github.com/iot-platform/device/internal/svc"
	"github.com/iot-platform/device/types/device"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMessageRequest) (*device.SendMessageReply, error) {
	if in.ProductKey == "" || in.DeviceKey == "" || in.Data == "" {
		return nil, errors.New("参数异常")
	}

	topic := fmt.Sprintf("/sys/%v/%v/receive", in.ProductKey, in.DeviceKey)
	// 发布
	if token := mqtt.MC.Publish(topic, 0, false, in.Data); token.Wait() && token.Error() != nil {
		logx.Error("[PUBLISH ERROR]: ", token.Error())
		return nil, token.Error()
	}

	return &device.SendMessageReply{}, nil
}
