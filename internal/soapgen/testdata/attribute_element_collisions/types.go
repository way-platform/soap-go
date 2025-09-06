package attribute_element_collisions

import (
	"encoding/xml"
)

// Complex types

// ConfigDataType represents the ConfigDataType complex type
type ConfigDataType struct {
	ConfigID     string  `xml:"ConfigID"`
	Status       string  `xml:"Status"`
	Data         string  `xml:"Data"`
	Priority     int32   `xml:"Priority"`
	StatusAttr   *string `xml:"status,attr,omitempty"`
	PriorityAttr *int32  `xml:"priority,attr,omitempty"`
	Enabled      *bool   `xml:"enabled,attr,omitempty"`
}

// DownloadRequestType represents the DownloadRequestType complex type
type DownloadRequestType struct {
	DriverID    string        `xml:"DriverID"`
	VehicleID   string        `xml:"VehicleID"`
	Version     int32         `xml:"Version"`
	TimeRange   TimeRangeType `xml:"TimeRange"`
	Sessionid   *string       `xml:"sessionid,attr,omitempty"`
	VersionAttr *int32        `xml:"version,attr,omitempty"`
	Limit       *int32        `xml:"limit,attr,omitempty"`
	Offset      *int32        `xml:"offset,attr,omitempty"`
}

// MetadataInfoType represents the MetadataInfoType complex type
type MetadataInfoType struct {
	Name        string  `xml:"Name"`
	Type        string  `xml:"Type"`
	ID          string  `xml:"ID"`
	Description string  `xml:"Description"`
	NameAttr    *string `xml:"name,attr,omitempty"`
	TypeAttr    *string `xml:"type,attr,omitempty"`
	Id          *string `xml:"id,attr,omitempty"`
	Category    *string `xml:"category,attr,omitempty"`
}

// TimeRangeType represents the TimeRangeType complex type
type TimeRangeType struct {
	Begin time.Time `xml:"Begin"`
	End   time.Time `xml:"End"`
}

// DownloadRequest represents the DownloadRequest element
type DownloadRequest struct {
	XMLName xml.Name            `xml:"DownloadRequest"`
	Value   DownloadRequestType `xml:",chardata"`
}

// ConfigData represents the ConfigData element
type ConfigData struct {
	XMLName xml.Name       `xml:"ConfigData"`
	Value   ConfigDataType `xml:",chardata"`
}

// MetadataInfo represents the MetadataInfo element
type MetadataInfo struct {
	XMLName xml.Name         `xml:"MetadataInfo"`
	Value   MetadataInfoType `xml:",chardata"`
}
