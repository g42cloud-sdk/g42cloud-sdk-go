package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type GlanceListImageMembersResponse struct {
	Members *[]GlanceImageMembers `json:"members,omitempty"`

	Schema         *string `json:"schema,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceListImageMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageMembersResponse struct{}"
	}

	return strings.Join([]string{"GlanceListImageMembersResponse", string(data)}, " ")
}
