package examples

//go:generate go install -C ../cmd/soap .

//go:generate soap gen --client -i NumberConversion.wsdl -d numberconversion
//go:generate soap doc -i NumberConversion.wsdl -o numberconversion/README.md

//go:generate soap gen --client -i GlobalWeather.wsdl -d globalweather
//go:generate soap doc -i GlobalWeather.wsdl -o globalweather/README.md

//go:generate soap gen --client -i KitchenSink.wsdl -d kitchensink
//go:generate soap doc -i KitchenSink.wsdl -o kitchensink/README.md
