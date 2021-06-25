package week9

import (
	"bytes"
	"testing"
)

func TestProtocol(t *testing.T) {
	p := NewProto(0, 1, 2, []byte("test"))
	res := &Protocol{}
	res.Decode(p.Encode())

	if p.Ver != res.Ver {
		t.Errorf("decode failed, expect %s:%d, but got %d", "Ver", p.Ver, res.Ver)
	}
	if p.Op != res.Op {
		t.Errorf("decode failed, expect %s:%d, but got %d", "Op", p.Op, res.Op)
	}
	if p.SeqId != res.SeqId {
		t.Errorf("decode failed, expect %s:%d, but got %d", "SeqId", p.SeqId, res.SeqId)
	}
	if bytes.Compare(p.Body, res.Body) != 0 {
		t.Errorf("decode failed, expect %s:%s, but got %s", "Body", string(p.Body), string(res.Body))
	}
}
