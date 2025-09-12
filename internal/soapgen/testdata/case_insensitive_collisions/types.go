package case_insensitive_collisions

import (
	"encoding/xml"
	"time"
)

// Complex types

// FleetDataType represents the FleetDataType complex type
type FleetDataType struct {
	VehicleCount int32  `xml:"vehicleCount"`
	Status       string `xml:"status"`
}

// RecordType represents the RecordType complex type
type RecordType struct {
	Id        string    `xml:"id"`
	Timestamp time.Time `xml:"timestamp"`
}

// UserRequestType represents the UserRequestType complex type
type UserRequestType struct {
	UserID      string `xml:"userID"`
	RequestType string `xml:"requestType"`
}

// UserRequestWrapper represents the UserRequest element
type UserRequestWrapper struct {
	XMLName xml.Name        `xml:"http://example.com/collisions UserRequest"`
	Value   UserRequestType `xml:",chardata"`
}

// GetFleetResponseWrapper represents the GetFleetResponse element
type GetFleetResponseWrapper struct {
	XMLName xml.Name      `xml:"http://example.com/collisions GetFleetResponse"`
	Value   FleetDataType `xml:",chardata"`
}

// DataRecord represents the DataRecord element
type DataRecord struct {
	XMLName xml.Name   `xml:"DataRecord"`
	Value   RecordType `xml:",chardata"`
}
