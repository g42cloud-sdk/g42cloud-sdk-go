package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ChargeInfoResponse struct {
	ChargeMode ChargeInfoResponseChargeMode `json:"charge_mode"`
}

func (o ChargeInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfoResponse struct{}"
	}

	return strings.Join([]string{"ChargeInfoResponse", string(data)}, " ")
}

type ChargeInfoResponseChargeMode struct {
	value string
}

type ChargeInfoResponseChargeModeEnum struct {
	PRE_PAID  ChargeInfoResponseChargeMode
	POST_PAID ChargeInfoResponseChargeMode
}

func GetChargeInfoResponseChargeModeEnum() ChargeInfoResponseChargeModeEnum {
	return ChargeInfoResponseChargeModeEnum{
		PRE_PAID: ChargeInfoResponseChargeMode{
			value: "prePaid",
		},
		POST_PAID: ChargeInfoResponseChargeMode{
			value: "postPaid",
		},
	}
}

func (c ChargeInfoResponseChargeMode) Value() string {
	return c.value
}

func (c ChargeInfoResponseChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChargeInfoResponseChargeMode) UnmarshalJSON(b []byte) error {
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
