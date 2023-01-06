package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ObsInfo struct {
	BucketName *string `json:"bucket_name,omitempty"`

	FilePrefixName *string `json:"file_prefix_name,omitempty"`

	IsObsCreated *bool `json:"is_obs_created,omitempty"`

	IsAuthorizedBucket *bool `json:"is_authorized_bucket,omitempty"`

	BucketLifecycle *int64 `json:"bucket_lifecycle,omitempty"`
}

func (o ObsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsInfo struct{}"
	}

	return strings.Join([]string{"ObsInfo", string(data)}, " ")
}
