package CqGoClient

import (
	"dog/util"
	"github.com/go-resty/resty/v2"
)

type remoteService struct {
	httpClient *resty.Client
}

func newRemoteService() remoteService {
	return remoteService{
		httpClient: resty.New().SetHostURL("http://127.0.0.1:5700"),
	}
}

func (s *remoteService) getFriendList() (resp *resty.Response, err error) {
	return s.httpClient.R().SetQueryParams(map[string]string{
	}).Get("/get_friend_list")
}

type GroupList struct {
	Data []struct {
		GroupCreateTime int    `json:"group_create_time"`
		GroupID         int    `json:"group_id"`
		GroupLevel      int    `json:"group_level"`
		GroupMemo       string `json:"group_memo"`
		GroupName       string `json:"group_name"`
		MaxMemberCount  int    `json:"max_member_count"`
		MemberCount     int    `json:"member_count"`
	} `json:"data"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

func (s remoteService) getGroupList() (resp *resty.Response, err error) {
	return s.httpClient.R().Get("/get_group_list")
}

func (s remoteService) sendGroupMsg(group_id int64,message string) (resp *resty.Response, err error) {
	return s.httpClient.R().SetBody(map[string]interface{}{
		"group_id": util.ToString(group_id),
		"message":  message,
	}).Get("/send_group_msg")
}
