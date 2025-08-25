package element_references

// PersonName represents the PersonName element
type PersonName struct {
	Value string `xml:",chardata"`
}

// PersonAge represents the PersonAge element
type PersonAge struct {
	Value int32 `xml:",chardata"`
}

// Address represents the Address element
type Address struct {
	Street  string `xml:"street"`
	City    string `xml:"city"`
	ZipCode string `xml:"zipCode"`
}

// Person represents the Person element
type Person struct {
	PersonName PersonName  `xml:"PersonName"`
	PersonAge  PersonAge   `xml:"PersonAge"`
	Address    Address     `xml:"Address"`
	PersonName *PersonName `xml:"PersonName"`
	Address    []Address   `xml:"Address"`
}
