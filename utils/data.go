package utils

import (
	"cyclonedx-enrich/models"

	"github.com/CycloneDX/cyclonedx-go"
)

var RulesEmpty = []models.RuleEntry{}
var Rules = []models.RuleEntry{
	{Rule: ".*com.example.*", Licenses: []string{"Example License"}},
	{Rule: ".*io.example.*", Licenses: []string{"MIT License"}},
	{Rule: "pkg:maven/com.sun.xml.ws/release-documentation@2.3.2.*", Licenses: []string{"CDDL", "GPL-2.0"}},
	{Rule: "pkg:maven/com.sun.xml.ws/samples@2.3.2.*", Licenses: []string{"CDDL", "GPL-2.0"}},
	{Rule: "pkg:maven/javax.servlet/servlet-api@2.5.*", Licenses: []string{"Apache-2.0"}},
	{Rule: "pkg:maven/org.antlr/antlr-runtime@3.4.*", Licenses: []string{"BSD-1-Clause"}},
	{Rule: "pkg:maven/org.javamoney/moneta@1.4.*", Licenses: []string{"Apache-2.0"}},
	{Rule: "pkg:maven/org.openapitools/jackson-databind-nullable@0.2.2.*", Licenses: []string{"Apache-2.0"}},
	{Rule: "pkg:maven/com.sun.xml.ws/jaxws-ri@2.3.2", Licenses: []string{"Apache 2.0", "EDL 1.0"},
		Hashes: map[string]string{
			"SHA-1":   "15ee0d893c452db7c4865eacb37cf5355c6c5cd7",
			"SHA-256": "b3f164c170d7a281a5b869b0ead0fbca3f3c9e06a01b1be521460816eef861e5",
			"SHA-512": "3cc8334f21ca594e6b62d5d18e6219bb781d2a51edd790bc981b80bcd701310ad4cb83541e19dd17214ff61d8c2fda09c83b2c6473da3a2360c770e8fdad53ce",
			"SHA-384": "52b81b209ec0b5f4920420203fff8911f7d61defcbcef588debdfa85d2dc22f8d1126d02905092e7131a23698afec639",
		}, Properties: map[string]string{
			"cdx:npm:package:path":  "node_modules/@angular/cdk/node_modules/parse5",
			"cdx:npm:package:path2": "node_modules/@angular/cdk/node_modules/parse52",
		}, References: []models.Reference{
			{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "website"},
			{URL: "https://oss.sonatype.org/service/local/staging/deploy/maven2/", Type: "distribution"},
			{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "vcs"},
		}},
}

var ComponentEmpty = &cyclonedx.Component{}

var ComponentWithData = &cyclonedx.Component{
	Type:    "library",
	BOMRef:  "pkg:maven/com.sun.xml.ws/jaxws-ri@2.3.2?type=pom",
	Group:   "com.sun.xml.ws",
	Name:    "jaxws-ri",
	Version: "2.3.2",
	Scope:   "required",
	Hashes: &[]cyclonedx.Hash{
		{Algorithm: "MD5", Value: "a0cbf1399c88a0435be995c0f68a3baa"},
		{Algorithm: "SHA-1", Value: "15ee0d893c452db7c4865eacb37cf5355c6c5cd7"},
		{Algorithm: "SHA-256", Value: "b3f164c170d7a281a5b869b0ead0fbca3f3c9e06a01b1be521460816eef861e5"},
		{Algorithm: "SHA-512", Value: "3cc8334f21ca594e6b62d5d18e6219bb781d2a51edd790bc981b80bcd701310ad4cb83541e19dd17214ff61d8c2fda09c83b2c6473da3a2360c770e8fdad53ce"},
		{Algorithm: "SHA-384", Value: "52b81b209ec0b5f4920420203fff8911f7d61defcbcef588debdfa85d2dc22f8d1126d02905092e7131a23698afec639"},
	},
	Licenses: &cyclonedx.Licenses{
		cyclonedx.LicenseChoice{License: &cyclonedx.License{Name: "Apache 2.0"}},
		cyclonedx.LicenseChoice{License: &cyclonedx.License{Name: "EDL 1.0"}},
	},
	PackageURL: "pkg:maven/com.sun.xml.ws/jaxws-ri@2.3.2?type=pom",
	ExternalReferences: &[]cyclonedx.ExternalReference{
		{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "website"},
		{URL: "https://oss.sonatype.org/service/local/staging/deploy/maven2/", Type: "distribution"},
		{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "vcs"},
	},
	Properties: &[]cyclonedx.Property{
		{Name: "cdx:npm:package:path", Value: "node_modules/@angular/cdk/node_modules/parse5"},   //TODO: GET BETTER DATA
		{Name: "cdx:npm:package:path2", Value: "node_modules/@angular/cdk/node_modules/parse52"}, //TODO: GET BETTER DATA
	},
}

