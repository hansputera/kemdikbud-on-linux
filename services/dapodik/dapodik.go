package dapodik

type DapodikPatch struct {
	PatchName  string
	PatchUrl   string
	IsVokasi   bool
	Categories []string
}

type DapodikVersion struct {
	Version   string
	Url       string
	VokasiUrl string
	Patches   []*DapodikPatch
}
