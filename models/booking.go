package models

import "time"


type Booking struct {
	ID                  uint`json:"id"`
	ProductID           uint`json:"productId"`
	CheckInTime         time.Time`json:"checkInTime"`
	CheckOutTime        time.Time`json:"checkOutTime"`
	NumberOfAdults      int`json:"numberOfAdults"`
	NumberOfChildren    int`json:"numberOfChildren"`
	SpecialNotes        string`json:"specialNotes"`
	UserID              int`json:"userId"`
}

func (booking *Booking) SetBooking(bookingInput Booking) {
	booking.ProductID = bookingInput.ProductID
	booking.CheckInTime = bookingInput.CheckInTime
	booking.CheckOutTime = bookingInput.CheckOutTime
	booking.NumberOfAdults = bookingInput.NumberOfAdults
	booking.NumberOfChildren = bookingInput.NumberOfChildren
	booking.SpecialNotes = bookingInput.SpecialNotes
	booking.UserID = bookingInput.UserID
}
