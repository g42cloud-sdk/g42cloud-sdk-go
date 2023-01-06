package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Billing struct {
	Allocated int32 `json:"allocated"`

	ChargingMode BillingChargingMode `json:"charging_mode"`

	CloudType *BillingCloudType `json:"cloud_type,omitempty"`

	ConsistentLevel BillingConsistentLevel `json:"consistent_level"`

	ObjectType *BillingObjectType `json:"object_type,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	ProductId *string `json:"product_id,omitempty"`

	ProtectType BillingProtectType `json:"protect_type"`

	Size int32 `json:"size"`

	SpecCode BillingSpecCode `json:"spec_code"`

	Status BillingStatus `json:"status"`

	StorageUnit *string `json:"storage_unit,omitempty"`

	Used int32 `json:"used"`

	FrozenScene *string `json:"frozen_scene,omitempty"`
}

func (o Billing) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Billing struct{}"
	}

	return strings.Join([]string{"Billing", string(data)}, " ")
}

type BillingChargingMode struct {
	value string
}

type BillingChargingModeEnum struct {
	PRE_PAID  BillingChargingMode
	POST_PAID BillingChargingMode
}

func GetBillingChargingModeEnum() BillingChargingModeEnum {
	return BillingChargingModeEnum{
		PRE_PAID: BillingChargingMode{
			value: "pre_paid",
		},
		POST_PAID: BillingChargingMode{
			value: "post_paid",
		},
	}
}

func (c BillingChargingMode) Value() string {
	return c.value
}

func (c BillingChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingChargingMode) UnmarshalJSON(b []byte) error {
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

type BillingCloudType struct {
	value string
}

type BillingCloudTypeEnum struct {
	PUBLIC BillingCloudType
	HYBRID BillingCloudType
}

func GetBillingCloudTypeEnum() BillingCloudTypeEnum {
	return BillingCloudTypeEnum{
		PUBLIC: BillingCloudType{
			value: "public",
		},
		HYBRID: BillingCloudType{
			value: "hybrid",
		},
	}
}

func (c BillingCloudType) Value() string {
	return c.value
}

func (c BillingCloudType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingCloudType) UnmarshalJSON(b []byte) error {
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

type BillingConsistentLevel struct {
	value string
}

type BillingConsistentLevelEnum struct {
	APP_CONSISTENT   BillingConsistentLevel
	CRASH_CONSISTENT BillingConsistentLevel
}

func GetBillingConsistentLevelEnum() BillingConsistentLevelEnum {
	return BillingConsistentLevelEnum{
		APP_CONSISTENT: BillingConsistentLevel{
			value: "app_consistent",
		},
		CRASH_CONSISTENT: BillingConsistentLevel{
			value: "crash_consistent",
		},
	}
}

func (c BillingConsistentLevel) Value() string {
	return c.value
}

func (c BillingConsistentLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingConsistentLevel) UnmarshalJSON(b []byte) error {
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

type BillingObjectType struct {
	value string
}

type BillingObjectTypeEnum struct {
	SERVER BillingObjectType
	DISK   BillingObjectType
}

func GetBillingObjectTypeEnum() BillingObjectTypeEnum {
	return BillingObjectTypeEnum{
		SERVER: BillingObjectType{
			value: "server",
		},
		DISK: BillingObjectType{
			value: "disk",
		},
	}
}

func (c BillingObjectType) Value() string {
	return c.value
}

func (c BillingObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingObjectType) UnmarshalJSON(b []byte) error {
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

type BillingProtectType struct {
	value string
}

type BillingProtectTypeEnum struct {
	BACKUP      BillingProtectType
	REPLICATION BillingProtectType
	HYBRID      BillingProtectType
}

func GetBillingProtectTypeEnum() BillingProtectTypeEnum {
	return BillingProtectTypeEnum{
		BACKUP: BillingProtectType{
			value: "backup",
		},
		REPLICATION: BillingProtectType{
			value: "replication",
		},
		HYBRID: BillingProtectType{
			value: "hybrid",
		},
	}
}

func (c BillingProtectType) Value() string {
	return c.value
}

func (c BillingProtectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingProtectType) UnmarshalJSON(b []byte) error {
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

type BillingSpecCode struct {
	value string
}

type BillingSpecCodeEnum struct {
	VAULT_BACKUP_SERVER_NORMAL BillingSpecCode
	VAULT_BACKUP_VOLUME_NORMAL BillingSpecCode
}

func GetBillingSpecCodeEnum() BillingSpecCodeEnum {
	return BillingSpecCodeEnum{
		VAULT_BACKUP_SERVER_NORMAL: BillingSpecCode{
			value: "vault.backup.server.normal",
		},
		VAULT_BACKUP_VOLUME_NORMAL: BillingSpecCode{
			value: "vault.backup.volume.normal",
		},
	}
}

func (c BillingSpecCode) Value() string {
	return c.value
}

func (c BillingSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BillingSpecCode) UnmarshalJSON(b []byte) error {
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
