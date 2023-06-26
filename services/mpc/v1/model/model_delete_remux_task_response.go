package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteRemuxTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRemuxTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteRemuxTaskResponse", string(data)}, " ")
}
