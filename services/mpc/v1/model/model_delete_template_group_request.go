package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTemplateGroupRequest struct {
	GroupId string `json:"group_id"`
}

func (o DeleteTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupRequest", string(data)}, " ")
}
