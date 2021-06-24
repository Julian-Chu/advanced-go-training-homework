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

func NewProtocol(ver uint16, op uint32, seqId uint32, message string) *Protocol {

	return &Protocol{Ver: ver, Op: op, SeqId: seqId, Body: []byte(message)}
}

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
