package licenses

import (
	"log/slog"

	"github.com/CycloneDX/cyclonedx-go"
)

var log = slog.Default()

func skip(component *cyclonedx.Component) bool {
	if len(component.PackageURL) == 0 || component.Licenses != nil || (component.Licenses == nil && len(*component.Licenses) == 0) {
		//SKIP
		return true
	}
	return false
}
