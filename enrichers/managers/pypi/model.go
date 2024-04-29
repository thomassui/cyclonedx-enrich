package pypi

import "time"

type Package struct {
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
	CommentText       string      `json:"comment_text"`
	Digests           Digests     `json:"digests"`
	Downloads         int64       `json:"downloads"`
	Filename          string      `json:"filename"`
	HasSig            bool        `json:"has_sig"`
	Md5Digest         string      `json:"md5_digest"`
	Packagetype       string      `json:"packagetype"`
	PythonVersion     string      `json:"python_version"`
	RequiresPython    interface{} `json:"requires_python"`
	Size              int64       `json:"size"`
	UploadTime        time.Time   `json:"upload_time"`
	UploadTimeISO8601 time.Time   `json:"upload_time_iso_8601"`
	URL               string      `json:"url"`
	Yanked            bool        `json:"yanked"`
	YankedReason      interface{} `json:"yanked_reason"`
}

type Digests struct {
	Blake2B256 string `json:"blake2b_256"`
	Md5        string `json:"md5"`
	Sha256     string `json:"sha256"`
}
