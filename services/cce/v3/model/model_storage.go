package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Storage struct {
	StorageSelectors []StorageSelectors `json:"storageSelectors"`

	StorageGroups []StorageGroups `json:"storageGroups"`
}

func (o Storage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Storage struct{}"
	}

	return strings.Join([]string{"Storage", string(data)}, " ")
}
