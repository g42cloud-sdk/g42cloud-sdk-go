package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type HeaderBody struct {
	Headers *HeaderMap `json:"headers,omitempty"`
}

func (o HeaderBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeaderBody struct{}"
	}

	return strings.Join([]string{"HeaderBody", string(data)}, " ")
}
