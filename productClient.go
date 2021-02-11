/*
* @Time    : 2021-02-11 11:47
* @Author  : CoderCharm
* @File    : productClient.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/

package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/wxyMicro/product/common"
	goMicroServiceProduct "github.com/wxyMicro/product/proto/product"
	"log"
)

func main() {
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	t, io, err := common.NewTracer("go.micro.service.product.client", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		//添加注册中心
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := goMicroServiceProduct.NewProductService("go.micro.service.product", service.Client())

	productAdd := &goMicroServiceProduct.ProductInfo{
		ProductName:        "产品名字",
		ProductSku:         "sku型号",
		ProductPrice:       1.1,
		ProductDescription: "介绍....",
		ProductCategoryId:  1,
		ProductImage: []*goMicroServiceProduct.ProductImage{
			{
				ImageName: "image_name_01",
				ImageCode: "image01",
				ImageUrl:  "https://image.3001.net/images/20200504/1588558613_5eaf7b159c8e9.jpeg",
			},
			{
				ImageName: "image_name_02",
				ImageCode: "image02",
				ImageUrl:  "https://image.3001.net/images/20200504/1588558613_5eaf7b159c8e9.jpeg",
			},
		},
		ProductSize: []*goMicroServiceProduct.ProductSize{
			{
				SizeName: "product-size",
				SizeCode: "product-size-code",
			},
		},
		ProductSeo: &goMicroServiceProduct.ProductSeo{
			SeoTitle:       "product-seo",
			SeoKeywords:    "product-seo",
			SeoDescription: "product-seo",
			SeoCode:        "product-seo",
		},
	}
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
