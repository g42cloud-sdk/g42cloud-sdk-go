package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePostgresqlDbUserRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *PostgresqlUserForCreation `json:"body,omitempty"`
}

func (o CreatePostgresqlDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDbUserRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDbUserRequest", string(data)}, " ")
}
