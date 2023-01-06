package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowIpInfoResponse struct {
	CdnIps         *[]CdnIps `json:"cdn_ips,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowIpInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowIpInfoResponse", string(data)}, " ")
}
