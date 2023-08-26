package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/sdktime"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SecurityGroup struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	ProjectId string `json:"project_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}
