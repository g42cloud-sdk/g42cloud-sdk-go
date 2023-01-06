package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListQuotasResponse struct {
	Resources      *[]Quota `json:"resources,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListQuotasResponse", string(data)}, " ")
}
