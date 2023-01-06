package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type GetKieConfigs struct {
	Id *string `json:"id,omitempty"`

	Key *string `json:"key,omitempty"`

	Labels *interface{} `json:"labels,omitempty"`

	Value *string `json:"value,omitempty"`

	ValueType *string `json:"value_type,omitempty"`

	Status *string `json:"status,omitempty"`

	CreateTime *int32 `json:"create_time,omitempty"`

	UpdateTime *int32 `json:"update_time,omitempty"`

	CreateRevision *int64 `json:"create_revision,omitempty"`

	UpdateRevision *int64 `json:"update_revision,omitempty"`
}

func (o GetKieConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetKieConfigs struct{}"
	}

	return strings.Join([]string{"GetKieConfigs", string(data)}, " ")
}
