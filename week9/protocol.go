package week9

type Protocol struct {
	PackLen   int32
	HeaderLen int16
	Ver       int16
	Op        int32
	SeqId     int32
	Body      []byte
}
