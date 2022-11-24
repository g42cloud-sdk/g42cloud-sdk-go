package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
