package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListJobInfoRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	Id string `json:"id"`
}

func (o ListJobInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobInfoRequest struct{}"
	}

	return strings.Join([]string{"ListJobInfoRequest", string(data)}, " ")
}
