package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
