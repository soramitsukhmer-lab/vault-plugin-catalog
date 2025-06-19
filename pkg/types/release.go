package types

type Releases struct {
	Linux  ReleasesPlatform `json:"linux,omitempty"`
	Darwin ReleasesPlatform `json:"darwin,omitempty"`
}

type ReleasesPlatform struct {
	Amd64 Release `json:"amd64,omitempty"`
	Arm64 Release `json:"arm64,omitempty"`
}

type Release struct {
	Url    string `json:"url"`
	Sha256 string `json:"sha256"`
}
