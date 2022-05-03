package feishu

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Sdk struct {
	appID     string
	appSecret string
	token     string
	client    http.Client
}

func NewSDK(appID string, appSecret string) *Sdk {
	s := &Sdk{
		appID:     appID,
		appSecret: appSecret,
		client:    http.Client{},
	}

	if appID != "" && appSecret != "" {
		s.refreshToken()
	}

	return s
}

type tokenRequest struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type tokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

// TenantAccessToken https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
func (s Sdk) TenantAccessToken() (*tokenResponse, error) {
	request := tokenRequest{
		AppID:     s.appID,
		AppSecret: s.appSecret,
	}
	var response tokenResponse
	err := s.post("https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/", s.token, request, &response)
	if err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, errors.New(fmt.Sprintf("code: %d, err: %s", response.Code, response.Msg))
	}

	return &response, nil
}

// wired response:
// response of success
//{
//    "Extra": null,
//    "StatusCode": 0,
//    "StatusMessage": "success"
//}
// response of failure
//{
//    "code": 99991300,
//    "msg": "invalid request body: not json, invalid character '\\n' in string literal"
//}
type webhookV2Response struct {
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
	Code          int    `json:"code"`
	Msg           string `json:"msg"`
}

func (s Sdk) WebhookV2(webhook string, body io.Reader) error {
	req, err := http.NewRequest("POST", webhook, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	do, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer do.Body.Close()

	var resp webhookV2Response
	err = json.NewDecoder(do.Body).Decode(&resp)
	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return errors.New(fmt.Sprintf("code: %d, err: %s", resp.Code, resp.Msg))
	}

	return nil
}

func (s Sdk) get(url string, auth string, responseBody interface{}) error {
	return s.call("GET", url, auth, nil, responseBody)
}

func (s Sdk) post(url string, auth string, requestBody, responseBody interface{}) error {
	return s.call("POST", url, auth, requestBody, responseBody)
}

func (s Sdk) call(method string, url string, auth string, requestBody, responseBody interface{}) error {
	var body io.Reader
	if requestBody != nil {
		bs, err := json.Marshal(requestBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(bs)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if auth != "" {
		req.Header.Set("Authorization", "Bearer "+auth)
	}

	do, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer do.Body.Close()

	err = json.NewDecoder(do.Body).Decode(&responseBody)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sdk) refreshToken() {
	response, err := s.TenantAccessToken()
	if err != nil {
		// sleep and try again
		time.Sleep(time.Second * 1)
		s.refreshToken()
		return
	}
	s.token = response.TenantAccessToken

	// https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
	// Token 有效期为 2 小时，在此期间调用该接口 token 不会改变。当 token 有效期小于 30 分的时候，再次请求获取 token 的时候，会生成一个新的 token，与此同时老的 token 依然有效。
	// 在过期前 1 分钟刷新
	time.AfterFunc(time.Second*time.Duration(response.Expire-60), func() {
		s.refreshToken()
	})
}
