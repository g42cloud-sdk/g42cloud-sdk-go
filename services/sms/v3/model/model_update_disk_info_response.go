package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDiskInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDiskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDiskInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDiskInfoResponse", string(data)}, " ")
}
