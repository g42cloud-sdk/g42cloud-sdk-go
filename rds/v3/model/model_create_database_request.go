package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDatabaseRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *DatabaseForCreation `json:"body,omitempty"`
}

func (o CreateDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequest", string(data)}, " ")
}
