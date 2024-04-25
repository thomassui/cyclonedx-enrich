package models

type RuleEntry struct {
	Rule       string
	Licenses   []string
	Properties map[string]string
	Hashes     map[string]string
	References []RuleReference
}

type RuleReference struct {
	URL     string
	Type    string
	Comment string
}
