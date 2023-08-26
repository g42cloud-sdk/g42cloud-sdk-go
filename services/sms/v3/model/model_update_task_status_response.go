package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTaskStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResponse", string(data)}, " ")
}
