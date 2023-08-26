package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateOriginHostResponse struct {
	OriginHost     *DomainOriginHost `json:"origin_host,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateOriginHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOriginHostResponse struct{}"
	}

	return strings.Join([]string{"UpdateOriginHostResponse", string(data)}, " ")
}
