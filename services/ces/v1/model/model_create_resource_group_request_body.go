package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResourceGroupRequestBody struct {
	GroupName string `json:"group_name"`

	Resources []CreateResourceGroup `json:"resources"`
}

func (o CreateResourceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupRequestBody", string(data)}, " ")
}
