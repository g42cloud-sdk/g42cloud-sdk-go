package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupportVersions struct {
	ClusterType string `json:"clusterType"`

	ClusterVersion []string `json:"clusterVersion"`
}

func (o SupportVersions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportVersions struct{}"
	}

	return strings.Join([]string{"SupportVersions", string(data)}, " ")
}
