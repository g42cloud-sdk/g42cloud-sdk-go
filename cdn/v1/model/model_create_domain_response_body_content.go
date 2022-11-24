package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDomainResponseBodyContent struct {
	Id *string `json:"id,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	BusinessType *string `json:"business_type,omitempty"`

	ServiceArea *string `json:"service_area,omitempty"`

	UserDomainId *string `json:"user_domain_id,omitempty"`

	DomainStatus *string `json:"domain_status,omitempty"`

	Cname *string `json:"cname,omitempty"`

	Sources *[]Sources `json:"sources,omitempty"`

	DomainOriginHost *DomainOriginHost `json:"domain_origin_host,omitempty"`

	HttpsStatus *int32 `json:"https_status,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	ModifyTime *int64 `json:"modify_time,omitempty"`

	Disabled *int32 `json:"disabled,omitempty"`

	Locked *int32 `json:"locked,omitempty"`

	RangeStatus *string `json:"range_status,omitempty"`

	FollowStatus *string `json:"follow_status,omitempty"`

	OriginStatus *string `json:"origin_status,omitempty"`

	AutoRefreshPreheat *int32 `json:"auto_refresh_preheat,omitempty"`
}

func (o CreateDomainResponseBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainResponseBodyContent struct{}"
	}

	return strings.Join([]string{"CreateDomainResponseBodyContent", string(data)}, " ")
}
