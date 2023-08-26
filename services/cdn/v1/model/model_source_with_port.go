package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type SourceWithPort struct {
	IpOrDomain string `json:"ip_or_domain"`

	OriginType SourceWithPortOriginType `json:"origin_type"`

	ActiveStandby int32 `json:"active_standby"`

	EnableObsWebHosting *int32 `json:"enable_obs_web_hosting,omitempty"`

	HttpPort *int32 `json:"http_port,omitempty"`

	HttpsPort *int32 `json:"https_port,omitempty"`
}

func (o SourceWithPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceWithPort struct{}"
	}

	return strings.Join([]string{"SourceWithPort", string(data)}, " ")
}

type SourceWithPortOriginType struct {
	value string
}

type SourceWithPortOriginTypeEnum struct {
	IPADDR     SourceWithPortOriginType
	DOMAIN     SourceWithPortOriginType
	OBS_BUCKET SourceWithPortOriginType
}

func GetSourceWithPortOriginTypeEnum() SourceWithPortOriginTypeEnum {
	return SourceWithPortOriginTypeEnum{
		IPADDR: SourceWithPortOriginType{
			value: "ipaddr",
		},
		DOMAIN: SourceWithPortOriginType{
			value: "domain",
		},
		OBS_BUCKET: SourceWithPortOriginType{
			value: "obs_bucket",
		},
	}
}

func (c SourceWithPortOriginType) Value() string {
	return c.value
}

func (c SourceWithPortOriginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceWithPortOriginType) UnmarshalJSON(b []byte) error {
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
