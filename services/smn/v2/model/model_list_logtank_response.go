package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListLogtankResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Logtanks       *[]LogtankItem `json:"logtanks,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtankResponse struct{}"
	}

	return strings.Join([]string{"ListLogtankResponse", string(data)}, " ")
}
