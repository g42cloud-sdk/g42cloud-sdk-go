package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Sources struct {
	IpOrDomain string `json:"ip_or_domain"`

	OriginType SourcesOriginType `json:"origin_type"`

	ActiveStandby int32 `json:"active_standby"`

	EnableObsWebHosting *int32 `json:"enable_obs_web_hosting,omitempty"`
}

func (o Sources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Sources struct{}"
	}

	return strings.Join([]string{"Sources", string(data)}, " ")
}

type SourcesOriginType struct {
	value string
}

type SourcesOriginTypeEnum struct {
	IPADDR     SourcesOriginType
	DOMAIN     SourcesOriginType
	OBS_BUCKET SourcesOriginType
}

func GetSourcesOriginTypeEnum() SourcesOriginTypeEnum {
	return SourcesOriginTypeEnum{
		IPADDR: SourcesOriginType{
			value: "ipaddr",
		},
		DOMAIN: SourcesOriginType{
			value: "domain",
		},
		OBS_BUCKET: SourcesOriginType{
			value: "obs_bucket",
		},
	}
}

func (c SourcesOriginType) Value() string {
	return c.value
}

func (c SourcesOriginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourcesOriginType) UnmarshalJSON(b []byte) error {
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
