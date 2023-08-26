package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListAllObsObjListResponse struct {
	Objects        *[]ObsObject `json:"objects,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAllObsObjListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllObsObjListResponse struct{}"
	}

	return strings.Join([]string{"ListAllObsObjListResponse", string(data)}, " ")
}
