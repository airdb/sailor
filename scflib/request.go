package scflib

import (
	"fmt"
	"strconv"

	"github.com/tencentyun/scf-go-lib/events"
)

type ListReq struct {
	// PageNo   uint `url:"page_no"`
	// PageSize uint `url:"page_size"`
	PageNo   uint `json:"p"`
	PageSize uint `url:"from"`
}

func GetReqParm(qs events.APIGatewayQueryString) (pageNo, pageSize uint) {
	a := []byte{}
	// var a ListReq

	fmt.Println("xx", a)
	fmt.Println("xx", qs["page_no"])
	fmt.Println("xx", qs["page_size"])
	// pageNo := qs["page_no"][0]
	no64, _ := strconv.ParseUint(qs["pageNo"][0], 10, 32)
	pageNo = uint(no64)

	size64, _ := strconv.ParseUint(qs["pageSize"][0], 10, 32)
	pageSize = uint(size64)

	fmt.Println("xxxx", pageNo, pageSize)

	return
}

func GetReqID(qs events.APIGatewayQueryString) (id uint) {
	no64, _ := strconv.ParseUint(qs["id"][0], 10, 32)
	id = uint(no64)

	return id
}

// type queryBinding struct{}.
func BindQuery(qs events.APIGatewayQueryString, ptr interface{}) error {
	return mapForm(ptr, qs)
}
