package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteManualBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManualBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteManualBackupResponse", string(data)}, " ")
}
