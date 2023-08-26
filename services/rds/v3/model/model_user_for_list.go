package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UserForList struct {
	Name string `json:"name"`
}

func (o UserForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserForList struct{}"
	}

	return strings.Join([]string{"UserForList", string(data)}, " ")
}
