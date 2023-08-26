package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
