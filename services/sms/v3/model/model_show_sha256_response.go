package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowSha256Response struct {
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSha256Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSha256Response struct{}"
	}

	return strings.Join([]string{"ShowSha256Response", string(data)}, " ")
}
