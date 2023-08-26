package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTranscodingTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTranscodingTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTranscodingTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingTaskResponse", string(data)}, " ")
}
