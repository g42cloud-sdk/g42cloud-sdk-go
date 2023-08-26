package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RestoreBackupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreBackupResponse struct{}"
	}

	return strings.Join([]string{"RestoreBackupResponse", string(data)}, " ")
}
