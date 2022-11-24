package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Domains struct {
	Id *string `json:"id,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	BusinessType *string `json:"business_type,omitempty"`

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

	AutoRefreshPreheat *int32 `json:"auto_refresh_preheat,omitempty"`

	ServiceArea *DomainsServiceArea `json:"service_area,omitempty"`

	RangeStatus *string `json:"range_status,omitempty"`

	FollowStatus *string `json:"follow_status,omitempty"`

	OriginStatus *string `json:"origin_status,omitempty"`

	BannedReason *string `json:"banned_reason,omitempty"`

	LockedReason *string `json:"locked_reason,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o Domains) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Domains struct{}"
	}

	return strings.Join([]string{"Domains", string(data)}, " ")
}

type DomainsServiceArea struct {
	value string
}

type DomainsServiceAreaEnum struct {
	MAINLAND_CHINA         DomainsServiceArea
	OUTSIDE_MAINLAND_CHINA DomainsServiceArea
	GLOBAL                 DomainsServiceArea
}

func GetDomainsServiceAreaEnum() DomainsServiceAreaEnum {
	return DomainsServiceAreaEnum{
		MAINLAND_CHINA: DomainsServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: DomainsServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: DomainsServiceArea{
			value: "global",
		},
	}
}

func (c DomainsServiceArea) Value() string {
	return c.value
}

func (c DomainsServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainsServiceArea) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
