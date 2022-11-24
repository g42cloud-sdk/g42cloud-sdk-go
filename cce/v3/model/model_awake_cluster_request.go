package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AwakeClusterRequest struct {
	ClusterId string `json:"cluster_id"`
}

func (o AwakeClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AwakeClusterRequest struct{}"
	}

	return strings.Join([]string{"AwakeClusterRequest", string(data)}, " ")
}
