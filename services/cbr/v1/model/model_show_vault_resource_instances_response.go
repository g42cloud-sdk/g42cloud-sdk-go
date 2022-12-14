package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowVaultResourceInstancesResponse struct {
	Resources *[]TagResource `json:"resources,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowVaultResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowVaultResourceInstancesResponse", string(data)}, " ")
}
