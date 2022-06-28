package middleware

import "time"

var Users = map[string]string{
	"vania": "iya",
}
var Sessions = map[string]Session{}

type Session struct {
	Username string
	Expired  time.Time
}

type Models struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func (s *Session) IsExpired() bool {
	return s.Expired.Before(time.Now())
}
