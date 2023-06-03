package obs

import (
	"encoding/json"
	"google.golang.org/grpc/codes"
)

type Status struct {
	s statusInternal
}

func New(code codes.Code, event string) *Status {
	s := &Status{
		s: statusInternal{
			Code:  code,
			Event: event,
		},
	}
	return s
}

func (s *Status) Err() error {
	if s.GetCode() == codes.OK {
		return nil
	} else {
		return &Error{s}
	}
}

func (s *Status) GetCode() codes.Code {
	if s == nil {
		return codes.OK
	} else {
		return s.s.Code
	}
}

func (s *Status) GetMsg() string {
	if s == nil {
		return ""
	} else {
		return s.s.Msg
	}
}

func (s *Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.s)
}

type statusInternal struct {
	Code  codes.Code `json:"code"`
	Msg   string     `json:"msg,omitempty"`
	Event string     `json:"event,omitempty"`
	Pc    uintptr    `json:"file,omitempty"`
}

type Error struct {
	s *Status
}

func (e Error) Error() string {
	return e.s.GetMsg()
}
