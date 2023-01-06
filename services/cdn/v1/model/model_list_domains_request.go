package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListDomainsRequest struct {
	DomainName *string `json:"domain_name,omitempty"`

	BusinessType *ListDomainsRequestBusinessType `json:"business_type,omitempty"`

	DomainStatus *ListDomainsRequestDomainStatus `json:"domain_status,omitempty"`

	ServiceArea *ListDomainsRequestServiceArea `json:"service_area,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNumber *int32 `json:"page_number,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainsRequest", string(data)}, " ")
}

type ListDomainsRequestBusinessType struct {
	value string
}

type ListDomainsRequestBusinessTypeEnum struct {
	WEB        ListDomainsRequestBusinessType
	DOWNLOAD   ListDomainsRequestBusinessType
	VIDEO      ListDomainsRequestBusinessType
	WHOLE_SITE ListDomainsRequestBusinessType
}

func GetListDomainsRequestBusinessTypeEnum() ListDomainsRequestBusinessTypeEnum {
	return ListDomainsRequestBusinessTypeEnum{
		WEB: ListDomainsRequestBusinessType{
			value: "web",
		},
		DOWNLOAD: ListDomainsRequestBusinessType{
			value: "download",
		},
		VIDEO: ListDomainsRequestBusinessType{
			value: "video",
		},
		WHOLE_SITE: ListDomainsRequestBusinessType{
			value: "wholeSite",
		},
	}
}

func (c ListDomainsRequestBusinessType) Value() string {
	return c.value
}

func (c ListDomainsRequestBusinessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDomainsRequestBusinessType) UnmarshalJSON(b []byte) error {
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

type ListDomainsRequestDomainStatus struct {
	value string
}

type ListDomainsRequestDomainStatusEnum struct {
	ONLINE           ListDomainsRequestDomainStatus
	OFFLINE          ListDomainsRequestDomainStatus
	CONFIGURING      ListDomainsRequestDomainStatus
	CONFIGURE_FAILED ListDomainsRequestDomainStatus
	CHECKING         ListDomainsRequestDomainStatus
	CHECK_FAILED     ListDomainsRequestDomainStatus
	DELETING         ListDomainsRequestDomainStatus
}

func GetListDomainsRequestDomainStatusEnum() ListDomainsRequestDomainStatusEnum {
	return ListDomainsRequestDomainStatusEnum{
		ONLINE: ListDomainsRequestDomainStatus{
			value: "online",
		},
		OFFLINE: ListDomainsRequestDomainStatus{
			value: "offline",
		},
		CONFIGURING: ListDomainsRequestDomainStatus{
			value: "configuring",
		},
		CONFIGURE_FAILED: ListDomainsRequestDomainStatus{
			value: "configure_failed",
		},
		CHECKING: ListDomainsRequestDomainStatus{
			value: "checking",
		},
		CHECK_FAILED: ListDomainsRequestDomainStatus{
			value: "check_failed",
		},
		DELETING: ListDomainsRequestDomainStatus{
			value: "deleting",
		},
	}
}

func (c ListDomainsRequestDomainStatus) Value() string {
	return c.value
}

func (c ListDomainsRequestDomainStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDomainsRequestDomainStatus) UnmarshalJSON(b []byte) error {
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

type ListDomainsRequestServiceArea struct {
	value string
}

type ListDomainsRequestServiceAreaEnum struct {
	MAINLAND_CHINA         ListDomainsRequestServiceArea
	OUTSIDE_MAINLAND_CHINA ListDomainsRequestServiceArea
	GLOBAL                 ListDomainsRequestServiceArea
}

func GetListDomainsRequestServiceAreaEnum() ListDomainsRequestServiceAreaEnum {
	return ListDomainsRequestServiceAreaEnum{
		MAINLAND_CHINA: ListDomainsRequestServiceArea{
			value: "mainland_china",
		},
		OUTSIDE_MAINLAND_CHINA: ListDomainsRequestServiceArea{
			value: "outside_mainland_china",
		},
		GLOBAL: ListDomainsRequestServiceArea{
			value: "global",
		},
	}
}

func (c ListDomainsRequestServiceArea) Value() string {
	return c.value
}

func (c ListDomainsRequestServiceArea) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDomainsRequestServiceArea) UnmarshalJSON(b []byte) error {
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
