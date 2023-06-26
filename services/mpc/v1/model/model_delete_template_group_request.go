package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTemplateGroupRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	GroupId string `json:"group_id"`
}

func (o DeleteTemplateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupRequest", string(data)}, " ")
}
