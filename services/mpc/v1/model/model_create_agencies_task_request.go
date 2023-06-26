package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAgenciesTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *AgenciesTaskReq `json:"body,omitempty"`
}

func (o CreateAgenciesTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgenciesTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAgenciesTaskRequest", string(data)}, " ")
}
