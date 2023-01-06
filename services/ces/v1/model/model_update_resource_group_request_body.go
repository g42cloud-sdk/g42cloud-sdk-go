package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateResourceGroupRequestBody struct {
	GroupName string `json:"group_name"`

	Resources []CreateResourceGroup `json:"resources"`
}

func (o UpdateResourceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupRequestBody", string(data)}, " ")
}
