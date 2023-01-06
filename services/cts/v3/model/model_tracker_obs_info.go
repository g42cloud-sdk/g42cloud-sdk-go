package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TrackerObsInfo struct {
	BucketName *string `json:"bucket_name,omitempty"`

	FilePrefixName *string `json:"file_prefix_name,omitempty"`

	IsObsCreated *bool `json:"is_obs_created,omitempty"`

	BucketLifecycle *int32 `json:"bucket_lifecycle,omitempty"`
}

func (o TrackerObsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerObsInfo struct{}"
	}

	return strings.Join([]string{"TrackerObsInfo", string(data)}, " ")
}
