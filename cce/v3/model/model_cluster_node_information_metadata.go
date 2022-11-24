package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterNodeInformationMetadata struct {
	Name string `json:"name"`
}

func (o ClusterNodeInformationMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNodeInformationMetadata struct{}"
	}

	return strings.Join([]string{"ClusterNodeInformationMetadata", string(data)}, " ")
}
