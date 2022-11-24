package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PostPaidServerEipBandwidth struct {
	Size *int32 `json:"size,omitempty"`

	Sharetype PostPaidServerEipBandwidthSharetype `json:"sharetype"`

	Chargemode *string `json:"chargemode,omitempty"`

	Id *string `json:"id,omitempty"`
}

func (o PostPaidServerEipBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerEipBandwidth struct{}"
	}

	return strings.Join([]string{"PostPaidServerEipBandwidth", string(data)}, " ")
}

type PostPaidServerEipBandwidthSharetype struct {
	value string
}

type PostPaidServerEipBandwidthSharetypeEnum struct {
	PER   PostPaidServerEipBandwidthSharetype
	WHOLE PostPaidServerEipBandwidthSharetype
}

func GetPostPaidServerEipBandwidthSharetypeEnum() PostPaidServerEipBandwidthSharetypeEnum {
	return PostPaidServerEipBandwidthSharetypeEnum{
		PER: PostPaidServerEipBandwidthSharetype{
			value: "PER",
		},
		WHOLE: PostPaidServerEipBandwidthSharetype{
			value: "WHOLE",
		},
	}
}

func (c PostPaidServerEipBandwidthSharetype) Value() string {
	return c.value
}

func (c PostPaidServerEipBandwidthSharetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerEipBandwidthSharetype) UnmarshalJSON(b []byte) error {
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
