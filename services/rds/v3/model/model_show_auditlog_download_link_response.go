package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
