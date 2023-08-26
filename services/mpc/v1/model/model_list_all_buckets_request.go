package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAllBucketsRequest struct {
}

func (o ListAllBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllBucketsRequest struct{}"
	}

	return strings.Join([]string{"ListAllBucketsRequest", string(data)}, " ")
}
