package types

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RoomID     primitive.ObjectID `bson:"roomID,omitempty" json:"roomID,omitempty"`
	UserID     primitive.ObjectID `bson:"userID,omitempty" json:"userID,omitempty"`
	NumPersons int                `bson:"numPerson,omitempty" json:"numPerson,omitempty"`
	FromDate   time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate   time.Time          `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
	Canceled   bool               `bson:"canceled,omitempty" json:"canceled,omitempty"`
}

type BookRoomParams struct {
	RoomID     string `json:"roomID"`
	FromDate   string `json:"fromDate"`
	TillDate   string `json:"tillDate"`
	NumPersons int    `json:"numPersons"`
}

func (p BookRoomParams) Validate() error {
	_, err := primitive.ObjectIDFromHex(p.RoomID)
	if err != nil {
		return errors.New("invalid room identity")
	}
	formDate, err := ParseStringToTime(p.FromDate)
	if err != nil {
		return errors.New("you should check out the from-date")
	}
	tillDate, err := ParseStringToTime(p.TillDate)
	if err != nil {
		return errors.New("you should check out the till-date")
	}
	now := time.Now()
	if now.After(formDate) || now.After(tillDate) {
		return errors.New("cannot book a room in the past")
	}
	if formDate.After(tillDate) {
		return errors.New("cannot book the room,please check out the date")
	}
	return nil
}

func NewBookingFromParams(params BookRoomParams, userID primitive.ObjectID) *Booking {
	formDate, _ := ParseStringToTime(params.FromDate)
	tillDate, _ := ParseStringToTime(params.TillDate)
	roomID, _ := primitive.ObjectIDFromHex(params.RoomID)
	return &Booking{
		UserID:     userID,
		RoomID:     roomID,
		TillDate:   tillDate,
		FromDate:   formDate,
		NumPersons: params.NumPersons,
		Canceled:   false,
	}
}

func ParseStringToTime(timestr string) (time.Time, error) {
	formatTime, err := time.Parse("2006-01-02 15:04:05", timestr)
	if err != nil {
		return time.Now(), err
	}
	return formatTime, nil
}
