package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTaskResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskResponse", string(data)}, " ")
}
