package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAllObsObjListRequest struct {
	Authorization *string `json:"Authorization,omitempty"`

	XProjectId *string `json:"X-Project_Id,omitempty"`

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	Bucket string `json:"bucket"`

	Prefix *string `json:"prefix,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o ListAllObsObjListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllObsObjListRequest struct{}"
	}

	return strings.Join([]string{"ListAllObsObjListRequest", string(data)}, " ")
}
