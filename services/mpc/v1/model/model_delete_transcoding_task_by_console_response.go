package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTranscodingTaskByConsoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTranscodingTaskByConsoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTranscodingTaskByConsoleResponse struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingTaskByConsoleResponse", string(data)}, " ")
}
