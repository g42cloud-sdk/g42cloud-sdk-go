package v1

import (
	http_client "github.com/g42cloud-sdk/g42cloud-sdk-go/core"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/invoker"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/services/cse/v1/model"
)

type CseClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCseClient(hcClient *http_client.HcHttpClient) *CseClient {
	return &CseClient{HcClient: hcClient}
}

func CseClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

func (c *CseClient) CreateEngine(request *model.CreateEngineRequest) (*model.CreateEngineResponse, error) {
	requestDef := GenReqDefForCreateEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEngineResponse), nil
	}
}

func (c *CseClient) CreateEngineInvoker(request *model.CreateEngineRequest) *CreateEngineInvoker {
	requestDef := GenReqDefForCreateEngine()
	return &CreateEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CseClient) DeleteEngine(request *model.DeleteEngineRequest) (*model.DeleteEngineResponse, error) {
	requestDef := GenReqDefForDeleteEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEngineResponse), nil
	}
}

func (c *CseClient) DeleteEngineInvoker(request *model.DeleteEngineRequest) *DeleteEngineInvoker {
	requestDef := GenReqDefForDeleteEngine()
	return &DeleteEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CseClient) DownloadKie(request *model.DownloadKieRequest) (*model.DownloadKieResponse, error) {
	requestDef := GenReqDefForDownloadKie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadKieResponse), nil
	}
}

func (c *CseClient) DownloadKieInvoker(request *model.DownloadKieRequest) *DownloadKieInvoker {
	requestDef := GenReqDefForDownloadKie()
	return &DownloadKieInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CseClient) ListEngines(request *model.ListEnginesRequest) (*model.ListEnginesResponse, error) {
	requestDef := GenReqDefForListEngines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnginesResponse), nil
	}
}

func (c *CseClient) ListEnginesInvoker(request *model.ListEnginesRequest) *ListEnginesInvoker {
	requestDef := GenReqDefForListEngines()
	return &ListEnginesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CseClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

func (c *CseClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CseClient) ShowEngine(request *model.ShowEngineRequest) (*model.ShowEngineResponse, error) {
	requestDef := GenReqDefForShowEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineResponse), nil
	}
}

func (c *CseClient) ShowEngineInvoker(request *model.ShowEngineRequest) *ShowEngineInvoker {
	requestDef := GenReqDefForShowEngine()
	return &ShowEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CseClient) ShowEngineJob(request *model.ShowEngineJobRequest) (*model.ShowEngineJobResponse, error) {
	requestDef := GenReqDefForShowEngineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineJobResponse), nil
	}
}

func (c *CseClient) ShowEngineJobInvoker(request *model.ShowEngineJobRequest) *ShowEngineJobInvoker {
	requestDef := GenReqDefForShowEngineJob()
	return &ShowEngineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

func (c *CseClient) UploadKie(request *model.UploadKieRequest) (*model.UploadKieResponse, error) {
	requestDef := GenReqDefForUploadKie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadKieResponse), nil
	}
}

func (c *CseClient) UploadKieInvoker(request *model.UploadKieRequest) *UploadKieInvoker {
	requestDef := GenReqDefForUploadKie()
	return &UploadKieInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
