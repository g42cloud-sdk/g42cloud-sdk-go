package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAnimatedGraphicsTaskResponse struct {
	Total *int32 `json:"total,omitempty"`

	Tasks          *[]AnimatedGraphicsTask `json:"tasks,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListAnimatedGraphicsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"ListAnimatedGraphicsTaskResponse", string(data)}, " ")
}
