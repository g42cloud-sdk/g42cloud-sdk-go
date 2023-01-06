package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDomainFullConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDomainFullConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainFullConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainFullConfigResponse", string(data)}, " ")
}
