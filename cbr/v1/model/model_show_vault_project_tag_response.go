package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVaultProjectTagResponse struct {
	Tags           *[]TagsResp `json:"tags,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowVaultProjectTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultProjectTagResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultProjectTagResponse", string(data)}, " ")
}
