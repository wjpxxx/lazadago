package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductsProductsResponseEntity struct{
    PrimaryCategory	int	`json:"primary_category"`
    Attributes	GetProductsAttributesResponseEntity	`json:"attributes"`
    Skus	[]GetProductsSkusResponseEntity	`json:"skus"`
    ItemId	int	`json:"item_id"`
    CreatedTime	string	`json:"created_time"`
    UpdatedTime	string	`json:"updated_time"`
    Images	string	`json:"images"`
}
func (g GetProductsProductsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
//GetProductsAttributesResponseEntity
type GetProductsAttributesResponseEntity struct{
    ShortDescription string `json:"short_description"`
    Name string `json:"name"`
    Description string `json:"description"`
    WarrantyType string `json:"warranty_type"`
    Brand string `json:"brand"`
}
func (g GetProductsAttributesResponseEntity) String() string {
    return lib.ObjectToString(g)
}
//GetProductsSkusResponseEntity
type GetProductsSkusResponseEntity struct{
    Status string `json:"Status"`
    SkuId int64 `json:"SkuId"`
    Quantity int `json:"quantity"`
    ProductWeight string `json:"product_weight"`
    Images []string `json:"Images"`
    SellerSku string `json:"SellerSku"`
    ShopSku string `json:"ShopSku"`
    Url string `json:"Url"`
    PackageWidth string `json:"package_width"`
    SpecialToTime string `json:"special_to_time"`
    SpecialFromTime string `json:"special_from_time"`
    PackageHeight string `json:"package_height"`
    SpecialPrice float32 `json:"special_price"`
    Price float32 `json:"price"`
    PackageLength string `json:"package_length"`
    PackageWeight string `json:"package_weight"`
    Available int `json:"Available"`
    SpecialoDate string `json:"special_to_date"`
}
func (g GetProductsSkusResponseEntity) String() string {
    return lib.ObjectToString(g)
}