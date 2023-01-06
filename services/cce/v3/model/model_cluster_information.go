package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ClusterInformation struct {
	Spec *ClusterInformationSpec `json:"spec"`
}

func (o ClusterInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInformation struct{}"
	}

	return strings.Join([]string{"ClusterInformation", string(data)}, " ")
}
