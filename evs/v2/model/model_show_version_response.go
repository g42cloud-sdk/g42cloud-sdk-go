package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVersionResponse struct {
	Versions       *[]Versions `json:"versions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowVersionResponse", string(data)}, " ")
}
