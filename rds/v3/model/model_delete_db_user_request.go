package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDbUserRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	UserName string `json:"user_name"`
}

func (o DeleteDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteDbUserRequest", string(data)}, " ")
}
