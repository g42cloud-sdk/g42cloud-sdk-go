package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CollectLogRequest struct {
	TaskId string `json:"task_id"`

	Body *UploadLogRequestBody `json:"body,omitempty"`
}

func (o CollectLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectLogRequest struct{}"
	}

	return strings.Join([]string{"CollectLogRequest", string(data)}, " ")
}
