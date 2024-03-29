package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type Spec struct {
	Id *int64 `json:"id,omitempty"`

	EngineId *string `json:"engine_id,omitempty"`

	SpecType *SpecSpecType `json:"spec_type,omitempty"`

	Cluster *string `json:"cluster,omitempty"`

	ClusterId *string `json:"cluster_id,omitempty"`

	ClusterNodes *SpecClusterNode `json:"cluster_nodes,omitempty"`

	Flavor *string `json:"flavor,omitempty"`

	Region *string `json:"region,omitempty"`

	Version *string `json:"version,omitempty"`

	ExtendParam *string `json:"extend_param,omitempty"`
}

func (o Spec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Spec struct{}"
	}

	return strings.Join([]string{"Spec", string(data)}, " ")
}

type SpecSpecType struct {
	value string
}

type SpecSpecTypeEnum struct {
	CCE          SpecSpecType
	CSE          SpecSpecType
	SPRING_CLOUD SpecSpecType
}

func GetSpecSpecTypeEnum() SpecSpecTypeEnum {
	return SpecSpecTypeEnum{
		CCE: SpecSpecType{
			value: "CCE",
		},
		CSE: SpecSpecType{
			value: "CSE",
		},
		SPRING_CLOUD: SpecSpecType{
			value: "SpringCloud",
		},
	}
}

func (c SpecSpecType) Value() string {
	return c.value
}

func (c SpecSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SpecSpecType) UnmarshalJSON(b []byte) error {
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
