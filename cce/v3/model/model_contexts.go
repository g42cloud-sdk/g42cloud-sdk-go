package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Contexts struct {
	Name *string `json:"name,omitempty"`

	Context *Context `json:"context,omitempty"`
}

func (o Contexts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Contexts struct{}"
	}

	return strings.Join([]string{"Contexts", string(data)}, " ")
}
