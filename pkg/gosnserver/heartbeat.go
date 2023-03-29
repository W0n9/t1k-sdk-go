package gosnserver

import (
	"io"

	"git.in.chaitin.net/patronus/t1k-sdk/sdk/go/pkg/t1k"
)

func DoHeartbeat(s io.ReadWriter) error {
	h := t1k.MakeHeader(t1k.MASK_FIRST|t1k.MASK_LAST, 0)
	_, err := s.Write(h.Serialize())
	if err != nil {
		return err
	}
	_, err = readDetectionResult(s)
	return err
}
