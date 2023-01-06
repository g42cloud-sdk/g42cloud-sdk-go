package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateDomainOriginResponse struct {
	Origin         *ResourceBody `json:"origin,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateDomainOriginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainOriginResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainOriginResponse", string(data)}, " ")
}
