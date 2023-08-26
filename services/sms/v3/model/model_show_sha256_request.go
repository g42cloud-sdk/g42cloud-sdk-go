package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSha256Request struct {
	Key string `json:"key"`
}

func (o ShowSha256Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSha256Request struct{}"
	}

	return strings.Join([]string{"ShowSha256Request", string(data)}, " ")
}
