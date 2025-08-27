# NumberConversion

**Namespace:** `http://www.dataaccess.com/webservicesserver/`

## NumberConversion

The Number Conversion Web Service, implemented with Visual DataFlex, provides functions that convert numbers into words or dollar amounts.

### NumberToWords

Returns the word corresponding to the positive number passed as parameter. Limited to quadrillions.

#### Request

- **NumberToWords** (complex)
  - **ubiNum** (xs:unsignedLong)

#### Response

- **NumberToWordsResponse** (complex)
  - **NumberToWordsResult** (xs:string)


### NumberToDollars

Returns the non-zero dollar amount of the passed number.

#### Request

- **NumberToDollars** (complex)
  - **dNum** (xs:decimal)

#### Response

- **NumberToDollarsResponse** (complex)
  - **NumberToDollarsResult** (xs:string)


