package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListQualityEnhanceDefaultTemplateResponse struct {
	TaskArray *[]QualityEnhanceTemplateInfo `json:"task_array,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListQualityEnhanceDefaultTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityEnhanceDefaultTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListQualityEnhanceDefaultTemplateResponse", string(data)}, " ")
}
