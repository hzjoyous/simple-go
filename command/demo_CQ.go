package command

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	resty "github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type demoCQ struct {
	ConsoleInterface
}

func init() {
	var command ConsoleInterface
	command = new(demoCQ)
	commandList[command.GetSignature()] = command
}

func (demoCQ demoCQ) GetSignature() string {
	return "demoCQ"
}

func (demoCQ demoCQ) GetDescription() string {
	return "this is a Description"
}

func (demoCQ demoCQ) Handle() {
	r := gin.Default()
	r.POST("/api/demo", demoCQHttpHandle)
	//work()
	go work()
	_ = r.Run(":8000")
}

func work() {
	fmt.Println("work  start run ")

	aiqHttpClient := newAIqHttpClient()
	resp, err := aiqHttpClient.say("王者荣耀", "12346")

	// Explore response object
	//cqHttpClient := newCQHttpClient()
	//resp, err := cqHttpClient.sendPrivateMsg(2776404988, "heiheihei[CQ:face,id=182]")
	fmt.Println("Response Info:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status     :", resp.Status())
	fmt.Println("Time       :", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Body       :", resp.String())

	value := gjson.Get(resp.String(), "result.response_list.0.action_list.#.say")
	for _, name := range value.Array() {
		fmt.Println(name.String())
	}
	fmt.Println(value.String())
}

func demoCQHttpHandle(c *gin.Context) {
	reqInfo := message{}
	//Headers were already written. Wanted to override status code 400 with 200
	//err := c.BindJSON(&reqInfo)
	err := c.ShouldBind(&reqInfo)

	if err != nil {
		fmt.Println("上报数据异常")
	} else {
		fmt.Println(reqInfo.PostType)
		fmt.Println(reqInfo)
	}
	cqHttpClient := newCQHttpClient()

	switch reqInfo.PostType {
	case postTypeMessage:
		switch reqInfo.MessageType {
		case messageTypeGroup:
			groupId := reqInfo.GroupId
			aiqHttpClient := newAIqHttpClient()
			resp, _ := aiqHttpClient.say(reqInfo.Message, strconv.FormatInt(reqInfo.GroupId, 10))
			value := gjson.Get(resp.String(), "result.response_list.0.action_list.#.say")
			for _, say := range value.Array() {
				resp, err := cqHttpClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupId, 10), say.String())
				fmt.Println(groupId, say.String(), reqInfo.GroupId)
				fmt.Println("Response Info:")
				fmt.Println("Error      :", err)
				fmt.Println("Body       :", resp.String())
				//resp, err = cqHttpClient.sendPrivateMsg(2776404988, "heiheihei[CQ:face,id=182]")
			}
			fmt.Println("处理类型群消息,来自", reqInfo.GroupId, "内容", reqInfo.Message, reqInfo.RawMessage)
			break
		case messageTypePrivate:
			fmt.Println("处理类型私人信息,来自", reqInfo.UserId, "内容", reqInfo.Message, reqInfo.RawMessage)
			break
		default:
			break
		}
		break
	case postTypeNotice:
		break
	case postTypeRequest:
		break
	default:
		break
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

const (
	postTypeMessage string = "message"
	postTypeNotice  string = "notice"
	postTypeRequest string = "request"
)
const (
	messageTypePrivate string = "private"
	messageTypeGroup   string = "group"
)

type message struct {
	MessageType string `json:"message_type"`
	PostType    string `json:"post_type"`
	UserId      string `json:"user_id"`
	GroupId     int64  `json:"group_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
}

// newCQHttpClient
func newCQHttpClient() cqHttpClient {
	cqHttpClient := cqHttpClient{httpClient: resty.New().SetHostURL("http://127.0.0.1:6701")}
	return cqHttpClient
}

type cqHttpClient struct {
	httpClient *resty.Client
}

func (cqHttpClient cqHttpClient) getUserInfo(userId int, noCache bool) (resp *resty.Response, err error) {
	resp, err = cqHttpClient.httpClient.R().SetQueryParams(map[string]string{
		"user_id":  strconv.Itoa(userId),
		"no_cache": strconv.FormatBool(false),
	}).Get("/send_private_msg")
	return

}
func (cqHttpClient cqHttpClient) sendPrivateMsg(userId int, message string) (resp *resty.Response, err error) {
	resp, err = cqHttpClient.httpClient.R().SetQueryParams(map[string]string{
		"user_id":     strconv.Itoa(userId),
		"message":     message,
		"auto_escape": strconv.FormatBool(false),
	}).Get("/send_private_msg")
	return
}
func (cqHttpClient cqHttpClient) sendGroupMsg(groupId string, message string) (resp *resty.Response, err error) {
	resp, err = cqHttpClient.httpClient.R().SetQueryParams(map[string]string{
		"group_id":    groupId,
		"message":     message,
		"auto_escape": strconv.FormatBool(false),
	}).Get("/send_group_msg")
	return
}

// newAIqHttpClient
func newAIqHttpClient() aiqHttpClient {
	aiqHttpClient := aiqHttpClient{httpClient: resty.New().SetHostURL("https://aip.baidubce.com")}
	//aiqHttpClient := aiqHttpClient{httpClient: resty.New().SetHostURL("http://localhost:8000")}
	return aiqHttpClient
}

type aiqHttpClient struct {
	httpClient *resty.Client
}

func (aiqHttpClient aiqHttpClient) getAccessToken() (resp *resty.Response, err error) {
	resp, err = aiqHttpClient.httpClient.R().SetQueryParams(map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     "LEVlCI9ymTsByK5PPIis41zV",
		"client_secret": "dsEzNQOnpyZ7TK2xxP2ouKaPrjtPFyhQ",
	}).Post("/oauth/2.0/token")
	return
}

func (aiqHttpClient aiqHttpClient) say(message string, sessionId string) (resp *resty.Response, err error) {
	request := aiqHttpClientSayRequest{
		LogId:     "UNITTEST_10000",
		Version:   "2.0",
		ServiceId: "S29166",
		SessionId: sessionId,
		Request: aiqHttpClientSayRequestRequest{
			Query:  message,
			UserId: "1234567890",
		},
		DialogState: aiqHttpClientSayRequestDialogState{
			Contexts: aiqHttpClientSayRequestSysRememberedSkills{
				SysRememberedSkills: make([]string, 0),
			},
		},
	}
	jsonString, _ := json.Marshal(request)
	requestStr := string(jsonString)

	fmt.Println(requestStr)
	resp, err = aiqHttpClient.httpClient.R().SetQueryParams(map[string]string{
		"access_token": "24.9b509b5640882f31ce29f4152f660768.2592000.1590129949.282335-19549928",
	}).
		SetHeader("Content-Type", "application/json").
		SetBody(requestStr).
		Post("/rpc/2.0/unit/service/chat")
	return
}

type aiqHttpClientSayRequest struct {
	LogId       string                             `json:"log_id"`
	Version     string                             `json:"version"`
	ServiceId   string                             `json:"service_id"`
	SessionId   string                             `json:"session_id"`
	Request     interface{}                        `json:"request"`
	DialogState aiqHttpClientSayRequestDialogState `json:"dialog_state"`
}

type aiqHttpClientSayRequestRequest struct {
	Query  string `json:"query"`
	UserId string `json:"user_id"`
}

type aiqHttpClientSayRequestDialogState struct {
	Contexts aiqHttpClientSayRequestSysRememberedSkills `json:"contexts"`
}

type aiqHttpClientSayRequestSysRememberedSkills struct {
	SysRememberedSkills []string `json:"SYS_REMEMBERED_SKILLS"`
}
