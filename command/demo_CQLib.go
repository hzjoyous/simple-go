package command

import (
	"math/rand"
	"time"
)

func randTrue(Molecular int,Denominator int) bool{
	return rand.Intn(Denominator) <Molecular
}

func atCQCode(userId string) string{
	return "[CQ:at,qq=" + userId + "]"
}

func init(){
	rand.Seed(time.Now().UnixNano())
}
