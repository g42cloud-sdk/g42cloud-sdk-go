package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListStorageTypesResponse struct {
	StorageType *[]Storage `json:"storage_type,omitempty"`

	DsspoolInfo    *[]DssPoolInfo `json:"dsspool_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListStorageTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageTypesResponse struct{}"
	}

	return strings.Join([]string{"ListStorageTypesResponse", string(data)}, " ")
}
