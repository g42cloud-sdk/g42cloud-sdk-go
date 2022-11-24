package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteSqlserverDatabaseExRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	DbName string `json:"db_name"`

	Body *DropDatabaseV3Req `json:"body,omitempty"`
}

func (o DeleteSqlserverDatabaseExRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlserverDatabaseExRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlserverDatabaseExRequest", string(data)}, " ")
}
