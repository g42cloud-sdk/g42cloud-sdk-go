package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type DomainBody struct {
	DomainName string `json:"domain_name"`

	BusinessType DomainBodyBusinessType `json:"business_type"`

	Sources []Sources `json:"sources"`

	ServiceArea DomainBodyServiceArea `json:"service_area"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DomainBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainBody struct{}"
	}

	return strings.Join([]string{"DomainBody", string(data)}, " ")
}

type DomainBodyBusinessType struct {
	value string
}

type DomainBodyBusinessTypeEnum struct {
	WEB        DomainBodyBusinessType
	DOWNLOAD   DomainBodyBusinessType
	VIDEO      DomainBodyBusinessType
	WHOLE_SITE DomainBodyBusinessType
}

func GetDomainBodyBusinessTypeEnum() DomainBodyBusinessTypeEnum {
	return DomainBodyBusinessTypeEnum{
		WEB: DomainBodyBusinessType{
			value: "web",
		},
		DOWNLOAD: DomainBodyBusinessType{
			value: "download",
		},
		VIDEO: DomainBodyBusinessType{
			value: "video",
		},
		WHOLE_SITE: DomainBodyBusinessType{
			value: "wholeSite",
		},
	}
}

func (c DomainBodyBusinessType) Value() string {
	return c.value
}

func (c DomainBodyBusinessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainBodyBusinessType) UnmarshalJSON(b []byte) error {
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

type DomainBodyServiceArea struct {
	value string
}

type DomainBodyServiceAreaEnum struct {
	MAINLAND_CHINA         DomainBodyServiceArea
	OUTSIDE_MAINLAND_CHINA DomainBodyServiceArea
	GLOBAL                 DomainBodyServiceArea
}

func GetDomainBodyServiceAreaEnum() DomainBodyServiceAreaEnum {
	return DomainBodyServiceAreaEnum{
		MAINLAND_CHINA: DomainBodyServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: DomainBodyServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: DomainBodyServiceArea{
			value: "global",
		},
	}
}

func (c DomainBodyServiceArea) Value() string {
	return c.value
}

func (c DomainBodyServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DomainBodyServiceArea) UnmarshalJSON(b []byte) error {
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
