package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type BssParamForCreateVolume struct {
	ChargingMode *BssParamForCreateVolumeChargingMode `json:"chargingMode,omitempty"`

	IsAutoPay *BssParamForCreateVolumeIsAutoPay `json:"isAutoPay,omitempty"`

	IsAutoRenew *BssParamForCreateVolumeIsAutoRenew `json:"isAutoRenew,omitempty"`

	PeriodNum *int32 `json:"periodNum,omitempty"`

	PeriodType *BssParamForCreateVolumePeriodType `json:"periodType,omitempty"`
}

func (o BssParamForCreateVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParamForCreateVolume struct{}"
	}

	return strings.Join([]string{"BssParamForCreateVolume", string(data)}, " ")
}

type BssParamForCreateVolumeChargingMode struct {
	value string
}

type BssParamForCreateVolumeChargingModeEnum struct {
	POST_PAID BssParamForCreateVolumeChargingMode
	PRE_PAID  BssParamForCreateVolumeChargingMode
}

func GetBssParamForCreateVolumeChargingModeEnum() BssParamForCreateVolumeChargingModeEnum {
	return BssParamForCreateVolumeChargingModeEnum{
		POST_PAID: BssParamForCreateVolumeChargingMode{
			value: "postPaid",
		},
		PRE_PAID: BssParamForCreateVolumeChargingMode{
			value: "prePaid",
		},
	}
}

func (c BssParamForCreateVolumeChargingMode) Value() string {
	return c.value
}

func (c BssParamForCreateVolumeChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForCreateVolumeChargingMode) UnmarshalJSON(b []byte) error {
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

type BssParamForCreateVolumeIsAutoPay struct {
	value string
}

type BssParamForCreateVolumeIsAutoPayEnum struct {
	TRUE  BssParamForCreateVolumeIsAutoPay
	FALSE BssParamForCreateVolumeIsAutoPay
}

func GetBssParamForCreateVolumeIsAutoPayEnum() BssParamForCreateVolumeIsAutoPayEnum {
	return BssParamForCreateVolumeIsAutoPayEnum{
		TRUE: BssParamForCreateVolumeIsAutoPay{
			value: "true",
		},
		FALSE: BssParamForCreateVolumeIsAutoPay{
			value: "false",
		},
	}
}

func (c BssParamForCreateVolumeIsAutoPay) Value() string {
	return c.value
}

func (c BssParamForCreateVolumeIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForCreateVolumeIsAutoPay) UnmarshalJSON(b []byte) error {
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

type BssParamForCreateVolumeIsAutoRenew struct {
	value string
}

type BssParamForCreateVolumeIsAutoRenewEnum struct {
	TRUE  BssParamForCreateVolumeIsAutoRenew
	FALSE BssParamForCreateVolumeIsAutoRenew
}

func GetBssParamForCreateVolumeIsAutoRenewEnum() BssParamForCreateVolumeIsAutoRenewEnum {
	return BssParamForCreateVolumeIsAutoRenewEnum{
		TRUE: BssParamForCreateVolumeIsAutoRenew{
			value: "true",
		},
		FALSE: BssParamForCreateVolumeIsAutoRenew{
			value: "false",
		},
	}
}

func (c BssParamForCreateVolumeIsAutoRenew) Value() string {
	return c.value
}

func (c BssParamForCreateVolumeIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForCreateVolumeIsAutoRenew) UnmarshalJSON(b []byte) error {
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

type BssParamForCreateVolumePeriodType struct {
	value string
}

type BssParamForCreateVolumePeriodTypeEnum struct {
	MONTH BssParamForCreateVolumePeriodType
	YEAR  BssParamForCreateVolumePeriodType
}

func GetBssParamForCreateVolumePeriodTypeEnum() BssParamForCreateVolumePeriodTypeEnum {
	return BssParamForCreateVolumePeriodTypeEnum{
		MONTH: BssParamForCreateVolumePeriodType{
			value: "month",
		},
		YEAR: BssParamForCreateVolumePeriodType{
			value: "year",
		},
	}
}

func (c BssParamForCreateVolumePeriodType) Value() string {
	return c.value
}

func (c BssParamForCreateVolumePeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForCreateVolumePeriodType) UnmarshalJSON(b []byte) error {
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
