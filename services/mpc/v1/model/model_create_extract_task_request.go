package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateExtractTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *CreateExtractTaskReq `json:"body,omitempty"`
}

func (o CreateExtractTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExtractTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateExtractTaskRequest", string(data)}, " ")
}
