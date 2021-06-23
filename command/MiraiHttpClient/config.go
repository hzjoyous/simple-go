package MiraiHttpClient

import "strconv"

type qqFriendEntity struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

func (receiver qqFriendEntity) getQQNumber() string {
	return strconv.Itoa(receiver.ID)
}

var qqFriendEntityList []qqFriendEntity

func init() {
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "董珈毓", ID: 3256439706})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "倪家楠", ID: 774340277})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "杨晓琼", ID: 626726829})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "涂青青", ID: 3346445171})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "乌鱼子", ID: 1492799403})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "张柳", ID: 1320457569})
	//qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "佩佩", ID: 2508764466})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "柠夏初开", ID: 2338900259})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "☾  ☾  ☾", ID: 2030204736})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "", ID: 2686298697})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "", ID: 2039799133})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "", ID: 1182996608})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "", ID: 1182996608})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "", ID: 2033254708})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "", ID: 2033254708})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "", ID: 1411880115})
	qqFriendEntityList = append(qqFriendEntityList, qqFriendEntity{Remark: "", ID: 88958033})
}

func getQQFriendEntityList() []qqFriendEntity {
	return qqFriendEntityList
}
