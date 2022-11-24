package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserForCreation struct {
	Name string `json:"name"`

	Password string `json:"password"`
}

func (o UserForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserForCreation struct{}"
	}

	return strings.Join([]string{"UserForCreation", string(data)}, " ")
}
