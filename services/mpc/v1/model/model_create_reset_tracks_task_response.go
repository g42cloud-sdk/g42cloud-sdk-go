package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResetTracksTaskResponse struct {
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResetTracksTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResetTracksTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateResetTracksTaskResponse", string(data)}, " ")
}
