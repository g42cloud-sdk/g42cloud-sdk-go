package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResetTracksTaskRequest struct {
	Body *CreateResetTracksReq `json:"body,omitempty"`
}

func (o CreateResetTracksTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResetTracksTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateResetTracksTaskRequest", string(data)}, " ")
}
