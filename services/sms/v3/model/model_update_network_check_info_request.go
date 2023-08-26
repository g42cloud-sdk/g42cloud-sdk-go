package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateNetworkCheckInfoRequest struct {
	TaskId string `json:"task_id"`

	Body *NetworkCheckInfoRequestBody `json:"body,omitempty"`
}

func (o UpdateNetworkCheckInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkCheckInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateNetworkCheckInfoRequest", string(data)}, " ")
}
