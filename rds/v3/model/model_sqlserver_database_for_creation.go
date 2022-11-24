package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlserverDatabaseForCreation struct {
	Name string `json:"name"`
}

func (o SqlserverDatabaseForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlserverDatabaseForCreation struct{}"
	}

	return strings.Join([]string{"SqlserverDatabaseForCreation", string(data)}, " ")
}
