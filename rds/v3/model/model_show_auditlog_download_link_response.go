package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAuditlogDownloadLinkResponse struct {
	Links          *[]string `json:"links,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowAuditlogDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditlogDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditlogDownloadLinkResponse", string(data)}, " ")
}
