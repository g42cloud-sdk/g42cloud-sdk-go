package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResetTracksTaskRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *CreateResetTracksReq `json:"body,omitempty"`
}

func (o CreateResetTracksTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResetTracksTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateResetTracksTaskRequest", string(data)}, " ")
}
