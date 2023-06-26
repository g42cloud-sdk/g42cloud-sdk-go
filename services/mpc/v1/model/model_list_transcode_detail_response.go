package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTranscodeDetailResponse struct {
	TaskArray      *[]TaskDetailInfo `json:"task_array,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTranscodeDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeDetailResponse struct{}"
	}

	return strings.Join([]string{"ListTranscodeDetailResponse", string(data)}, " ")
}
