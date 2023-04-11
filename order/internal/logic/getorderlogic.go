package logic

import (
	"context"
	"rpc-common/userclient"

	"order/internal/svc"
	"order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	userId := l.getOrderById(req.Id)
	re, _ := l.svcCtx.UserRpc.GetUser(context.Background(), &userclient.IdRequest{
		Id: userId,
	})

	return &types.OrderReply{
		Id:       re.Id,
		Name:     re.Name,
		UserName: re.Name,
	}, nil
}

func (l *GetOrderLogic) getOrderById(id string) string {
	return "1"
}
