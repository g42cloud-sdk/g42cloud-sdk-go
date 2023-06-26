package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateRemuxTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *CreateRemuxTaskReq `json:"body,omitempty"`
}

func (o CreateRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateRemuxTaskRequest", string(data)}, " ")
}
