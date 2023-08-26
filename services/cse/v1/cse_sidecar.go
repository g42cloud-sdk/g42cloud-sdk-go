package v1

import "github.com/g42cloud-sdk/g42cloud-sdk-go/services/cse/v1/model"

func (c *CseClient) GetCreateEngineRequest() *model.CreateEngineRequest {
	return new(model.CreateEngineRequest)
}

func (c *CseClient) GetDeleteEngineRequest() *model.DeleteEngineRequest {
	return new(model.DeleteEngineRequest)
}

func (c *CseClient) GetDownloadKieRequest() *model.DownloadKieRequest {
	return new(model.DownloadKieRequest)
}

func (c *CseClient) GetListEnginesRequest() *model.ListEnginesRequest {
	return new(model.ListEnginesRequest)
}

func (c *CseClient) GetListFlavorsRequest() *model.ListFlavorsRequest {
	return new(model.ListFlavorsRequest)
}

func (c *CseClient) GetShowEngineRequest() *model.ShowEngineRequest {
	return new(model.ShowEngineRequest)
}

func (c *CseClient) GetShowEngineJobRequest() *model.ShowEngineJobRequest {
	return new(model.ShowEngineJobRequest)
}

func (c *CseClient) GetUploadKieRequest() *model.UploadKieRequest {
	return new(model.UploadKieRequest)
}
