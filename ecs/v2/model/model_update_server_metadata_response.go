package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateServerMetadataResponse struct {
	Metadata       map[string]string `json:"metadata,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateServerMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerMetadataResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerMetadataResponse", string(data)}, " ")
}
