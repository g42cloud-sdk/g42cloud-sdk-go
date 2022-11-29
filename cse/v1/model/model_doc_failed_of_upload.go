package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DocFailedOfUpload struct {
	Key *string `json:"key,omitempty"`

	Labels *interface{} `json:"labels,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o DocFailedOfUpload) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocFailedOfUpload struct{}"
	}

	return strings.Join([]string{"DocFailedOfUpload", string(data)}, " ")
}
