package detection

import (
	"git.in.chaitin.net/patronus/t1k-sdk/sdk/go/pkg/misc"
)

type DetectionContext struct {
	UUID         string
	Scheme       string
	ProxyName    string
	RemoteAddr   string
	RemotePort   uint16
	LocalAddr    string
	LocalPort    uint16
	ReqBeginTime int64
	RspBeginTime int64

	T1KContext []byte

	Request  Request
	Response Response
}

func New() *DetectionContext {
	return &DetectionContext{
		UUID:       misc.GenUUID(),
		Scheme:     "http",
		ProxyName:  "go-sdk",
		RemoteAddr: "127.0.0.1",
		RemotePort: 30001,
		LocalAddr:  "127.0.0.1",
		LocalPort:  80,
	}
}

func (dc *DetectionContext) ProcessResult(r *Result) {
	if r.Objective == RO_REQUEST {
		dc.T1KContext = r.T1KContext
	}
}
