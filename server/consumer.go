package server

import (
	"consumer-service/pkg/bapi"
	"consumer-service/pkg/errcode"
	pb "consumer-service/proto"
	"context"
	"encoding/json"
)

type GetConsumerListResponse struct {
	Code      int           `json:"code"`
	Message   string        `json:"msg"`
	RequestId string        `json:"requestId"`
	Data      *ConsumerList `json:"result"`
}

type ConsumerList struct {
	Data       []*ConsumerInfo `json:"data"`
	TotalCount int             `json:"totalCount"`
	PageNumber int             `json:"pageNumber"`
	PageSize   int             `json:"pageSize"`
}

type ConsumerInfo struct {
	ConsumerId    string `json:"consumerId"`
	Source        string `json:"source"`
	SourceName    string `json:"sourceName"`
	ConsumerName  string `json:"consumerName"`
	CompanyName   string `json:"companyName"`
	UserGroup     string `json:"userGroup"`
	UserGroupName string `json:"userGroupName"`
	UpdateTime    string `json:"updateTime"`
	UpdateErp     string `json:"updateErp"`
}

type ConsumerServer struct {
	auth *Auth
}

func NewConsumerServer() *ConsumerServer {
	return &ConsumerServer{}
}

func (c *ConsumerServer) GetConsumerList(ctx context.Context, r *pb.GetConsumerListReq) (*pb.GetConsumerListResp, error) {
	if err := c.auth.Check(ctx); err != nil {
		return nil, err
	}

	api := bapi.NewAPI("http://127.0.0.1:8888")
	body, err := api.GetConsumerList(ctx, r.GetSource())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetConsumerListFail, err)
	}
	consumerResp := new(GetConsumerListResponse)
	consumerList := &pb.GetConsumerListResp{}
	if err = json.Unmarshal(body, consumerResp); err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorJsonConvertFail, err)
	}

	resp, err := json.Marshal(consumerResp.Data)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorJsonConvertFail, err)
	}

	//RPC请求屏蔽网络层实现，因此只关心数据即可，不关心code，message等，用rpc error可反馈异常
	if err = json.Unmarshal(resp, consumerList); err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorJsonConvertFail, err)
	}

	return consumerList, nil
}
