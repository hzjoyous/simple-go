package MiraiHttpClient

import "github.com/go-resty/resty/v2"

type HiTokotoClient struct {
	httpClient        *resty.Client
}

func newHiTokotoClient(host string) HiTokotoClient {
	if len(host) == 0 {
		host = "https://v1.hitokoto.cn/"
	}
	httpClient := HiTokotoClient{
		httpClient:    resty.New().SetHostURL("https://v1.hitokoto.cn/"),
	}
	return httpClient
}


type HiTokotoResponse struct {
	ID         int         `json:"id"`
	UUID       string      `json:"uuid"`
	Hitokoto   string      `json:"hitokoto"`
	Type       string      `json:"type"`
	From       string      `json:"from"`
	FromWho    interface{} `json:"from_who"`
	Creator    string      `json:"creator"`
	CreatorUID int         `json:"creator_uid"`
	Reviewer   int         `json:"reviewer"`
	CommitFrom string      `json:"commit_from"`
	CreatedAt  string      `json:"created_at"`
	Length     int         `json:"length"`
}

func (receiver HiTokotoClient) getOneTokoto()(resp *resty.Response, err error){
	return receiver.httpClient.R().Get("/")
}

