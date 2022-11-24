package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePrivateBucketAccessRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DomainId string `json:"domain_id"`

	Body *UpdatePrivateBucketAccessBody `json:"body,omitempty"`
}

func (o UpdatePrivateBucketAccessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateBucketAccessRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateBucketAccessRequest", string(data)}, " ")
}
