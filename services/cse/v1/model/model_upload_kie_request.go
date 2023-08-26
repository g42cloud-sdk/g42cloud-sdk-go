package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type UploadKieRequest struct {
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	XEngineId string `json:"x-engine-id"`

	Override UploadKieRequestOverride `json:"override"`

	Label *string `json:"label,omitempty"`

	Body *UploadKieRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadKieRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKieRequest struct{}"
	}

	return strings.Join([]string{"UploadKieRequest", string(data)}, " ")
}

type UploadKieRequestOverride struct {
	value string
}

type UploadKieRequestOverrideEnum struct {
	FORCE UploadKieRequestOverride
	ABORT UploadKieRequestOverride
	SKIP  UploadKieRequestOverride
}

func GetUploadKieRequestOverrideEnum() UploadKieRequestOverrideEnum {
	return UploadKieRequestOverrideEnum{
		FORCE: UploadKieRequestOverride{
			value: "force",
		},
		ABORT: UploadKieRequestOverride{
			value: "abort",
		},
		SKIP: UploadKieRequestOverride{
			value: "skip",
		},
	}
}

func (c UploadKieRequestOverride) Value() string {
	return c.value
}

func (c UploadKieRequestOverride) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadKieRequestOverride) UnmarshalJSON(b []byte) error {
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
