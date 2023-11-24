package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// OverviewDataGet 获取企业号基础数据
func OverviewDataGet(clt *core.SDKClient, accessToken string, req *enterprise.ReportRequest) (*enterprise.ReportResult, error) {
	var resp enterprise.ReportResponse
	if err := clt.Get("v1.0/enterprise/overview/data/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
