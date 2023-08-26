package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateBucketAuthorizedRequest struct {
	Body *BucketAuthorizedReq `json:"body,omitempty"`
}

func (o UpdateBucketAuthorizedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBucketAuthorizedRequest struct{}"
	}

	return strings.Join([]string{"UpdateBucketAuthorizedRequest", string(data)}, " ")
}
