package util

import (
	"encoding/json"
	"github.com/HyperService-Consortium/NSB/application/response"
)

func MustUnmarshal(data []byte, load interface{}) {
	err := json.Unmarshal(data, &load)
	if err != nil {
		panic(response.DecodeJsonError(err))
	}
}
