package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type BillingCreate struct {
	CloudType *string `json:"cloud_type,omitempty"`

	ConsistentLevel string `json:"consistent_level"`

	ObjectType string `json:"object_type"`

	ProtectType string `json:"protect_type"`

	Size int32 `json:"size"`

	ChargingMode *string `json:"charging_mode,omitempty"`

	PeriodType *BillingCreatePeriodType `json:"period_type,omitempty"`

	PeriodNum *int32 `json:"period_num,omitempty"`

	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	ConsoleUrl *string `json:"console_url,omitempty"`

	IsMultiAz *bool `json:"is_multi_az,omitempty"`

	PromotionInfo *string `json:"promotion_info,omitempty"`

	PurchaseMode *string `json:"purchase_mode,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	ExtraInfo *BillbingCreateExtraInfo `json:"extra_info,omitempty"`
}

func (o BillingCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingCreate struct{}"
	}

	return strings.Join([]string{"BillingCreate", string(data)}, " ")
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
