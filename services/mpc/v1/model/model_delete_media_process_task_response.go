package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteMediaProcessTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMediaProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMediaProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteMediaProcessTaskResponse", string(data)}, " ")
}
