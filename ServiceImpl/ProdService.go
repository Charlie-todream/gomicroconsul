package ServiceImpl

import (
	"context"
	"github.com/charlie/gomicroconsul/Services"
	"strconv"
	"time"
)

type ProdService struct {
}

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}

// 实现ProdService.pb.micro.go  ProdServiceHandler interface
func (*ProdService) GetProdList(ctx context.Context, in *Services.ProdRequest, res *Services.ProdListResponse) error {

	time.Sleep(time.Second * 3)  // 设置3秒延迟
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, newProd(100+i, "prodName"+strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}
