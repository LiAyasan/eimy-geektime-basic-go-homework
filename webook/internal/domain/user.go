package domain

import "time"

// User 领域对象，是 DDD 中的聚合根或 entity
// BO(business object)
type User struct {
	Id       int64
	Email    string
	Password string
	Ctime    time.Time
	UserExtend
}

type UserExtend struct {
	Nickname string `json:"nickname"`
	Birthday string `json:"birthday"`
	Details  string `json:"details"`
}
