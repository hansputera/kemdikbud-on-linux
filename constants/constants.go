package constants

const (
	DAPO_UNDUHAN     = "https://dapo.dikdasmen.go.id/unduhan"
	DAPO_CDN_RELEASE = "https://cdn-dapodik.kemdikbud.go.id/rilis/"
)

const (
	APP_VERSION = "0.0.1"
)

var FILE_SIGNATURES = map[string]string{
	"7z\xBC\xAF\x27\x1C": "7z",
	"Rar!\x1A\x07\x00":   "rar",
	"Nullsoft":           "nsis",
	"Inno Setup":         "inno",
	"MSI ":               "msi",
	"ISc(":               "installshield",
}
