package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowBlackWhiteListResponse struct {
	Type *int32 `json:"type,omitempty"`

	IpList         *[]string `json:"ip_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBlackWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlackWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ShowBlackWhiteListResponse", string(data)}, " ")
}
