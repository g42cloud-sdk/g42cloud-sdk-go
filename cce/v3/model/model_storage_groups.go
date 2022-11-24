package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StorageGroups struct {
	Name string `json:"name"`

	CceManaged *bool `json:"cceManaged,omitempty"`

	SelectorNames []string `json:"selectorNames"`

	VirtualSpaces []VirtualSpace `json:"virtualSpaces"`
}

func (o StorageGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageGroups struct{}"
	}

	return strings.Join([]string{"StorageGroups", string(data)}, " ")
}
