package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPlatformProductsResult struct{
    Data	[]GetPlatformProductsDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetPlatformProductsDetailResponseEntity	`json:"detail"`
}
func (g GetPlatformProductsResult) String() string {
    return lib.ObjectToString(g)
}
type GetPlatformProductsDetailResponseEntity struct{}
