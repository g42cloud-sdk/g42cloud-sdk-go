package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTasksRequest struct {
	Body *DeleteTasksReq `json:"body,omitempty"`
}

func (o DeleteTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTasksRequest struct{}"
	}

	return strings.Join([]string{"DeleteTasksRequest", string(data)}, " ")
}
