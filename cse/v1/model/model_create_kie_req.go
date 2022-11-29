package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateKieReq struct {
	Key *string `json:"key,omitempty"`

	Labels *interface{} `json:"labels,omitempty"`

	Value *string `json:"value,omitempty"`

	ValueType *string `json:"value_type,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o CreateKieReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKieReq struct{}"
	}

	return strings.Join([]string{"CreateKieReq", string(data)}, " ")
}
