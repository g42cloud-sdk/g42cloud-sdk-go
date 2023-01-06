package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type EngineCreateReq struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Payment EngineCreateReqPayment `json:"payment"`

	Flavor EngineCreateReqFlavor `json:"flavor"`

	AzList []string `json:"azList"`

	AuthType EngineCreateReqAuthType `json:"authType"`

	Vpc string `json:"vpc"`

	NetworkId string `json:"networkId"`

	SubnetCidr string `json:"subnetCidr"`

	PublicIpId *string `json:"publicIpId,omitempty"`

	AuthCred *EngineRbacPwd `json:"auth_cred,omitempty"`

	SpecType EngineCreateReqSpecType `json:"specType"`

	Inputs map[string]string `json:"inputs,omitempty"`
}

func (o EngineCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineCreateReq struct{}"
	}

	return strings.Join([]string{"EngineCreateReq", string(data)}, " ")
}

type EngineCreateReqPayment struct {
	value string
}

type EngineCreateReqPaymentEnum struct {
	E_1 EngineCreateReqPayment
}

func GetEngineCreateReqPaymentEnum() EngineCreateReqPaymentEnum {
	return EngineCreateReqPaymentEnum{
		E_1: EngineCreateReqPayment{
			value: "1",
		},
	}
}

func (c EngineCreateReqPayment) Value() string {
	return c.value
}

func (c EngineCreateReqPayment) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineCreateReqPayment) UnmarshalJSON(b []byte) error {
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

type EngineCreateReqFlavor struct {
	value string
}

type EngineCreateReqFlavorEnum struct {
	CSE_S1_SMALL  EngineCreateReqFlavor
	CSE_S1_MEDIUM EngineCreateReqFlavor
	CSE_S1_LARGE  EngineCreateReqFlavor
	CSE_S1_XLARGE EngineCreateReqFlavor
}

func GetEngineCreateReqFlavorEnum() EngineCreateReqFlavorEnum {
	return EngineCreateReqFlavorEnum{
		CSE_S1_SMALL: EngineCreateReqFlavor{
			value: "cse.s1.small",
		},
		CSE_S1_MEDIUM: EngineCreateReqFlavor{
			value: "cse.s1.medium",
		},
		CSE_S1_LARGE: EngineCreateReqFlavor{
			value: "cse.s1.large",
		},
		CSE_S1_XLARGE: EngineCreateReqFlavor{
			value: "cse.s1.xlarge",
		},
	}
}

func (c EngineCreateReqFlavor) Value() string {
	return c.value
}

func (c EngineCreateReqFlavor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineCreateReqFlavor) UnmarshalJSON(b []byte) error {
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

type EngineCreateReqAuthType struct {
	value string
}

type EngineCreateReqAuthTypeEnum struct {
	RBAC EngineCreateReqAuthType
	NONE EngineCreateReqAuthType
}

func GetEngineCreateReqAuthTypeEnum() EngineCreateReqAuthTypeEnum {
	return EngineCreateReqAuthTypeEnum{
		RBAC: EngineCreateReqAuthType{
			value: "RBAC",
		},
		NONE: EngineCreateReqAuthType{
			value: "NONE",
		},
	}
}

func (c EngineCreateReqAuthType) Value() string {
	return c.value
}

func (c EngineCreateReqAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineCreateReqAuthType) UnmarshalJSON(b []byte) error {
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

type EngineCreateReqSpecType struct {
	value string
}

type EngineCreateReqSpecTypeEnum struct {
	CSE  EngineCreateReqSpecType
	CSE2 EngineCreateReqSpecType
}

func GetEngineCreateReqSpecTypeEnum() EngineCreateReqSpecTypeEnum {
	return EngineCreateReqSpecTypeEnum{
		CSE: EngineCreateReqSpecType{
			value: "CSE",
		},
		CSE2: EngineCreateReqSpecType{
			value: "CSE2",
		},
	}
}

func (c EngineCreateReqSpecType) Value() string {
	return c.value
}

func (c EngineCreateReqSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineCreateReqSpecType) UnmarshalJSON(b []byte) error {
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
