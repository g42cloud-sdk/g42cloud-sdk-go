package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ShowTopUrlRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`

	DomainName string `json:"domain_name"`

	ServiceArea *ShowTopUrlRequestServiceArea `json:"service_area,omitempty"`

	StatType ShowTopUrlRequestStatType `json:"stat_type"`
}

func (o ShowTopUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopUrlRequest struct{}"
	}

	return strings.Join([]string{"ShowTopUrlRequest", string(data)}, " ")
}

type ShowTopUrlRequestServiceArea struct {
	value string
}

type ShowTopUrlRequestServiceAreaEnum struct {
	MAINLAND_CHINA         ShowTopUrlRequestServiceArea
	OUTSIDE_MAINLAND_CHINA ShowTopUrlRequestServiceArea
}

func GetShowTopUrlRequestServiceAreaEnum() ShowTopUrlRequestServiceAreaEnum {
	return ShowTopUrlRequestServiceAreaEnum{
		MAINLAND_CHINA: ShowTopUrlRequestServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: ShowTopUrlRequestServiceArea{
			value: "outside_mainland_china",
		},
	}
}

func (c ShowTopUrlRequestServiceArea) Value() string {
	return c.value
}

func (c ShowTopUrlRequestServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTopUrlRequestServiceArea) UnmarshalJSON(b []byte) error {
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

type ShowTopUrlRequestStatType struct {
	value string
}

type ShowTopUrlRequestStatTypeEnum struct {
	FLUX    ShowTopUrlRequestStatType
	REQ_NUM ShowTopUrlRequestStatType
}

func GetShowTopUrlRequestStatTypeEnum() ShowTopUrlRequestStatTypeEnum {
	return ShowTopUrlRequestStatTypeEnum{
		FLUX: ShowTopUrlRequestStatType{
			value: "flux",
		},
		REQ_NUM: ShowTopUrlRequestStatType{
			value: "req_num",
		},
	}
}

func (c ShowTopUrlRequestStatType) Value() string {
	return c.value
}

func (c ShowTopUrlRequestStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTopUrlRequestStatType) UnmarshalJSON(b []byte) error {
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
