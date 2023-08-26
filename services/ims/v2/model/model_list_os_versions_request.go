package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListOsVersionsRequest struct {
	Tag *string `json:"tag,omitempty"`
}

func (o ListOsVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOsVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListOsVersionsRequest", string(data)}, " ")
}
