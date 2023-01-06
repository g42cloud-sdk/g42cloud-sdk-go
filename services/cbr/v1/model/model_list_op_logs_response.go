package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListOpLogsResponse struct {
	OperationLogs *[]OperationLog `json:"operation_logs,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOpLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpLogsResponse struct{}"
	}

	return strings.Join([]string{"ListOpLogsResponse", string(data)}, " ")
}
