package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAllBucketsRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`
}

func (o ListAllBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllBucketsRequest struct{}"
	}

	return strings.Join([]string{"ListAllBucketsRequest", string(data)}, " ")
}
