package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDomainsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Domains        *[]Domains `json:"domains,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainsResponse", string(data)}, " ")
}
