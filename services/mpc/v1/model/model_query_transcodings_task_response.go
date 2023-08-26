package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"errors"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/converter"

	"strings"
)

type QueryTranscodingsTaskResponse struct {
	TaskId *string `json:"task_id,omitempty"`

	Status *QueryTranscodingsTaskResponseStatus `json:"status,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	TransTemplateId *[]int32 `json:"trans_template_id,omitempty"`

	Input *ObsObjInfo `json:"input,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`

	OutputFileName *[]string `json:"output_file_name,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	Description *string `json:"description,omitempty"`

	Tips *string `json:"tips,omitempty"`

	TranscodeDetail *TranscodeDetail `json:"transcode_detail,omitempty"`

	ThumbnailOutput *ObsObjInfo `json:"thumbnail_output,omitempty"`

	ThumbnailOutputname *string `json:"thumbnail_outputname,omitempty"`

	PicInfo *[]PicInfo `json:"pic_info,omitempty"`

	AvParameters *[]AvParameters `json:"av_parameters,omitempty"`
}

func (o QueryTranscodingsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTranscodingsTaskResponse struct{}"
	}

	return strings.Join([]string{"QueryTranscodingsTaskResponse", string(data)}, " ")
}

type QueryTranscodingsTaskResponseStatus struct {
	value string
}

type QueryTranscodingsTaskResponseStatusEnum struct {
	NO_TASK          QueryTranscodingsTaskResponseStatus
	WAITING          QueryTranscodingsTaskResponseStatus
	TRANSCODING      QueryTranscodingsTaskResponseStatus
	SUCCEEDED        QueryTranscodingsTaskResponseStatus
	FAILED           QueryTranscodingsTaskResponseStatus
	CANCELED         QueryTranscodingsTaskResponseStatus
	NEED_TO_BE_AUDIT QueryTranscodingsTaskResponseStatus
}

func GetQueryTranscodingsTaskResponseStatusEnum() QueryTranscodingsTaskResponseStatusEnum {
	return QueryTranscodingsTaskResponseStatusEnum{
		NO_TASK: QueryTranscodingsTaskResponseStatus{
			value: "NO_TASK",
		},
		WAITING: QueryTranscodingsTaskResponseStatus{
			value: "WAITING",
		},
		TRANSCODING: QueryTranscodingsTaskResponseStatus{
			value: "TRANSCODING",
		},
		SUCCEEDED: QueryTranscodingsTaskResponseStatus{
			value: "SUCCEEDED",
		},
		FAILED: QueryTranscodingsTaskResponseStatus{
			value: "FAILED",
		},
		CANCELED: QueryTranscodingsTaskResponseStatus{
			value: "CANCELED",
		},
		NEED_TO_BE_AUDIT: QueryTranscodingsTaskResponseStatus{
			value: "NEED_TO_BE_AUDIT",
		},
	}
}

func (c QueryTranscodingsTaskResponseStatus) Value() string {
	return c.value
}

func (c QueryTranscodingsTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryTranscodingsTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
