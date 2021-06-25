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

func NewProto(ver uint16, op uint32, seqId uint32, body []byte) *Protocol {
	headerSize := packLenSize + headerLenSize + verSize + opSize + seqIdSize
	packLen := headerSize
	if body != nil {
		packLen += len(body)
	}
	return &Protocol{PackLen: uint32(packLen), HeaderLen: uint16(headerSize), Ver: ver, Op: op, SeqId: seqId, Body: body}
}

const (
	packLenSize   = 4
	headerLenSize = 2
	verSize       = 2
	opSize        = 4
	seqIdSize     = 4
	headerSize    = packLenSize + headerLenSize + verSize + opSize + seqIdSize
)

const (
	packLenOffset   = 0
	headerLenOffset = packLenOffset + packLenSize
	verOffset       = headerLenOffset + headerLenSize
	opOffset        = verOffset + verSize
	seqIdOffset     = opOffset + opSize
	bodyOffset      = seqIdOffset + seqIdSize
)

func (p Protocol) Encode() []byte {
	bufSize := headerSize
	if p.Body != nil {
		bufSize += len(p.Body)
	}
	buf := make([]byte, bufSize)
	binary.BigEndian.PutUint32(buf[packLenOffset:], p.PackLen)
	binary.BigEndian.PutUint16(buf[headerLenOffset:], p.HeaderLen)
	binary.BigEndian.PutUint16(buf[verOffset:], p.Ver)
	binary.BigEndian.PutUint32(buf[opOffset:], p.Op)
	binary.BigEndian.PutUint32(buf[seqIdOffset:], p.SeqId)
	if p.Body != nil {
		copy(buf[bodyOffset:], p.Body)
	}
	return buf
}

func (p *Protocol) Decode(buf []byte) {
	p.PackLen = binary.BigEndian.Uint32(buf[packLenOffset:headerLenOffset])
	p.HeaderLen = binary.BigEndian.Uint16(buf[headerLenOffset:verOffset])
	p.Ver = binary.BigEndian.Uint16(buf[verOffset:opOffset])
	p.Op = binary.BigEndian.Uint32(buf[opOffset:seqIdOffset])
	p.SeqId = binary.BigEndian.Uint32(buf[seqIdOffset:bodyOffset])
	bodySize := p.PackLen - uint32(p.HeaderLen)
	if bodySize > 0 {
		p.Body = make([]byte, bodySize)
		copy(p.Body, buf[bodyOffset:])
	}
}
