package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlserverUserForCreation struct {
	Name string `json:"name"`

	Password string `json:"password"`
}

func (o SqlserverUserForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlserverUserForCreation struct{}"
	}

	return strings.Join([]string{"SqlserverUserForCreation", string(data)}, " ")
}
