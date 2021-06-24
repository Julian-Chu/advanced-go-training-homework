package week9

import "encoding/binary"

type Protocol struct {
	PackLen   uint32
	HeaderLen uint16
	Ver       uint16
	Op        uint32
	SeqId     uint32
	Body      []byte
}

const (
	packLenSize   = 4
	headerLenSize = 2
	verSize       = 2
	opSize        = 4
	SeqIdSize     = 4
)

//func NewProtocol(ver uint16, op uint32, seqId uint32, message string) *Protocol {
//	return &Protocol{Ver: ver, Op: op, SeqId: seqId, Body: []byte(message)}
//}

func Encode(p Protocol) []byte {
	headerSize := packLenSize + headerLenSize + verSize + opSize + SeqIdSize
	bufSize := headerSize
	if p.Body != nil {
		bufSize += len(p.Body)
	}
	buf := make([]byte, bufSize)
	offset := 0
	binary.BigEndian.PutUint32(buf[offset:], p.PackLen)
	offset += packLenSize
	binary.BigEndian.PutUint16(buf[offset:], p.HeaderLen)
	offset += headerLenSize
	binary.BigEndian.PutUint16(buf[offset:], p.Ver)
	offset += verSize
	binary.BigEndian.PutUint32(buf[offset:], p.Op)
	offset += opSize
	binary.BigEndian.PutUint32(buf[offset:], p.SeqId)
	if p.Body != nil {
		offset += SeqIdSize
		copy(buf[offset:], p.Body)
	}
	return buf
}

func Decode(buf []byte) Protocol {
	p := Protocol{}
	start := 0
	end := packLenSize
	p.PackLen = binary.BigEndian.Uint32(buf[start:end])
	start = end
	end += headerLenSize
	p.HeaderLen = binary.BigEndian.Uint16(buf[start:end])
	start = end
	end += verSize
	p.Ver = binary.BigEndian.Uint16(buf[start:end])
	start = end
	end += opSize
	p.Op = binary.BigEndian.Uint32(buf[start:end])
	start = end
	end += SeqIdSize
	p.SeqId = binary.BigEndian.Uint32(buf[start:end])
	bodySize := p.PackLen - uint32(p.HeaderLen)
	if bodySize > 0 {
		p.Body = make([]byte, bodySize)
		copy(p.Body, buf[end:])
	}
	return p
}
