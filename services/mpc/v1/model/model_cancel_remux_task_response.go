package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CancelRemuxTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelRemuxTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelRemuxTaskResponse", string(data)}, " ")
}
