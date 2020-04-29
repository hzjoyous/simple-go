package command

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/api/demo", demoCQHttpHandle)
	//work()
	go work()
	_ = r.Run(":8000")
}

type replyData struct {
	counter int64
}

func work() {
	fmt.Println("work  start run ")

	//aiqHttpClient := newAIqHttpClient()
	//resp, err := aiqHttpClient.say("王者荣耀", "12346")
	//
	//fmt.Println("Response Info:")
	//fmt.Println("Error      :", err)
	//fmt.Println("Status Code:", resp.StatusCode())
	//fmt.Println("Status     :", resp.Status())
	//fmt.Println("Time       :", resp.Time())
	//fmt.Println("Received At:", resp.ReceivedAt())
	//fmt.Println("Body       :", resp.String())
	//
	//value := gjson.Get(resp.String(), "result.response_list.0.action_list.#.say")
	//for _, name := range value.Array() {
	//	fmt.Println(name.String())
	//}
	//fmt.Println(value.String())
}

func demoCQHttpHandle(c *gin.Context) {
	reqInfo := message{}
	//Headers were already written. Wanted to override status code 400 with 200
	//err := c.BindJSON(&reqInfo)
	err := c.ShouldBind(&reqInfo)

	if err != nil {
		fmt.Println("上报数据异常")
	}
	cqHttpClient := newCQHttpClient()
	needContinueArr := []string{"然后呢", "请继续", "接着说", "我不知道说什么了", "请注意文明用语", "你可以任性去理解", "有电脑为什么要手机改啊？", "在思考吗？","一整就来些符号，也太考验我的理解力了。。。"}
	switch reqInfo.PostType {
	case postTypeMessage:
		aiqHttpClient := newAIqHttpClient()
		switch reqInfo.MessageType {
		case messageTypeGroup:
			// init
			rand.Seed(time.Now().UnixNano())
			groupId := reqInfo.GroupId
			userId := reqInfo.UserId
			sessionId := strconv.FormatInt(reqInfo.GroupId, 10) + strconv.FormatInt(reqInfo.UserId, 10)

			// check
			// check1 群号过滤
			needArr := []int64{325405886, 46938920, 57419059, 93050305, 97431784, 122917448, 160308765, 169294352, 171062298, 181043219, 228667664, 230078413, 241142401, 275879130, 291028285, 295305303, 308127235, 332463685, 340630794, 363324037, 370767642, 429732029, 459530943, 467204389, 467941966, 492548647, 536231481, 564784122, 584657835, 612908090, 625012253, 635290770, 674584784, 693931666, 731990104, 739654999, 789788805, 810919826, 820698944, 858684210, 874415430, 936046286, 970683037, 979359071, 1083478826}
			if inArrayInt64(groupId, needArr) {
				break
			}
			tmpNeedArr := []int64{733530788, 484614174}
			if inArrayInt64(groupId, tmpNeedArr) {
				break
			}

			// check2 图片不予回复
			if strings.Contains(reqInfo.Message, "CQ:image,file") {
				if randTrue(1, 5) {
					randFaceCode := getRandFaceRepeat(3)
					_, _ = cqHttpClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupId, 10), randFaceCode)
					fmt.Println("[CQ-hard]", date(), " |来自:", reqInfo.GroupId, "| rawMessage:", reqInfo.RawMessage, "|send:", randFaceCode)
				} else {
					fmt.Println("[CQ-hard]", date(), " |来自:", reqInfo.GroupId, "| cantRun is True | rawMessage:", reqInfo.RawMessage)
				}
				break
			}

			cantRun := false
			group1d4NeedArr := []int64{733530788}
			group1d3NeedArr := []int64{893422240}
			if inArrayInt64(reqInfo.GroupId, group1d4NeedArr) {
				cantRun = randTrue(3, 4)
			} else if inArrayInt64(groupId, group1d3NeedArr) {
				cantRun = randTrue(2, 3)
			} else {
				cantRun = randTrue(3, 4)
			}

			atRand := false

			// 用户特定筛选
			switch userId {
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
					_, _ = cqHttpClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupId, 10), atCQCode(strconv.FormatInt(userId, 10))+"\n[贤者模式-running-1540025138-3521207082]\n以热爱祖国为荣，以危害祖国为耻。\n以服务人民为荣，以背离人民为耻。\n以崇尚科学为荣，以愚昧无知为耻。\n以辛勤劳动为荣，以好逸恶劳为耻。\n以团结互助为荣，以损人利己为耻。\n以诚实守信为荣，以见利忘义为耻。\n以遵纪守法为荣，以违法乱纪为耻。\n以艰苦奋斗为荣，以骄奢淫逸为耻。")

				}
				return
			default:
				atRand = randTrue(1, 5)
				break
			}

			if cantRun {
				_, _ = aiqHttpClient.say(reqInfo.Message, strconv.FormatInt(reqInfo.GroupId, 10)+strconv.FormatInt(reqInfo.UserId, 10))
				fmt.Println("[CQ-hard]", date(), "|来自", groupId, "|跳过原因cantRun:", cantRun, "| rawMessage:", reqInfo.RawMessage)
				break
			}

			atMessage := ""
			if atRand {
				atMessage = atCQCode(strconv.FormatInt(userId, 10)) + getRandFace()
			}

			time.Sleep(3 * time.Second)

			// run
			resp, _ := aiqHttpClient.say(reqInfo.Message, sessionId)
			value := gjson.Get(resp.String(), "result.response_list.0.action_list.#.say")

			var sayCounter int
			var sayMessage string

			for _, say := range value.Array() {
				sayCounter += 1

				sayMessage = say.String()
				if inArrayString(sayMessage, needContinueArr) {
					sayMessage = getRandMessage()
				}
				resp, _ := cqHttpClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupId, 10), atMessage+sayMessage)

				fmt.Println(groupId, ":")
				fmt.Println("	Message	:", reqInfo.Message)
				fmt.Println("	Say		:", sayMessage)
				fmt.Println("	Body    :", resp.String())
				break
			}
			if sayCounter == 0 {
				resp, _ = cqHttpClient.sendGroupMsg(strconv.FormatInt(reqInfo.GroupId, 10), "😴,,,,")
			}

			fmt.Println("[CQ-SUCCESS]", date(), "|groupId:", reqInfo.GroupId, ",内容:", reqInfo.RawMessage)
			break

		case messageTypePrivate:
			resp, _ := aiqHttpClient.say(reqInfo.Message, strconv.FormatInt(reqInfo.GroupId, 10))
			value := gjson.Get(resp.String(), "result.response_list.0.action_list.#.say")
			userId := reqInfo.UserId

			if strings.Contains(reqInfo.Message, "CQ:image,file") {
				randFaceCode := getRandFaceRepeat(3)
				_, _ = cqHttpClient.sendPrivateMsg(strconv.FormatInt(reqInfo.UserId, 10), randFaceCode)
				fmt.Println("[CQ-hard]", date(), " |来自:", reqInfo.UserId, "| rawMessage:", reqInfo.RawMessage, "|send:", randFaceCode)
				break
			}
			sayMessage := ""
			for _, say := range value.Array() {
				sayMessage = say.String()
				if inArrayString(sayMessage, needContinueArr) {
					sayMessage = getRandMessage()
				}
				resp, err := cqHttpClient.sendPrivateMsg(strconv.FormatInt(reqInfo.UserId, 10), sayMessage)
				fmt.Println(userId, say.String())
				fmt.Println("Error      :", err)
				fmt.Println("Body       :", resp.String())
			}
			time.Sleep(3 * time.Second)

			fmt.Println("处理类型私人信息,来自", strconv.FormatInt(userId, 10), "内容", reqInfo.Message, reqInfo.RawMessage)
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
	UserId      int64  `json:"user_id"`
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
func (cqHttpClient cqHttpClient) sendPrivateMsg(userId string, message string) (resp *resty.Response, err error) {
	resp, err = cqHttpClient.httpClient.R().SetQueryParams(map[string]string{
		"user_id":     userId,
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
