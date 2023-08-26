package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowHistoryTaskDetailsResponse struct {
	Id *string `json:"id,omitempty"`

	TaskType *string `json:"task_type,omitempty"`

	Status *string `json:"status,omitempty"`

	Urls *[]UrlObject `json:"urls,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Processing *int32 `json:"processing,omitempty"`

	Succeed *int32 `json:"succeed,omitempty"`

	Failed *int32 `json:"failed,omitempty"`

	Total *int32 `json:"total,omitempty"`

	FileType       *string `json:"file_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHistoryTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryTaskDetailsResponse", string(data)}, " ")
}
