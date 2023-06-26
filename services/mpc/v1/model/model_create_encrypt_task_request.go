package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEncryptTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *CreateEncryptReq `json:"body,omitempty"`
}

func (o CreateEncryptTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateEncryptTaskRequest", string(data)}, " ")
}
