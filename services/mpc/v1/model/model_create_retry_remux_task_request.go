package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateRetryRemuxTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *RemuxRetryReq `json:"body,omitempty"`
}

func (o CreateRetryRemuxTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryRemuxTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateRetryRemuxTaskRequest", string(data)}, " ")
}
