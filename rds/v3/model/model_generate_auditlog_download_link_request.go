package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenerateAuditlogDownloadLinkRequest struct {
	Ids []string `json:"ids"`
}

func (o GenerateAuditlogDownloadLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateAuditlogDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"GenerateAuditlogDownloadLinkRequest", string(data)}, " ")
}
