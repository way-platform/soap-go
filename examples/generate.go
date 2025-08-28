package examples

//go:generate go install -C ../cmd/soap .
//go:generate soap gen -i NumberConversion.wsdl -d numberconversion
//go:generate soap doc -i NumberConversion.wsdl -o NumberConversion.md
//go:generate soap gen -i GlobalWeather.wsdl -d globalweather
//go:generate soap doc -i GlobalWeather.wsdl -o GlobalWeather.md
//go:generate soap gen -i KitchenSink.wsdl -d kitchensink
//go:generate soap doc -i KitchenSink.wsdl -o KitchenSink.md
