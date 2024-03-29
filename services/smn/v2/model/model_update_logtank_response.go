package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateLogtankResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Logtank        *LogtankItem `json:"logtank,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateLogtankResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogtankResponse", string(data)}, " ")
}
