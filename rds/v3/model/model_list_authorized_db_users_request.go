package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAuthorizedDbUsersRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	DbName string `json:"db-name"`

	Page int32 `json:"page"`

	Limit int32 `json:"limit"`
}

func (o ListAuthorizedDbUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizedDbUsersRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizedDbUsersRequest", string(data)}, " ")
}
