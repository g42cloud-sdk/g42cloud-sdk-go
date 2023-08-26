package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ClusterNodeInformation struct {
	Metadata *ClusterNodeInformationMetadata `json:"metadata"`
}

func (o ClusterNodeInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterNodeInformation struct{}"
	}

	return strings.Join([]string{"ClusterNodeInformation", string(data)}, " ")
}
