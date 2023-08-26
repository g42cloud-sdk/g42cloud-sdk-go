package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchRebootServersRequest struct {
	Body *BatchRebootServersRequestBody `json:"body,omitempty"`
}

func (o BatchRebootServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootServersRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootServersRequest", string(data)}, " ")
}
