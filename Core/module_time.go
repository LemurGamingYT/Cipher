package core

import (
	"time"
)

type TimeLib struct {
}

func (tl *TimeLib) TimeNow(_ []any, _ *Visitor) *StringObject {
	return NewStringObject(time.Now().String())
}

func NewTimeLib() TimeLib {
	return TimeLib{}
}
