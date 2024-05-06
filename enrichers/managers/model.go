package managers

// COCOAPODS
type CocoaPodsPackage struct {
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

// NPM
type NpmPackage struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	// Author       NpmUser    `json:"author"`  #NpmUser can also be a string
	// Contributors []NpmUser   `json:"contributors"` #NpmUser can also be a string
	License     *string     `json:"license"`
	Repository  *Repository `json:"repository"`
	Keywords    []string    `json:"keywords"`
	GitHead     string      `json:"gitHead"`
	Homepage    *string     `json:"homepage"`
	Dist        Dist        `json:"dist"`
	Maintainers *[]NpmUser  `json:"maintainers"`
}

type NpmUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Dist struct {
	Integrity    string      `json:"integrity"`
	Shasum       string      `json:"shasum"`
	Tarball      string      `json:"tarball"`
	FileCount    int64       `json:"fileCount"`
	UnpackedSize int64       `json:"unpackedSize"`
	Signatures   []Signature `json:"signatures"`
}

type Signature struct {
	Keyid string `json:"keyid"`
	Sig   string `json:"sig"`
}

type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// PyPi
type PyPiPackage struct {
	Info            Info          `json:"info"`
	LastSerial      int64         `json:"last_serial"`
	Urls            []URL         `json:"urls"`
	Vulnerabilities []interface{} `json:"vulnerabilities"`
}

type Info struct {
	Author                 string        `json:"author"`
	AuthorEmail            string        `json:"author_email"`
	BugtrackURL            interface{}   `json:"bugtrack_url"`
	Classifiers            []interface{} `json:"classifiers"`
	Description            string        `json:"description"`
	DescriptionContentType interface{}   `json:"description_content_type"`
	DocsURL                interface{}   `json:"docs_url"`
	DownloadURL            string        `json:"download_url"`
	Downloads              Downloads     `json:"downloads"`
	Dynamic                interface{}   `json:"dynamic"`
	HomePage               string        `json:"home_page"`
	Keywords               string        `json:"keywords"`
	License                *string       `json:"license"`
	Maintainer             string        `json:"maintainer"`
	MaintainerEmail        string        `json:"maintainer_email"`
	Name                   string        `json:"name"`
	PackageURL             string        `json:"package_url"`
	Platform               string        `json:"platform"`
	ProjectURL             string        `json:"project_url"`
	ProjectUrls            ProjectUrls   `json:"project_urls"`
	ProvidesExtra          interface{}   `json:"provides_extra"`
	ReleaseURL             string        `json:"release_url"`
	RequiresDist           []string      `json:"requires_dist"`
	RequiresPython         string        `json:"requires_python"`
	Summary                string        `json:"summary"`
	Version                string        `json:"version"`
	Yanked                 bool          `json:"yanked"`
	YankedReason           interface{}   `json:"yanked_reason"`
}

type Downloads struct {
	LastDay   int64 `json:"last_day"`
	LastMonth int64 `json:"last_month"`
	LastWeek  int64 `json:"last_week"`
}

type ProjectUrls struct {
	Homepage string `json:"Homepage"`
}

type URL struct {
	CommentText    string      `json:"comment_text"`
	Digests        Digests     `json:"digests"`
	Downloads      int64       `json:"downloads"`
	Filename       string      `json:"filename"`
	HasSig         bool        `json:"has_sig"`
	Md5Digest      string      `json:"md5_digest"`
	Packagetype    string      `json:"packagetype"`
	PythonVersion  string      `json:"python_version"`
	RequiresPython interface{} `json:"requires_python"`
	Size           int64       `json:"size"`
	// UploadTime        time.Time   `json:"upload_time"`
	// UploadTimeISO8601 time.Time   `json:"upload_time_iso_8601"`
	URL          string      `json:"url"`
	Yanked       bool        `json:"yanked"`
	YankedReason interface{} `json:"yanked_reason"`
}

type Digests struct {
	Blake2B256 string `json:"blake2b_256"`
	Md5        string `json:"md5"`
	Sha256     string `json:"sha256"`
}
