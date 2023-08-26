package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/sdktime"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Vpc struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Cidr string `json:"cidr"`

	ExtendCidrs []string `json:"extend_cidrs"`

	Status string `json:"status"`

	ProjectId string `json:"project_id"`

	EnterpriseProjectId string `json:"enterprise_project_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	CloudResources []CloudResource `json:"cloud_resources"`

	Tags []Tag `json:"tags"`
}

func (o Vpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vpc struct{}"
	}

	return strings.Join([]string{"Vpc", string(data)}, " ")
}
