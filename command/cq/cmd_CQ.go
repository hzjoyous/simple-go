package cq

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"simple-go/command/console"
	"simple-go/command/util"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)



func init() {
	c := console.Console{Signature: "cmdCQ", Description: "this is a template", Handle: cmdCQ}
	commandList[c.Signature] = c
}


func cmdCQ() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/api/cmd", cmdCQHttpHandle)
	//work()
	go work()
	_ = r.Run(":8000")
}

type replyData struct {
	counter int64
}

func work() {
	fmt.Println("work  start run ")

	// aiqHTTPClient := newAIqHTTPClient()
	// resp, err := aiqHTTPClient.say("王者荣耀", "12346")

	// fmt.Println("Response Info:")
	// fmt.Println("Error      :", err)
	// fmt.Println("Status Code:", resp.StatusCode())
	// fmt.Println("Status     :", resp.Status())
	// fmt.Println("Time       :", resp.Time())
	// fmt.Println("Received At:", resp.ReceivedAt())
	// fmt.Println("Body       :", resp.String())

	// value := gjson.Get(resp.String(), "result.response_list.0.action_list.#.say")
	// for _, name := range value.Array() {
	// 	fmt.Println(name.String())
	// }
	// fmt.Println(value.String())
}

func cmdCQHttpHandle(c *gin.Context) {
	reqInfo := message{}
	//Headers were already written. Wanted to override status code 400 with 200
	//err := c.BindJSON(&reqInfo)
	err := c.ShouldBind(&reqInfo)

	if err != nil {
		fmt.Println("上报数据异常")
	}

	cqHTTPClient := newcqHTTPClient()
	needContinueArr := []string{"然后呢", "请继续", "接着说", "我不知道说什么了", "请注意文明用语", "你可以任性去理解", "有电脑为什么要手机改啊？", "在思考吗？", "一整就来些符号，也太考验我的理解力了。。。"}
	switch reqInfo.PostType {
	case postTypeMessage:
		aiqHTTPClient := newAIqHTTPClient()
		switch reqInfo.MessageType {
		case messageTypeGroup:
			// init
			rand.Seed(time.Now().UnixNano())
			groupID := reqInfo.GroupID
			userID := reqInfo.UserID
			sessionID := strconv.FormatInt(reqInfo.GroupID, 10) + strconv.FormatInt(reqInfo.UserID, 10)

			// check
			// check1 群号过滤
			needArr := []int64{325405886, 46938920, 57419059, 93050305, 97431784, 122917448, 160308765, 169294352, 171062298, 181043219, 228667664, 230078413, 241142401, 275879130, 291028285, 295305303, 308127235, 332463685, 340630794, 363324037, 370767642, 429732029, 459530943, 467204389, 467941966, 492548647, 536231481, 564784122, 584657835, 612908090, 625012253, 635290770, 674584784, 693931666, 731990104, 739654999, 789788805, 810919826, 820698944, 858684210, 874415430, 936046286, 970683037, 979359071, 1083478826}
			if util.InArrayInt64(groupID, needArr) {
				break
			}
			tmpNeedArr := []int64{733530788, 484614174}
			if util.InArrayInt64(groupID, tmpNeedArr) {
				break
			}

			// check2 图片不予回复
			if strings.Contains(reqInfo.Message, "CQ:image,file") {
				if randTrue(1, 5) {
					randFaceCode := getRandFaceRepeat(3)
					_, _ = cqHTTPClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupID, 10), randFaceCode)
					fmt.Println("[CQ-hard]", util.Date(), " |来自:", reqInfo.GroupID, "| rawMessage:", reqInfo.RawMessage, "|send:", randFaceCode)
				} else {
					fmt.Println("[CQ-hard]", util.Date(), " |来自:", reqInfo.GroupID, "| cantRun is True | rawMessage:", reqInfo.RawMessage)
				}
				break
			}

			cantRun := false
			group1d4NeedArr := []int64{733530788}
			group1d3NeedArr := []int64{893422240}
			if util.InArrayInt64(reqInfo.GroupID, group1d4NeedArr) {
				cantRun = randTrue(3, 4)
			} else if util.InArrayInt64(groupID, group1d3NeedArr) {
				cantRun = randTrue(2, 3)
			} else {
				cantRun = randTrue(3, 4)
			}

			atRand := false

			// 用户特定筛选
			switch userID {
			case 2033369740:
				cantRun = false
				atRand = randTrue(1, 3)
				break
			case 1426148118:
				if time.Now().Hour() < 5 {
					cantRun = false
				} else {
					cantRun = randTrue(1, 2)
				}
				atRand = randTrue(1, 2)
				break
			case 1540025138:
			case 3521207082:
				if randTrue(1, 2) {
					_, _ = cqHTTPClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupID, 10), atCQCode(strconv.FormatInt(userID, 10))+"\n[贤者模式-running-1540025138-3521207082]\n以热爱祖国为荣，以危害祖国为耻。\n以服务人民为荣，以背离人民为耻。\n以崇尚科学为荣，以愚昧无知为耻。\n以辛勤劳动为荣，以好逸恶劳为耻。\n以团结互助为荣，以损人利己为耻。\n以诚实守信为荣，以见利忘义为耻。\n以遵纪守法为荣，以违法乱纪为耻。\n以艰苦奋斗为荣，以骄奢淫逸为耻。")

				}
				return
			default:
				atRand = randTrue(1, 5)
				break
			}

			if cantRun {
				_, _ = aiqHTTPClient.say(reqInfo.Message, strconv.FormatInt(reqInfo.GroupID, 10)+strconv.FormatInt(reqInfo.UserID, 10))
				fmt.Println("[CQ-hard]", util.Date(), "|来自", groupID, "|跳过原因cantRun:", cantRun, "| rawMessage:", reqInfo.RawMessage)
				break
			}

			atMessage := ""
			if atRand {
				atMessage = atCQCode(strconv.FormatInt(userID, 10)) + getRandFace()
			}

			time.Sleep(3 * time.Second)

			// run
			resp, _ := aiqHTTPClient.say(reqInfo.Message, sessionID)
			value := gjson.Get(resp.String(), "result.response_list.0.action_list.#.say")

			var sayCounter int
			var sayMessage string

			for _, say := range value.Array() {
				sayCounter++

				sayMessage = say.String()
				if util.InArrayString(sayMessage, needContinueArr) {
					sayMessage = getRandMessage()
				}
				resp, _ := cqHTTPClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupID, 10), atMessage+sayMessage)

				fmt.Println(groupID, ":")
				fmt.Println("	Message	:", reqInfo.Message)
				fmt.Println("	Say		:", sayMessage)
				fmt.Println("	Body    :", resp.String())
				break
			}
			if sayCounter == 0 {
				resp, _ = cqHTTPClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupID, 10), "😴,,,,")
			}

			fmt.Println("[CQ-SUCCESS]", util.Date(), "|groupID:", reqInfo.GroupID, ",内容:", reqInfo.RawMessage)
			break

		case messageTypePrivate:
			resp, _ := aiqHTTPClient.say(reqInfo.Message, strconv.FormatInt(reqInfo.GroupID, 10))
			value := gjson.Get(resp.String(), "result.response_list.0.action_list.#.say")
			userID := reqInfo.UserID

			if strings.Contains(reqInfo.Message, "CQ:image,file") {
				randFaceCode := getRandFaceRepeat(3)
				_, _ = cqHTTPClient.sendPrivateMsg(strconv.FormatInt(reqInfo.UserID, 10), randFaceCode)
				fmt.Println("[CQ-hard]", util.Date(), " |来自:", reqInfo.UserID, "| rawMessage:", reqInfo.RawMessage, "|send:", randFaceCode)
				break
			}
			sayMessage := ""
			for _, say := range value.Array() {
				sayMessage = say.String()
				if util.InArrayString(sayMessage, needContinueArr) {
					sayMessage = getRandMessage()
				}
				resp, err := cqHTTPClient.sendPrivateMsg(strconv.FormatInt(reqInfo.UserID, 10), sayMessage)
				fmt.Println(userID, say.String())
				fmt.Println("Error      :", err)
				fmt.Println("Body       :", resp.String())
			}
			time.Sleep(3 * time.Second)

			fmt.Println("处理类型私人信息,来自", strconv.FormatInt(userID, 10), "内容", reqInfo.Message, reqInfo.RawMessage)
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
	UserID      int64  `json:"user_id"`
	GroupID     int64  `json:"group_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
}

// newcqHTTPClient
func newcqHTTPClient() cqHTTPClient {
	cqHTTPClient := cqHTTPClient{httpClient: resty.New().SetHostURL("http://127.0.0.1:6701")}
	return cqHTTPClient
}

type cqHTTPClient struct {
	httpClient *resty.Client
}

func (cqHTTPClient cqHTTPClient) getUserInfo(userID int, noCache bool) (resp *resty.Response, err error) {
	resp, err = cqHTTPClient.httpClient.R().SetQueryParams(map[string]string{
		"user_id":  strconv.Itoa(userID),
		"no_cache": strconv.FormatBool(false),
	}).Get("/send_private_msg")
	return

}
func (cqHTTPClient cqHTTPClient) sendPrivateMsg(userID string, message string) (resp *resty.Response, err error) {
	resp, err = cqHTTPClient.httpClient.R().SetQueryParams(map[string]string{
		"user_id":     userID,
		"message":     message,
		"auto_escape": strconv.FormatBool(false),
	}).Get("/send_private_msg")
	return
}
func (cqHTTPClient cqHTTPClient) sendGroupMsg(groupID string, message string) (resp *resty.Response, err error) {
	resp, err = cqHTTPClient.httpClient.R().SetQueryParams(map[string]string{
		"group_id":    groupID,
		"message":     message,
		"auto_escape": strconv.FormatBool(false),
	}).Get("/send_group_msg")
	return
}

// newAIqHTTPClient
func newAIqHTTPClient() aiqHTTPClient {
	aiqHTTPClient := aiqHTTPClient{httpClient: resty.New().SetHostURL("https://aip.baidubce.com")}
	//aiqHTTPClient := aiqHTTPClient{httpClient: resty.New().SetHostURL("http://localhost:8000")}
	return aiqHTTPClient
}

type aiqHTTPClient struct {
	httpClient *resty.Client
}

func (aiqHTTPClient aiqHTTPClient) getAccessToken() (resp *resty.Response, err error) {
	resp, err = aiqHTTPClient.httpClient.R().SetQueryParams(map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     "LEVlCI9ymTsByK5PPIis41zV",
		"client_secret": "dsEzNQOnpyZ7TK2xxP2ouKaPrjtPFyhQ",
	}).Post("/oauth/2.0/token")
	return
}

func (aiqHTTPClient aiqHTTPClient) say(message string, sessionID string) (resp *resty.Response, err error) {
	request := aiqHTTPClientSayRequest{
		LogID:     "UNITTEST_10000",
		Version:   "2.0",
		ServiceID: "S29166",
		SessionID: sessionID,
		Request: aiqHTTPClientSayRequestRequest{
			Query:  message,
			UserID: "1234567890",
		},
		DialogState: aiqHTTPClientSayRequestDialogState{
			Contexts: aiqHTTPClientSayRequestSysRememberedSkills{
				SysRememberedSkills: make([]string, 0),
			},
		},
	}
	jsonString, _ := json.Marshal(request)
	requestStr := string(jsonString)

	resp, err = aiqHTTPClient.httpClient.R().SetQueryParams(map[string]string{
		"access_token": "24.9b509b5640882f31ce29f4152f660768.2592000.1590129949.282335-19549928",
	}).
		SetHeader("Content-Type", "application/json").
		SetBody(requestStr).
		Post("/rpc/2.0/unit/service/chat")
	return
}

type aiqHTTPClientSayRequest struct {
	LogID       string                             `json:"log_id"`
	Version     string                             `json:"version"`
	ServiceID   string                             `json:"service_id"`
	SessionID   string                             `json:"session_id"`
	Request     interface{}                        `json:"request"`
	DialogState aiqHTTPClientSayRequestDialogState `json:"dialog_state"`
}

type aiqHTTPClientSayRequestRequest struct {
	Query  string `json:"query"`
	UserID string `json:"user_id"`
}

type aiqHTTPClientSayRequestDialogState struct {
	Contexts aiqHTTPClientSayRequestSysRememberedSkills `json:"contexts"`
}

type aiqHTTPClientSayRequestSysRememberedSkills struct {
	SysRememberedSkills []string `json:"SYS_REMEMBERED_SKILLS"`
}
