package data

type Header struct {
	HeaderId string `json:"header_id"`
	Header   string `json:"header"`
}

func NewHeader() *Header {
	return &Header{}
}
