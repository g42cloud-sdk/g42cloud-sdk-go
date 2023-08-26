package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListQualityEnhanceDefaultTemplateRequest struct {
}

func (o ListQualityEnhanceDefaultTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQualityEnhanceDefaultTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListQualityEnhanceDefaultTemplateRequest", string(data)}, " ")
}
