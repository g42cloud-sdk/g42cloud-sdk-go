package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdatePrivateBucketAccessResponse struct {
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdatePrivateBucketAccessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateBucketAccessResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateBucketAccessResponse", string(data)}, " ")
}
