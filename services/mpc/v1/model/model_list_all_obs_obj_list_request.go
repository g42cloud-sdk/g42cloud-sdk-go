package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAllObsObjListRequest struct {
	Bucket string `json:"bucket"`

	Prefix *string `json:"prefix,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o ListAllObsObjListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllObsObjListRequest struct{}"
	}

	return strings.Join([]string{"ListAllObsObjListRequest", string(data)}, " ")
}
