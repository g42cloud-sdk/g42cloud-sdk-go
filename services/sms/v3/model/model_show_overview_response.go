package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowOverviewResponse struct {
	Waiting *int32 `json:"waiting,omitempty"`

	Replicating *int32 `json:"replicating,omitempty"`

	Syncing *int32 `json:"syncing,omitempty"`

	Other          *int32 `json:"other,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowOverviewResponse", string(data)}, " ")
}
