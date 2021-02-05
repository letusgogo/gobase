package util

import (
	"encoding/json"
	"errors"
	"git.iothinking.com/base/gobase/conf"
	"git.iothinking.com/base/gobase/log"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"go.uber.org/zap"
)

////////////////////////////// 短信服务 //////////////////////////
type SmsService struct {
	smsClient *sdk.Client
	smsConf   *conf.SmsConf
}

func NewSmsService(smsClient *sdk.Client, smsConf *conf.SmsConf) *SmsService {
	return &SmsService{smsClient: smsClient, smsConf: smsConf}
}

// 短信的响应
type smsRsp struct {
	Message   string `json:"Message"`
	RequestID string `json:"RequestId"`
	BizID     string `json:"BizId"`
	Code      string `json:"Code"`
}

// telephone 发送短信的号码
// templateParam 发送的模板消息
func (s *SmsService) sendSmsMsg(telephone, templateParam string) error {
	// 开启了 debug 模式不真正发短信
	if s.smsConf.Debug {
		log.Debug("send sms msg", zap.String("msg", templateParam))
		return nil
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = telephone
	request.QueryParams["SignName"] = s.smsConf.SignName
	request.QueryParams["TemplateCode"] = s.smsConf.TemplateCode
	request.QueryParams["TemplateParam"] = templateParam

	response, err := s.smsClient.ProcessCommonRequest(request)
	if err != nil {
		return err
	}
	if response.GetHttpStatus() != 200 {
		return errors.New("ali sms error return 500, response:" + response.GetHttpContentString())
	}
	// 判断返回的 Code
	smsRsp := &smsRsp{}
	_ = json.Unmarshal(response.GetHttpContentBytes(), smsRsp)
	if smsRsp.Code != "OK" {
		return errors.New(smsRsp.Message)
	} else {
		return nil
	}
}

type SmsDevEventService struct {
	*SmsService
}

func NewSmsDevEventService(smsService *SmsService) *SmsDevEventService {
	return &SmsDevEventService{SmsService: smsService}
}

// 发送设备信息通知短信
type SmsEvent struct {
	Name     string `json:"name"`
	DevName  string `json:"devName"`
	RuleName string `json:"ruleName"`
}

func (s *SmsDevEventService) SendEventSms(telephone string, eventSms *SmsEvent) error {
	eventSmsJsonStr, _ := json.Marshal(eventSms)
	return s.sendSmsMsg(telephone, string(eventSmsJsonStr))
}
