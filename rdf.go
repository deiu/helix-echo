package helix

import ()

// import (
// 	"github.com/deiu/rdf2go"
// )

var mimeParser = map[string]string{
	"text/turtle":               "turtle",
	"application/ld+json":       "jsonld",
	"application/sparql-update": "internal",
}

var mimeSerializer = map[string]string{
	"application/ld+json": "jsonld",
	"text/turtle":         "turtle",
}

func gotRDF(ctype string) bool {
	if len(mimeParser[ctype]) > 0 {
		return true
	}
	return false
}
