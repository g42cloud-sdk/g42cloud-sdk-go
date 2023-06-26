package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowAgenciesTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`
}

func (o ShowAgenciesTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgenciesTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowAgenciesTaskRequest", string(data)}, " ")
}
