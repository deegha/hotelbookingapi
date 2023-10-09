package models

import (
	"time"
)

type ProductSearch struct {
	SearchQuery string`json:"searchQuery"`
	Limit               int`json:"limit"`
	Offset              int`json:"offSet"`
	DistrictId          int`json:"districtId"`
	ProvinceId          int`json:"provinceId"`
	Price               int`json:"price"`
	RatingScore         int`json:"ratingScore"`
	PropertyType        int`json:"propertyType"`
	BedPreferance       int`json:"bedPreferance"`	
	NumberOfAdults      int`json:"numberOfAdults"`
	NumberOfChildren    int`json:"numberOfChildren"`
	CheckInTime         time.Time`json:"checkInTime"`
	CheckOutTime        time.Time`json:"checkOutTime"`
}

type Product struct {
	ID uint`json:"id"`
  Title string`json:"title"`
  HotelId int`json:"hotelId"`
  Type int`json:"type"`
  Price int`json:"price"`
  ProvinceId int`json:"provinceId"`
  DistrictId int`json:"districtId"`
  CreatedBy int`json:"createdBy"`
	Media []Media`json:"media" gorm:"many2many:product_media;"`
	Status int`json:"status"`
	CreatedAt time.Time`json:"createdAt"`
	UpdatedAt time.Time`json:"updatedAt"`
}


func (product *Product) SetProduct(input Product) {
	if input.Title != "" {
		product.Title = input.Title
	}
	
	if input.HotelId != 0 {
		product.HotelId = input.HotelId
	}
	if input.Type != 0 {
		product.Type = input.Type
	}
	if input.Price != 0 {
		product.Price = input.Price
	}
  if input.ProvinceId != 0 {
		product.ProvinceId = input.ProvinceId
	}
	if input.DistrictId != 0 {
		product.DistrictId = input.DistrictId
	}

	if input.CreatedBy != 0 {
		product.CreatedBy = input.CreatedBy
	}	

	if input.Status != 0  {
		product.Status = input.Status
	}

	product.CreatedAt = time.Now()
	product.Media = input.Media
}

