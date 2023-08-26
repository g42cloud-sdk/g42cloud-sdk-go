package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaskResponse", string(data)}, " ")
}
