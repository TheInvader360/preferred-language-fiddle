package main

import (
	"golang.org/x/text/language"
)

var supportedLanguages = language.NewMatcher([]language.Tag{
	language.Make("en"), // used as default if no match found
	language.Make("cy"),
})

func isPreferredCy(acceptLanguageHeader string) bool {
	tags, _, _ := language.ParseAcceptLanguage(acceptLanguageHeader) // safely ignore error (fallback to default language)
	bestMatchTag, _, _ := supportedLanguages.Match(tags...)
	return bestMatchTag.String() == "cy"
}
