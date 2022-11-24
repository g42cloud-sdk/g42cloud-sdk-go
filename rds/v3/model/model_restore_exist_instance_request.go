package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreExistInstanceRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	Body *RestoreExistingInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreExistInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreExistInstanceRequest", string(data)}, " ")
}
