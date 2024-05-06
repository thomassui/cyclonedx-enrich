package models

type RuleEntry struct {
	Rule       string
	Licenses   []string
	Properties map[string]string
	Hashes     map[string]string
	References []Reference
}
