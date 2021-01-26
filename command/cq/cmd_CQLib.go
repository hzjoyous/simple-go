package cq

import (
	"math/rand"
	"time"
)

func randTrue(Molecular int, Denominator int) bool {
	return rand.Intn(Denominator) < Molecular
}

func atCQCode(userID string) string {
	return "[CQ:at,qq=" + userID + "]"
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
