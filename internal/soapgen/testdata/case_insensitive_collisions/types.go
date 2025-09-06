package case_insensitive_collisions

import (
	"encoding/xml"
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

// UserRequest represents the UserRequest element
type UserRequest struct {
	XMLName xml.Name        `xml:"UserRequest"`
	Value   UserRequestType `xml:",chardata"`
}

// UserRequestElement represents the userRequest element
type UserRequestElement struct {
	XMLName xml.Name        `xml:"userRequest"`
	Value   UserRequestType `xml:",chardata"`
}

// GetFleetResponse represents the GetFleetResponse element
type GetFleetResponse struct {
	XMLName xml.Name      `xml:"GetFleetResponse"`
	Value   FleetDataType `xml:",chardata"`
}

// GetFleetResponseElement represents the getFleetResponse element
type GetFleetResponseElement struct {
	XMLName xml.Name      `xml:"getFleetResponse"`
	Value   FleetDataType `xml:",chardata"`
}

// DataRecord represents the DataRecord element
type DataRecord struct {
	XMLName xml.Name   `xml:"DataRecord"`
	Value   RecordType `xml:",chardata"`
}

// DataRecordElement represents the dataRecord element
type DataRecordElement struct {
	XMLName xml.Name   `xml:"dataRecord"`
	Value   RecordType `xml:",chardata"`
}
