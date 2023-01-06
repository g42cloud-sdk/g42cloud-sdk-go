package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowOriginHostResponse struct {
	OriginHost     *DomainOriginHost `json:"origin_host,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowOriginHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOriginHostResponse struct{}"
	}

	return strings.Join([]string{"ShowOriginHostResponse", string(data)}, " ")
}
