package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateBucketAuthorizedResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateBucketAuthorizedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBucketAuthorizedResponse struct{}"
	}

	return strings.Join([]string{"UpdateBucketAuthorizedResponse", string(data)}, " ")
}
