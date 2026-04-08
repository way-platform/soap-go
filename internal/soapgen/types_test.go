package soapgen

import "testing"

func TestToGoName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// Basic XML names — casing preserved after first letter
		{"GetWeather", "GetWeather"},
		{"getWeather", "GetWeather"},
		{"", ""},

		// Underscore/hyphen/dot separators
		{"service_order", "ServiceOrder"},
		{"service-order", "ServiceOrder"},
		{"service.order", "ServiceOrder"},

		// Spaces in enum values (bug fix: "EIR Sync Error")
		{"EIR Sync Error", "EIRSyncError"},
		{"Invalid Request", "InvalidRequest"},
		{"Invalid IMEI block request", "InvalidIMEIBlockRequest"},
		{"SP requested block", "SPRequestedBlock"},
		{"Ready for EIR", "ReadyForEIR"},
		{"HUD Requery", "HUDRequery"},

		// Slash in MIME types (bug fix: "application/xml")
		{"*/*", "Value"},
		{"application/xml", "ApplicationXml"},
		{"application/atom+xml", "ApplicationAtomXml"},
		{"application/octet-stream", "ApplicationOctetStream"},
		{"application/x-www-form-urlencoded", "ApplicationXWwwFormUrlencoded"},
		{"text/html", "TextHtml"},
		{"image/jpeg", "ImageJpeg"},
		{"multipart/form-data", "MultipartFormData"},

		// Colon in enum values (bug fix: "1:M")
		{"1:M", "V1M"},
		{"1:0", "V10"},

		// Charset enum values
		{"UTF-8", "UTF8"},
		{"ISO-8859-1", "ISO88591"},
		{"US-ASCII", "USASCII"},

		// Edge cases
		{"  spaced  ", "Spaced"},
		{"OK", "OK"},
		{"ERROR", "ERROR"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toGoName(tt.input)
			if got != tt.want {
				t.Errorf("toGoName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
