package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteEncryptTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEncryptTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEncryptTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteEncryptTaskResponse", string(data)}, " ")
}
