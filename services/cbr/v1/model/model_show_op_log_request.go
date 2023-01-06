package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowOpLogRequest struct {
	OperationLogId string `json:"operation_log_id"`
}

func (o ShowOpLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOpLogRequest struct{}"
	}

	return strings.Join([]string{"ShowOpLogRequest", string(data)}, " ")
}
