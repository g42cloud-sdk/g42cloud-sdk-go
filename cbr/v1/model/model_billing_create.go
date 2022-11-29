package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BillingCreate struct {
	CloudType *BillingCreateCloudType `json:"cloud_type,omitempty"`

	ConsistentLevel BillingCreateConsistentLevel `json:"consistent_level"`

	ObjectType BillingCreateObjectType `json:"object_type"`

	ProtectType BillingCreateProtectType `json:"protect_type"`

	Size int32 `json:"size"`

	ChargingMode *BillingCreateChargingMode `json:"charging_mode,omitempty"`

	PeriodType *BillingCreatePeriodType `json:"period_type,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	ConsoleUrl *string `json:"console_url,omitempty"`

	ExtraInfo *BillbingCreateExtraInfo `json:"extra_info,omitempty"`
}

func (o BillingCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingCreate struct{}"
	}

	return strings.Join([]string{"BillingCreate", string(data)}, " ")
}

type BillingCreateCloudType struct {
	value string
}

type BillingCreateCloudTypeEnum struct {
	PUBLIC BillingCreateCloudType
	HYBRID BillingCreateCloudType
}

func GetBillingCreateCloudTypeEnum() BillingCreateCloudTypeEnum {
	return BillingCreateCloudTypeEnum{
		PUBLIC: BillingCreateCloudType{
			value: "public",
		},
		HYBRID: BillingCreateCloudType{
			value: "hybrid",
		},
	}
}

func (c BillingCreateCloudType) Value() string {
	return c.value
}

func (c BillingCreateCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateCloudType) UnmarshalJSON(b []byte) error {
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

type BillingCreateConsistentLevel struct {
	value string
}

type BillingCreateConsistentLevelEnum struct {
	APP_CONSISTENT   BillingCreateConsistentLevel
	CRASH_CONSISTENT BillingCreateConsistentLevel
}

func GetBillingCreateConsistentLevelEnum() BillingCreateConsistentLevelEnum {
	return BillingCreateConsistentLevelEnum{
		APP_CONSISTENT: BillingCreateConsistentLevel{
			value: "app_consistent",
		},
		CRASH_CONSISTENT: BillingCreateConsistentLevel{
			value: "crash_consistent",
		},
	}
}

func (c BillingCreateConsistentLevel) Value() string {
	return c.value
}

func (c BillingCreateConsistentLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateConsistentLevel) UnmarshalJSON(b []byte) error {
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

type BillingCreateObjectType struct {
	value string
}

type BillingCreateObjectTypeEnum struct {
	SERVER BillingCreateObjectType
	DISK   BillingCreateObjectType
	TURBO  BillingCreateObjectType
}

func GetBillingCreateObjectTypeEnum() BillingCreateObjectTypeEnum {
	return BillingCreateObjectTypeEnum{
		SERVER: BillingCreateObjectType{
			value: "server",
		},
		DISK: BillingCreateObjectType{
			value: "disk",
		},
		TURBO: BillingCreateObjectType{
			value: "turbo",
		},
	}
}

func (c BillingCreateObjectType) Value() string {
	return c.value
}

func (c BillingCreateObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateObjectType) UnmarshalJSON(b []byte) error {
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

type BillingCreateProtectType struct {
	value string
}

type BillingCreateProtectTypeEnum struct {
	BACKUP      BillingCreateProtectType
	REPLICATION BillingCreateProtectType
}

func GetBillingCreateProtectTypeEnum() BillingCreateProtectTypeEnum {
	return BillingCreateProtectTypeEnum{
		BACKUP: BillingCreateProtectType{
			value: "backup",
		},
		REPLICATION: BillingCreateProtectType{
			value: "replication",
		},
	}
}

func (c BillingCreateProtectType) Value() string {
	return c.value
}

func (c BillingCreateProtectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateProtectType) UnmarshalJSON(b []byte) error {
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

type BillingCreateChargingMode struct {
	value string
}

type BillingCreateChargingModeEnum struct {
	POST_PAID BillingCreateChargingMode
	PRE_PAID  BillingCreateChargingMode
}

func GetBillingCreateChargingModeEnum() BillingCreateChargingModeEnum {
	return BillingCreateChargingModeEnum{
		POST_PAID: BillingCreateChargingMode{
			value: "post_paid",
		},
		PRE_PAID: BillingCreateChargingMode{
			value: "pre_paid",
		},
	}
}

func (c BillingCreateChargingMode) Value() string {
	return c.value
}

func (c BillingCreateChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreateChargingMode) UnmarshalJSON(b []byte) error {
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

type BillingCreatePeriodType struct {
	value string
}

type BillingCreatePeriodTypeEnum struct {
	YEAR  BillingCreatePeriodType
	MONTH BillingCreatePeriodType
}

func GetBillingCreatePeriodTypeEnum() BillingCreatePeriodTypeEnum {
	return BillingCreatePeriodTypeEnum{
		YEAR: BillingCreatePeriodType{
			value: "year",
		},
		MONTH: BillingCreatePeriodType{
			value: "month",
		},
	}
}

func (c BillingCreatePeriodType) Value() string {
	return c.value
}

func (c BillingCreatePeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCreatePeriodType) UnmarshalJSON(b []byte) error {
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
