package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TopicAttribute struct {
	Version string `json:"Version"`

	Id string `json:"Id"`

	Statement []Statement `json:"Statement"`
}

func (o TopicAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicAttribute struct{}"
	}

	return strings.Join([]string{"TopicAttribute", string(data)}, " ")
}
