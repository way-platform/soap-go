package examples

//go:generate go install -C ../cmd/soap .
//go:generate soap gen -i NumberConversion.wsdl -d numberconversion
//go:generate soap gen -i GlobalWeather.wsdl -d globalweather
//go:generate soap gen -i KitchenSink.wsdl -d kitchensink
