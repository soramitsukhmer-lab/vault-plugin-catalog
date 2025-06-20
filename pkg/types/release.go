package types

type ReleaseSpec struct {
	Linux  ReleasePlatformSpec `json:"linux,omitempty"`
	Darwin ReleasePlatformSpec `json:"darwin,omitempty"`
}

type ReleasePlatformSpec struct {
	Amd64 ReleaseArchitectureSpec `json:"amd64,omitempty"`
	Arm64 ReleaseArchitectureSpec `json:"arm64,omitempty"`
}

type ReleaseArchitectureSpec struct {
	Url    string `json:"url"`
	Sha256 string `json:"sha256"`
}
