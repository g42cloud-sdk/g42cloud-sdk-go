package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type PostPaidServerEipExtendParam struct {
	ChargingMode *PostPaidServerEipExtendParamChargingMode `json:"chargingMode,omitempty"`
}

func (o PostPaidServerEipExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerEipExtendParam struct{}"
	}

	return strings.Join([]string{"PostPaidServerEipExtendParam", string(data)}, " ")
}

type PostPaidServerEipExtendParamChargingMode struct {
	value string
}

type PostPaidServerEipExtendParamChargingModeEnum struct {
	PRE_PAID  PostPaidServerEipExtendParamChargingMode
	POST_PAID PostPaidServerEipExtendParamChargingMode
}

func GetPostPaidServerEipExtendParamChargingModeEnum() PostPaidServerEipExtendParamChargingModeEnum {
	return PostPaidServerEipExtendParamChargingModeEnum{
		PRE_PAID: PostPaidServerEipExtendParamChargingMode{
			value: "prePaid",
		},
		POST_PAID: PostPaidServerEipExtendParamChargingMode{
			value: "postPaid",
		},
	}
}

func (c PostPaidServerEipExtendParamChargingMode) Value() string {
	return c.value
}

func (c PostPaidServerEipExtendParamChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerEipExtendParamChargingMode) UnmarshalJSON(b []byte) error {
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
