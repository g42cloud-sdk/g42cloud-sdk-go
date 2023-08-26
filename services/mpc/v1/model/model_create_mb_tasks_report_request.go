package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMbTasksReportRequest struct {
	Body *MbTasksReportReq `json:"body,omitempty"`
}

func (o CreateMbTasksReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMbTasksReportRequest struct{}"
	}

	return strings.Join([]string{"CreateMbTasksReportRequest", string(data)}, " ")
}
