package npm

type Package struct {
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
