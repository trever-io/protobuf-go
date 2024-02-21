package internal_gengo

import (
	"log"
	"regexp"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

var commentPqTypesRe *regexp.Regexp

func init() {
	var err error
	commentPqTypesRe, err = regexp.Compile(`(\s?|//)@pq_type\s`)
	if err != nil {
		log.Fatalf("compile comment pq types regexp failed. %s", err)
	}
}

func fieldPqType(currentGoType string, currentPointer bool, tailComment protogen.Comments) (goType string, pointer bool, newTailing protogen.Comments) {
	newTailing = tailComment

	if !commentPqTypesRe.MatchString(string(tailComment)) {
		return currentGoType, currentPointer, newTailing
	}

	newTailing = protogen.Comments(strings.Replace(string(tailComment), "@pq_type", "", 1))

	switch currentGoType {
	case "[]string":
		return "pq.StringArray", false, newTailing
	default:
		return currentGoType, currentPointer, newTailing
	}
}
