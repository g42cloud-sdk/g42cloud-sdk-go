package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Billing struct {
	Allocated int32 `json:"allocated"`

	ChargingMode string `json:"charging_mode"`

	CloudType *string `json:"cloud_type,omitempty"`

	ConsistentLevel string `json:"consistent_level"`

	ObjectType *string `json:"object_type,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	ProtectType string `json:"protect_type"`

	Size int32 `json:"size"`

	SpecCode string `json:"spec_code"`

	Status BillingStatus `json:"status"`

	StorageUnit *string `json:"storage_unit,omitempty"`

	Used int32 `json:"used"`

	FrozenScene *string `json:"frozen_scene,omitempty"`

	IsMultiAz *bool `json:"is_multi_az,omitempty"`
}

func (o Billing) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Billing struct{}"
	}

	return strings.Join([]string{"Billing", string(data)}, " ")
}

type BillingStatus struct {
	value string
}

type BillingStatusEnum struct {
	AVAILABLE BillingStatus
	LOCK      BillingStatus
	FROZEN    BillingStatus
	DELETING  BillingStatus
	ERROR     BillingStatus
}

func GetBillingStatusEnum() BillingStatusEnum {
	return BillingStatusEnum{
		AVAILABLE: BillingStatus{
			value: "available",
		},
		LOCK: BillingStatus{
			value: "lock",
		},
		FROZEN: BillingStatus{
			value: "frozen",
		},
		DELETING: BillingStatus{
			value: "deleting",
		},
		ERROR: BillingStatus{
			value: "error",
		},
	}
}

func (c BillingStatus) Value() string {
	return c.value
}

func (c BillingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingStatus) UnmarshalJSON(b []byte) error {
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
