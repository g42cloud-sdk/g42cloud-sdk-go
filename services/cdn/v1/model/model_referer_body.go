package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RefererBody struct {
	Referer *Referer `json:"referer"`
}

func (o RefererBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefererBody struct{}"
	}

	return strings.Join([]string{"RefererBody", string(data)}, " ")
}
