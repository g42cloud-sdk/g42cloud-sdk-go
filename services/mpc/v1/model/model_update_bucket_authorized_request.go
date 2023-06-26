package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateBucketAuthorizedRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Body *BucketAuthorizedReq `json:"body,omitempty"`
}

func (o UpdateBucketAuthorizedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBucketAuthorizedRequest struct{}"
	}

	return strings.Join([]string{"UpdateBucketAuthorizedRequest", string(data)}, " ")
}
