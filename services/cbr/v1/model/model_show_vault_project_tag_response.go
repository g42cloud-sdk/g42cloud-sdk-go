package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
