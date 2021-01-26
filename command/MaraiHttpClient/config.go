package MaraiHttpClient



type qqPeople struct {
	Name     string `json:"name"`
	QQNumber string `json:"QQNumber"`
}

var qqPeopleList []qqPeople

func init() {
	qqPeopleList = append(qqPeopleList, qqPeople{Name: "董珈毓", QQNumber: "3256439706"})
	qqPeopleList = append(qqPeopleList, qqPeople{Name: "倪家楠", QQNumber: "774340277"})
	qqPeopleList = append(qqPeopleList, qqPeople{Name: "杨晓琼", QQNumber: "626726829"})
	qqPeopleList = append(qqPeopleList, qqPeople{Name: "涂青青", QQNumber: "3346445171"})
	qqPeopleList = append(qqPeopleList, qqPeople{Name: "乌鱼子", QQNumber: "1492799403"})
	qqPeopleList = append(qqPeopleList, qqPeople{Name: "张柳", QQNumber: "1320457569"})
	//qqPeopleList = append(qqPeopleList, qqPeople{Name: "佩佩", QQNumber: "2508764466"})
	qqPeopleList = append(qqPeopleList, qqPeople{Name: "柠夏初开", QQNumber: "2338900259"})
}

func getQQPeopleList() []qqPeople {
	return qqPeopleList
}
