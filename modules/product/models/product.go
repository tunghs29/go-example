package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
)

type Alias = int

type ProductType uint

const (
	Good ProductType = iota
	Combo
	Produce
)

func (pt ProductType) String() string {
	return []string{"good", "combo", "produce"}[pt]
}

type Attribute struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Value    string             `json:"value" bson:"value"`
	Position int                `json:"position" bson:"position"`
}

type Attributes []Attribute

type Image struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FileType        string             `json:"file_type" bson:"file_type"`
	FileName        string             `json:"file_name" bson:"file_name"`
	FileSize        int                `json:"file_size" bson:"file_size"`
	FileDescription string             `json:"file_description,omitempty" bson:"file_description"`
	FileMineType    string             `json:"file_mine_type" bson:"file_mine_type"`
}

type Images []Image

type BaseUnit struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

type ProductOrigin struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Code string             `json:"code" bson:"code"`
}

type DeactiveBranch struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

type DeactiveBranchs []DeactiveBranch

type Unit struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Type       string             `json:"type" bson:"type"`
	Conversion float64            `json:"conversion" bson:"conversion"`
}

type ImageId struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
}

type ImageIds []ImageId

type Variant struct {
	Id                    primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name                  string             `json:"name" bson:"name"`
	VariantName           string             `json:"variant_name" bson:"variant_name"`
	Barcode               string             `json:"barcode" bson:"barcode"`
	Price                 float64            `json:"price" bson:"price"`
	PurchasePrice         float64            `json:"purchase_price" bson:"purchase_price"`
	Unit                  *Unit              `json:"unit" bson:"unit"`
	PackingSpecifications float64            `json:"packing_specifications" bson:"packing_specifications"`
	Attributes            *Attributes        `json:"attributes" bson:"attributes"`
	ImageIds              *ImageIds          `json:"image_ids" bson:"image_ids"`
	Packing               int                `json:"packing" bson:"packing"`
	PurchaseApproved      bool               `json:"purchase_approved" bson:"purchase_approved"`
	OosApproved           bool               `json:"oos_approved" bson:"oos_approved"`
	AutoRecommendPo       bool               `json:"auto_recommend_po" bson:"auto_recommend_po"`
	Status                int                `json:"status" bson:"status"`
	IsSecurityDevice      bool               `json:"is_security_device" bson:"is_security_device"`
	ProductBestPrice      bool               `json:"product_best_price" bson:"product_best_price"`
	ProductSpecialDeal    bool               `json:"product_special_deal" bson:"product_special_deal"`
	ProductFlashSales     bool               `json:"product_flash_sales" bson:"product_flash_sales"`
	ProductWeekend        bool               `json:"product_weekend" bson:"product_weekend"`
	ProductSpecialMember  bool               `json:"product_special_member" bson:"product_special_member"`
	NetWeight             float64            `json:"net_weight" bson:"net_weight"`
	SkuLong               float64            `json:"sku_long" bson:"sku_long"`
	SkuWide               float64            `json:"sku_wide" bson:"sku_wide"`
	SkuHigh               float64            `json:"sku_high" bson:"sku_high"`
	RoleSku               int                `json:"role_sku" bson:"role_sku"`
}

type Variants []Variant

type Vendor struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	ExchangeTime    float64            `json:"exchange_time" bson:"exchange_time"`
	ExchangePercent float64            `json:"exchange_percent" bson:"exchange_percent"`
	MoqType         string             `json:"moq_type" bson:"moq_type"`
	MoqValue        float64            `json:"moq_value" bson:"moq_value"`
	Unit            string             `json:"unit" bson:"unit"`
	//IsDefault       bool               `json:"is_default" bson:"is_default"`
}

type Vendors []Vendor

type SubVendor struct {
	Id                       primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	ShortName                string             `json:"short_name" bson:"short_name"`
	Name                     string             `json:"name" bson:"name"`
	Code                     string             `json:"code" bson:"code"`
	ExchangeTime             float64            `json:"exchange_time" bson:"exchange_time"`
	ExchangePercent          float64            `json:"exchange_percent" bson:"exchange_percent"`
	IsExchange               float64            `json:"is_exchange" bson:"is_exchange"`
	ExchangeAt               float64            `json:"exchange_at" bson:"exchange_at"`
	MoqType                  string             `json:"moq_type" bson:"moq_type"`
	MoqValue                 float64            `json:"moq_value" bson:"moq_value"`
	Unit                     string             `json:"unit" bson:"unit"`
	DirectDiscountPercentage float64            `json:"direct_discount_percentage" bson:"direct_discount_percentage"`
	Calculation              int                `json:"calculation" bson:"calculation"`
	IsDefault                bool               `json:"is_default" bson:"is_default"`
	MoqJump                  int                `json:"moq_jump" bson:"moq_jump"`
}

