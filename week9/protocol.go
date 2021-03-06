package week9

import (
	"bufio"
	"encoding/binary"
	"net"
)

type Protocol struct {
	Ver   uint16
	Op    uint32
	SeqId uint32
	Body  []byte
}

func NewProto(ver uint16, op uint32, seqId uint32, body []byte) *Protocol {
	return &Protocol{Ver: ver, Op: op, SeqId: seqId, Body: body}
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
	var packLen uint32 = headerSize
	if p.Body != nil {
		packLen += uint32(len(p.Body))
	}
	buf := make([]byte, packLen)
	binary.BigEndian.PutUint32(buf[packLenOffset:], packLen)
	binary.BigEndian.PutUint16(buf[headerLenOffset:], headerSize)
	binary.BigEndian.PutUint16(buf[verOffset:], p.Ver)
	binary.BigEndian.PutUint32(buf[opOffset:], p.Op)
	binary.BigEndian.PutUint32(buf[seqIdOffset:], p.SeqId)
	if p.Body != nil {
		copy(buf[bodyOffset:], p.Body)
	}
	return buf
}

func (p *Protocol) Decode(buf []byte) {
	packLen := binary.BigEndian.Uint32(buf[packLenOffset:headerLenOffset])
	p.Ver = binary.BigEndian.Uint16(buf[verOffset:opOffset])
	p.Op = binary.BigEndian.Uint32(buf[opOffset:seqIdOffset])
	p.SeqId = binary.BigEndian.Uint32(buf[seqIdOffset:bodyOffset])
	bodySize := packLen - uint32(headerSize)
	if bodySize > 0 {
		p.Body = make([]byte, bodySize)
		copy(p.Body, buf[bodyOffset:])
	}
}

type Reader struct {
	conn net.Conn
}

func NewReader(conn net.Conn) (*Reader, func() error) {
	return &Reader{conn: conn}, conn.Close
}

func (r *Reader) Decode() (*Protocol, error) {
	p := &Protocol{}
	reader := bufio.NewReader(r.conn)
	var err error
	Read := func(buf []byte) {
		if err != nil {
			return
		}
		_, err = reader.Read(buf)
	}
	buf := make([]byte, packLenSize)
	Read(buf)
	packLen := binary.BigEndian.Uint32(buf)

	buf = make([]byte, headerLenSize)
	Read(buf)
	//headerLen := binary.BigEndian.Uint16(buf)
	buf = make([]byte, verSize)
	Read(buf)
	p.Ver = binary.BigEndian.Uint16(buf)

	buf = make([]byte, opSize)
	Read(buf)
	p.Op = binary.BigEndian.Uint32(buf)

	buf = make([]byte, seqIdSize)
	Read(buf)
	p.SeqId = binary.BigEndian.Uint32(buf)
	bodySize := packLen - uint32(headerSize)
	if bodySize > 0 {
		p.Body = make([]byte, int(bodySize))
		Read(p.Body)
	}
	return p, err
}
