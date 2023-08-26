package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMbTasksReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMbTasksReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMbTasksReportResponse struct{}"
	}

	return strings.Join([]string{"CreateMbTasksReportResponse", string(data)}, " ")
}
