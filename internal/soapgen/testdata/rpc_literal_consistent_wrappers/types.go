package rpc_literal_consistent_wrappers

import (
	"encoding/xml"
	"time"
)

// AuthenticateWrapper represents the authenticate element
type AuthenticateWrapper struct {
	XMLName  xml.Name `xml:"http://example.com/rpc-literal-test authenticate"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

// AuthenticateResponseWrapper represents the authenticateResponse element
type AuthenticateResponseWrapper struct {
	XMLName xml.Name  `xml:"http://example.com/rpc-literal-test authenticateResponse"`
	Token   string    `xml:"token"`
	Expires time.Time `xml:"expires"`
}

// FetchDataWrapper represents the fetchData element
type FetchDataWrapper struct {
	XMLName xml.Name `xml:"http://example.com/rpc-literal-test fetchData"`
	Token   string   `xml:"token"`
	DataId  int64    `xml:"dataId"`
}

// FetchDataResponseWrapper represents the fetchDataResponse element
type FetchDataResponseWrapper struct {
	XMLName      xml.Name  `xml:"http://example.com/rpc-literal-test fetchDataResponse"`
	Data         string    `xml:"data"`
	LastModified time.Time `xml:"lastModified"`
}

// SystemStatus represents the SystemStatus element
type SystemStatus struct {
	XMLName xml.Name `xml:"SystemStatus"`
	Online  bool     `xml:"online"`
	Version string   `xml:"version"`
}
