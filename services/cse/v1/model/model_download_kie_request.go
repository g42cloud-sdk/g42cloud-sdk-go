package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type DownloadKieRequest struct {
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	XEngineId string `json:"x-engine-id"`

	Label *string `json:"label,omitempty"`

	Match *DownloadKieRequestMatch `json:"match,omitempty"`

	Body *DownloadKieReqBody `json:"body,omitempty"`
}

func (o DownloadKieRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieRequest struct{}"
	}

	return strings.Join([]string{"DownloadKieRequest", string(data)}, " ")
}

type DownloadKieRequestMatch struct {
	value string
}

type DownloadKieRequestMatchEnum struct {
	EXACT DownloadKieRequestMatch
}

func GetDownloadKieRequestMatchEnum() DownloadKieRequestMatchEnum {
	return DownloadKieRequestMatchEnum{
		EXACT: DownloadKieRequestMatch{
			value: "exact",
		},
	}
}

func (c DownloadKieRequestMatch) Value() string {
	return c.value
}

func (c DownloadKieRequestMatch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadKieRequestMatch) UnmarshalJSON(b []byte) error {
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
