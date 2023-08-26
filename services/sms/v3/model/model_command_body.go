package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CommandBody struct {
	CommandName string `json:"command_name"`

	Result string `json:"result"`

	ResultDetail *interface{} `json:"result_detail"`
}

func (o CommandBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommandBody struct{}"
	}

	return strings.Join([]string{"CommandBody", string(data)}, " ")
}
