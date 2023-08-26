package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type ListStatSummaryResponse struct {
	Summary *[]StatSummary `json:"summary,omitempty"`

	Total *float32 `json:"total,omitempty"`

	StatType       *ListStatSummaryResponseStatType `json:"stat_type,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListStatSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListStatSummaryResponse", string(data)}, " ")
}

type ListStatSummaryResponseStatType struct {
	value string
}

type ListStatSummaryResponseStatTypeEnum struct {
	VIDEO_DURATION        ListStatSummaryResponseStatType
	REMUX_FILE_DURATION   ListStatSummaryResponseStatType
	TRANSCODE_TASK_NUMBER ListStatSummaryResponseStatType
	TRANSCODE_DURATION    ListStatSummaryResponseStatType
}

func GetListStatSummaryResponseStatTypeEnum() ListStatSummaryResponseStatTypeEnum {
	return ListStatSummaryResponseStatTypeEnum{
		VIDEO_DURATION: ListStatSummaryResponseStatType{
			value: "video_duration",
		},
		REMUX_FILE_DURATION: ListStatSummaryResponseStatType{
			value: "remux_file_duration",
		},
		TRANSCODE_TASK_NUMBER: ListStatSummaryResponseStatType{
			value: "transcode_task_number",
		},
		TRANSCODE_DURATION: ListStatSummaryResponseStatType{
			value: "transcode_duration",
		},
	}
}

func (c ListStatSummaryResponseStatType) Value() string {
	return c.value
}

func (c ListStatSummaryResponseStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStatSummaryResponseStatType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
