package cocoapods

type Package struct {
	Name           string     `json:"name"`
	Version        string     `json:"version"`
	License        *string    `json:"license"`
	Summary        string     `json:"summary"`
	Homepage       *string    `json:"homepage"`
	Authors        *Authors   `json:"authors"`
	SocialMediaURL *string    `json:"social_media_url"`
	Source         *Source    `json:"source"`
	Platforms      *Platforms `json:"platforms"`
	SourceFiles    *string    `json:"source_files"`
	SwiftVersions  *string    `json:"swift_versions"`
	SwiftVersion   *string    `json:"swift_version"`
}

type Authors struct {
	RobertPayne string `json:"Robert Payne"`
}

type Platforms struct {
	Ios  string `json:"ios"`
	Osx  string `json:"osx"`
	Tvos string `json:"tvos"`
}

type Source struct {
	Git string `json:"git"`
	Tag string `json:"tag"`
}
