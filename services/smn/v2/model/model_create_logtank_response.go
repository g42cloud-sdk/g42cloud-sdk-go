package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateLogtankResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogtankResponse struct{}"
	}

	return strings.Join([]string{"CreateLogtankResponse", string(data)}, " ")
}
