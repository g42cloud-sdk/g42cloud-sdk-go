package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTranscodeDetailRequest struct {
	TaskId *[]string `json:"task_id,omitempty"`
}

func (o ListTranscodeDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeDetailRequest struct{}"
	}

	return strings.Join([]string{"ListTranscodeDetailRequest", string(data)}, " ")
}
