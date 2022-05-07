package util

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
	"github.com/letusgogo/gobase/conf"
	"github.com/letusgogo/gobase/log"
	"go.uber.org/zap"
)

// 发送语音
type TelService struct {
	telConf   *conf.TelConf
	telClient *dyvmsapi.Client
}

func NewTelService(telConf *conf.TelConf, telClient *dyvmsapi.Client) *TelService {
	return &TelService{telConf: telConf, telClient: telClient}
}

func (t *TelService) SendVoice(telephone, ttsParam string) error {
	// 开启了 debug 模式不真正打电话
	if t.telConf.Debug {
		log.Warn("send telephone msg", zap.String("msg", ttsParam))
		return nil
	}
	request := dyvmsapi.CreateSingleCallByTtsRequest()
	request.Scheme = "https"
	request.CalledNumber = telephone
	request.TtsCode = t.telConf.TemplateCode
	request.TtsParam = ttsParam

	// 播放次数
	request.PlayTimes = requests.NewInteger(2)

	response, err := t.telClient.SingleCallByTts(request)
	if err != nil {
		return err
	}

	if response.Code != "OK" {
		return errors.New(response.Message)
	} else {
		return nil
	}
}
