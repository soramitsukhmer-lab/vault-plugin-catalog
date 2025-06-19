package types

type Releases struct {
	Linux  ReleasesArch `json:"linux,omitempty"`
	Darwin ReleasesArch `json:"darwin,omitempty"`
}

type ReleasesArch struct {
	Amd64 Release `json:"amd64,omitempty"`
	Arm64 Release `json:"arm64,omitempty"`
}

type Release struct {
	Url    string `json:"url"`
	Sha256 string `json:"sha256"`
}
