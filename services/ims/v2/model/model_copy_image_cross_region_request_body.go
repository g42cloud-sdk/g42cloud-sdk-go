package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CopyImageCrossRegionRequestBody struct {
	AgencyName string `json:"agency_name"`

	Description *string `json:"description,omitempty"`

	Name string `json:"name"`

	ProjectName string `json:"project_name"`

	Region string `json:"region"`
}

func (o CopyImageCrossRegionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionRequestBody struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionRequestBody", string(data)}, " ")
}
