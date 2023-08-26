package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupResponse", string(data)}, " ")
}
