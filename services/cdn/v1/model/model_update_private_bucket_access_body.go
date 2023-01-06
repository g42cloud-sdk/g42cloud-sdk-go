package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateBucketAccessBody struct {
	Status *bool `json:"status,omitempty"`
}

func (o UpdatePrivateBucketAccessBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateBucketAccessBody struct{}"
	}

	return strings.Join([]string{"UpdatePrivateBucketAccessBody", string(data)}, " ")
}