type SubVendors []SubVendor

type Element struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Cost     float64            `json:"cost" bson:"cost"`
	Quantity float64            `json:"quantity" bson:"quantity"`
	Unit     *Unit              `json:"unit" bson:"unit"`
	IsReused bool               `json:"is_reused" bson:"is_reused"`
}

type Elements []Element

type Product struct {
	Id                    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name                  string             `json:"name" bson:"name"`
	NameNoneAccent        string             `json:"name_none_accent" bson:"name_none_accent"`
	Description           string             `json:"description" bson:"description"`
	DescriptionNoneAccent string             `json:"description_none_accent" bson:"description_none_accent"`
	CategoryId            primitive.ObjectID `json:"category_id" bson:"category_id"`
	CategorySecondaryId   primitive.ObjectID `json:"category_secondary_id" bson:"category_secondary_id"`
	ProductType           string             `json:"product_type,omitempty" bson:"product_type,omitempty"`
	ItemClassId           primitive.ObjectID `json:"item_class_id" bson:"item_class_id"`
	Status                bool               `json:"status" bson:"status"`
	BaseMeasure           string             `json:"base_measure" bson:"base_measure"`
	BaseMeasureValue      float64            `json:"base_measure_value" bson:"base_measure_value"`
	ExprireDateQuanity    float64            `json:"exprire_date_quanity" bson:"exprire_date_quanity"`
	Temperature           string             `json:"temperature" bson:"temperature"`
	CreatedBy             primitive.ObjectID `json:"created_by" bson:"created_by"`
	ModifiedBy            primitive.ObjectID `json:"modified_by" bson:"modified_by"`
	Attributes            *Attributes        `json:"attributes" bson:"attributes"`
	Images                *Images            `json:"images" bson:"images""`
	BaseUnit              *BaseUnit          `json:"base_unit" bson:"base_unit"`
	ProductOrigin         *ProductOrigin     `json:"product_origin" bson:"product_origin"`
	ActiveAllBranch       bool               `json:"active_all_branch" bson:"active_all_branch"`
	DeactiveBranchs       *DeactiveBranchs   `json:"deactive_branchs" bson:"deactive_branchs"`
	Variants              *Variants          `json:"variants" bson:"variants"`
	SubVendors            *SubVendors        `json:"sub_vendors" bson:"sub_vendors"`
	Elements              *Elements          `json:"elements" bson:"elements"`
	Vendors               *Vendors           `json:"vendors" bson:"vendors"`
	TaxCode               string             `json:"tax_code" bson:"tax_code"`
	DisplayZone           string             `json:"display_zone" bson:"display_zone"`
	FrequencyControl      int                `json:"frequency_control" bson:"frequency_control"`
	NormDisplay           int                `json:"norm_display" bson:"norm_display"`
	MinStockLevel         int                `json:"min_stock_level" bson:"min_stock_level"`
	MaxStockLevel         int                `json:"max_stock_level" bson:"max_stock_level"`
	ToOrderFormula        int                `json:"to_order_formula" bson:"to_order_formula"`
	ExpiredRemainPercent  float64            `json:"expired_remain_percent" bson:"expired_remain_percent"`
	Note                  string             `json:"note" bson:"note"`
	IsInTransit           bool               `json:"is_in_transit" bson:"is_in_transit"`
	Hash                  string             `json:"hash" bson:"hash"`
	CategoryPathCode      string             `json:"category_path_code" bson:"category_path_code"`
	CategoryPathName      string             `json:"category_path_name" bson:"category_path_name"`
	CategoryPathId        string             `json:"category_path_id" bson:"category_path_id"`
}

func (Product) GetProductDatabase() string {
	return "kf_products"
}

func (Product) GetProductCollection() string {
	return "kf_products"
}

type ProductImport struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
