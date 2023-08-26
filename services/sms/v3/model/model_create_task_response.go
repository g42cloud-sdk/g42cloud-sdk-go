package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateTaskResponse struct {
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateTaskResponse", string(data)}, " ")
}
