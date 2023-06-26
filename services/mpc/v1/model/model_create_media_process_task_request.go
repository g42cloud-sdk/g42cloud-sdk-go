package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMediaProcessTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *CreateMediaProcessReq `json:"body,omitempty"`
}

func (o CreateMediaProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMediaProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateMediaProcessTaskRequest", string(data)}, " ")
}
