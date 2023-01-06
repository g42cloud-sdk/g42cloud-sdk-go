package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type NodePoolSpec struct {
	Type *NodePoolSpecType `json:"type,omitempty"`

	NodeTemplate *NodeSpec `json:"nodeTemplate"`

	InitialNodeCount *int32 `json:"initialNodeCount,omitempty"`

	Autoscaling *NodePoolNodeAutoscaling `json:"autoscaling,omitempty"`

	NodeManagement *NodeManagement `json:"nodeManagement,omitempty"`

	PodSecurityGroups *[]SecurityId `json:"podSecurityGroups,omitempty"`
}

func (o NodePoolSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolSpec struct{}"
	}

	return strings.Join([]string{"NodePoolSpec", string(data)}, " ")
}

type NodePoolSpecType struct {
	value string
}

type NodePoolSpecTypeEnum struct {
	VM          NodePoolSpecType
	ELASTIC_BMS NodePoolSpecType
}

func GetNodePoolSpecTypeEnum() NodePoolSpecTypeEnum {
	return NodePoolSpecTypeEnum{
		VM: NodePoolSpecType{
			value: "vm",
		},
		ELASTIC_BMS: NodePoolSpecType{
			value: "ElasticBMS",
		},
	}
}

func (c NodePoolSpecType) Value() string {
	return c.value
}

func (c NodePoolSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodePoolSpecType) UnmarshalJSON(b []byte) error {
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
