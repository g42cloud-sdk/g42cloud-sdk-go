package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSqlserverDatabaseRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *SqlserverDatabaseForCreation `json:"body,omitempty"`
}

func (o CreateSqlserverDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlserverDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlserverDatabaseRequest", string(data)}, " ")
}
