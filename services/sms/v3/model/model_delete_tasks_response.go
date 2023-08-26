package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTasksResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTasksResponse struct{}"
	}

	return strings.Join([]string{"DeleteTasksResponse", string(data)}, " ")
}
