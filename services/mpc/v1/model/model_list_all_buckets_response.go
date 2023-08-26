package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAllBucketsResponse struct {
	Buckets        *[]ObsBucket `json:"buckets,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAllBucketsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllBucketsResponse struct{}"
	}

	return strings.Join([]string{"ListAllBucketsResponse", string(data)}, " ")
}
