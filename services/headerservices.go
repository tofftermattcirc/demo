package services

import (
	"circadence/demo/data"
	"encoding/json"
	"io"
)

type HeaderService struct {
}

func NewHeaderService() *HeaderService {
	return &HeaderService{}
}

/** POST servicer */
func (ref *HeaderService) RegisterHeaderServicer(rc io.ReadCloser) []byte {
	header := data.NewHeader()

	/** Decode incoming Header JSON */
	err := json.NewDecoder(rc).Decode(&header)
	IfErrorHandle(err)

	data, err := json.Marshal(header)
	IfErrorHandle(err)

	return data

}
