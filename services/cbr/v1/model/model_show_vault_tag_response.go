package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVaultTagResponse struct {
	Tags *[]Tag `json:"tags,omitempty"`

	SysTags        *[]SysTag `json:"sys_tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowVaultTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultTagResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultTagResponse", string(data)}, " ")
}
