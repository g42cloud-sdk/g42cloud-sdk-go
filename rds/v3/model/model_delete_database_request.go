package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDatabaseRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	DbName string `json:"db_name"`
}

func (o DeleteDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRequest", string(data)}, " ")
}