var ComponentWithoutData = &cyclonedx.Component{
	Type:       "library",
	BOMRef:     "pkg:maven/com.sun.xml.ws/jaxws-ri@2.3.2?type=pom",
	Group:      "com.sun.xml.ws",
	Name:       "jaxws-ri",
	Version:    "2.3.2",
	Scope:      "required",
	PackageURL: "pkg:maven/com.sun.xml.ws/jaxws-ri@2.3.2?type=pom",
}

var ComponentNotInDatabase = &cyclonedx.Component{
	Type:        "library",
	BOMRef:      "pkg:maven/org.openapitools/jackson-databind-nullable@0.2.4?type=jar",
	Group:       "org.openapitools",
	Name:        "jackson-databind-nullable",
	Version:     "0.2.4",
	Description: "JsonNullable wrapper class and Jackson module to support fields with meaningful null values.",
	Scope:       "required",
	Hashes: &[]cyclonedx.Hash{
		{Algorithm: "MD5", Value: "479311558bbca63453f8a79e2735aec1"},
		{Algorithm: "SHA-1", Value: "371a38c3d339833edb1b2a0d96c3d249a890bcc4"},
		{Algorithm: "SHA-256", Value: "22c73f6c44eb65cb2ebbd9a0ace61a3951cc259fdc29b89e31a80564cd116ad6"},
		{Algorithm: "SHA-512", Value: "41a4c682635a481f78602087a83a7bbd1f36c0fd8d8fe5daf2ab05907472ca2f345de086fa56bab2d554412f2a1546ec5a2e832e04b1751ba29e6612318b42dc"},
		{Algorithm: "SHA-384", Value: "740ff354152ae7d691590c75d9c0be6decbb18912f56e3aca86243b7e9f5c350c48ca0e97fb3e031aa8aaf82c49e0885"},
	},
	Licenses: &cyclonedx.Licenses{
		cyclonedx.LicenseChoice{License: &cyclonedx.License{ID: "Apache-2.0", URL: "https://www.apache.org/licenses/LICENSE-2.0"}},
	},
	PackageURL: "pkg:maven/org.openapitools/jackson-databind-nullable@0.2.4?type=jar",
	ExternalReferences: &[]cyclonedx.ExternalReference{
		{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "website"},
		{URL: "https://oss.sonatype.org/service/local/staging/deploy/maven2/", Type: "distribution"},
		{URL: "https://github.com/OpenAPITools/jackson-databind-nullable", Type: "vcs"},
	},
	Properties: &[]cyclonedx.Property{
		{Name: "cdx:npm:package:path", Value: "node_modules/@angular/cdk/node_modules/parse5"},   //TODO: GET BETTER DATA
		{Name: "cdx:npm:package:path2", Value: "node_modules/@angular/cdk/node_modules/parse52"}, //TODO: GET BETTER DATA
	},
}

var ComponentMaven = &cyclonedx.Component{
	Type:        "library",
	BOMRef:      "pkg:maven/org.openapitools/jackson-databind-nullable@0.2.4?type=jar",
	Group:       "org.openapitools",
	Name:        "jackson-databind-nullable",
	Version:     "0.2.4",
	Description: "JsonNullable wrapper class and Jackson module to support fields with meaningful null values.",
	Scope:       "required",
	PackageURL:  "pkg:maven/org.openapitools/jackson-databind-nullable@0.2.4?type=jar",
}

var ComponentCocoapods = &cyclonedx.Component{
	Type:        "library",
	BOMRef:      "pkg:cocoapods/AppAuth@1.6.2#Core",
	Name:        "AppAuth/Core",
	Version:     "1.6.2",
	Description: "AppAuth for iOS and macOS is a client SDK for communicating with [OAuth 2.0]\n(https://tools.ietf.org/html/rfc6749) and [OpenID Connect]\n(http://openid.net/specs/openid-connect-core-1_0.html) providers. It strives to\ndirectly map the requests and responses of those specifications, while following\nthe idiomatic style of the implementation language. In addition to mapping the\nraw protocol flows, convenience methods are available to assist with common\ntasks like performing an action with fresh tokens.\n\nIt follows the OAuth 2.0 for Native Apps best current practice\n([RFC 8252](https://tools.ietf.org/html/rfc8252)).",
	PackageURL:  "pkg:cocoapods/AppAuth@1.6.2#Core",
}

// TODO: CHANGE COMPONENT
var ComponentNpm = &cyclonedx.Component{
	Type:       "library",
	BOMRef:     "@angular/cdk@16.2.14|parse5@7.1.2",
	Name:       "parse5",
	Version:    "7.1.2",
	Scope:      "optional",
	PackageURL: "pkg:npm/parse5@7.1.2",
}

// TODO: CHANGE COMPONENT
var ComponentPypi = &cyclonedx.Component{
	Type:        "framework",
	BOMRef:      "pkg:pypi/aiohttp@3.9.5",
	Group:       "org.openapitools",
	Name:        "aiohttp",
	Version:     "3.9.5",
	Description: "Async http client/server framework (asyncio)",
	PackageURL:  "pkg:pypi/aiohttp@3.9.5",
}
