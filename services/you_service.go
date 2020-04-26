package services

import (
	"context"
	"strconv"
)

type Hello_Service struct{}

func (Hello_Service) Hello(ctx context.Context,r *Request) (*Response, error) {
	return &Response{Res:"hello "+r.Name},nil
}

type Product_Service struct {}
func (Product_Service) SomeProduct(ctx context.Context,r *QuerySize) (*ProductList,error){
	res := make([]*Response,0)
	for i := 1; i <= int(r.Size); i++{
		icopy := i
		res = append(res,&Response{Res:"product:"+strconv.Itoa(icopy)})
	}
	return &ProductList{Res: res},nil
}
func (Product_Service) ProductInfo(ctx context.Context, r *QuerySize) (*ProductInfo,error){
	productinfo := &ProductInfo{
		Pid:                  r.Size,
		Pname:                "product"+strconv.Itoa(int(r.Size)),
		Price:                199.00,
	}
	return productinfo,nil
}

type Order_Service struct{}

func (Order_Service) NewOrder(ctx context.Context,r *OrderRequest) (*OrderResponse, error) {
	//fmt.Println(r)
	return &OrderResponse{
		Status:               "ok",
		Message:              "succ",
	},nil
}
