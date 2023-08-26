package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteThumbnailsTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteThumbnailsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteThumbnailsTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteThumbnailsTaskRequest", string(data)}, " ")
}
