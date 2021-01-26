package MaraiHttpClient

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type maraiHttpClient struct {
	httpClient    *resty.Client
	sessionKey    string
	adminQQNumber string
	authKey       string
}

func newMaraiClient(authKey string, adminQQNumber string) maraiHttpClient {
	maraiHttpClient := maraiHttpClient{
		httpClient:    resty.New().SetHostURL("http://127.0.0.1:8080"),
		adminQQNumber: adminQQNumber,
		authKey:       authKey,
	}
	return maraiHttpClient
}

func (receiver *maraiHttpClient) checkSessionByCode(code int) error {

	var err error
	if code == 3 || code == 4 {
		return receiver.verifySession()
	}
	return err
}

func (receiver *maraiHttpClient) verifySession() error {
	var err error
	err = nil
	var authR authResponse
	result, _ := receiver.auth(receiver.authKey)
	err = json.Unmarshal([]byte(result.String()), &authR)

	if err != nil {
		fmt.Println(err)
		return err
	}

	receiver.setSessionKey(authR.Session)

	result, err = receiver.verify(receiver.adminQQNumber)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func (receiver *maraiHttpClient) setSessionKey(session string) {
	fmt.Println(session)
	receiver.sessionKey = session
}

type aboutResponse struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	ErrorMessage string `json:"errorMessage"`
	Data         struct {
		Version string `json:"version"`
	} `json:"data"`
}

/**
 */
func (receiver maraiHttpClient) getAbout() (resp *resty.Response, err error) {
	return receiver.httpClient.R().SetQueryParams(map[string]string{
	}).Get("/about")
}

type authResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Session string `json:"session"`
}

/**
获取session
*/
func (receiver maraiHttpClient) auth(authKey string) (resp *resty.Response, err error) {
	return receiver.httpClient.R().SetBody(map[string]string{
		"authKey": authKey,
	}).Post("/auth")
}

type verifyResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (receiver maraiHttpClient) verify(qq string) (resp *resty.Response, err error) {
	return receiver.httpClient.R().SetBody(map[string]string{
		"sessionKey": receiver.sessionKey,
		"qq":         qq,
	}).Post("/verify")
}

type releaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (receiver maraiHttpClient) release() (resp *resty.Response, err error) {
	return receiver.httpClient.R().SetQueryParams(map[string]string{
		"sessionKey": "YourSessionKey",
		"qq":         "123456789",
	}).Post("/release")
}

type sendFriendMessageRequest struct {
	SessionKey   string `json:"sessionKey"`
	Target       int    `json:"target,omitempty"`
	MessageChain []struct {
		Type string `json:"type"`
		Text string `json:"text,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"messageChain"`
}

type sendFriendMessageResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	MessageID int    `json:"messageId"`
}

func (receiver maraiHttpClient) sendFriendMessage(qq string, message string) (resp *resty.Response, err error) {
	fmt.Println(receiver.sessionKey)
	return receiver.httpClient.R().SetBody(map[string]interface{}{
		"sessionKey": receiver.sessionKey,
		"target":     qq,
		"messageChain": []interface{}{
			map[string]interface{}{"type": "Plain", "text": message},
		},
	}).Post("/sendFriendMessage")
}

// todo 临时会话
func (receiver maraiHttpClient) sendTempMessage() (resp *resty.Response, err error) {
	return receiver.httpClient.R().SetQueryParams(map[string]string{

	}).Post("")
}

type sendGroupMessageRequest struct {
	SessionKey   string `json:"sessionKey"`
	Target       string `json:"target"`
	MessageChain []struct {
		Type string `json:"type"`
		Text string `json:"text,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"messageChain"`
}

type sendGroupMessageResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	MessageID int    `json:"messageId"`
}

func (receiver maraiHttpClient) sendGroupMessage(GroupId string, message string) (resp *resty.Response, err error) {
	fmt.Println(receiver.sessionKey)
	return receiver.httpClient.R().SetBody(sendGroupMessageRequest{
		SessionKey: receiver.sessionKey,
		Target:     GroupId,
		MessageChain: []struct {
			Type string `json:"type"`
			Text string `json:"text,omitempty"`
			URL  string `json:"url,omitempty"`
		}{{Type: "Plain", Text: message}},
	}).Post("/sendGroupMessage")
}

type friendListResponse []struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

func (receiver maraiHttpClient) friendList(sessionKey string) (resp *resty.Response, err error) {
	return receiver.httpClient.R().SetQueryParams(map[string]string{
		"sessionKey": sessionKey,
	}).Get("/friendList")
}
