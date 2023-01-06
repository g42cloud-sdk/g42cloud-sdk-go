package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteResourceGroupRequest struct {
	ContentType string `json:"Content-Type"`

	GroupId string `json:"group_id"`
}

func (o DeleteResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupRequest", string(data)}, " ")
}
