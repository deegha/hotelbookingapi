package models

type Amenities struct {
	ID uint`json:"id"`
	ProductID int`json:"productId"`
	Name string`json:"name"`
}

func (ameniti *Amenities) SetAmenities(newAmenity Amenities) {
  ameniti.Name = newAmenity.Name
}
