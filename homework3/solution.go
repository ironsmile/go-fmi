package main

import "fmt"
import "regexp"
import "strings"

type MarkdownParser struct {
	rawText string
}

func NewMarkdownParser(text string) (parser *MarkdownParser) {
	parser = new(MarkdownParser)
	parser.rawText = text
	return
}

func (mp *MarkdownParser) getThings(regexpStr string) (things []string) {
	things = getThings(regexpStr, mp.rawText)
	return
}

func (mp *MarkdownParser) Headers() (headers []string) {
	headers = merge(mp.getThings(`(?m)^(.+)\n=+$`),
		mp.getThings(`(?m)^\s*#\s+(.+?)(\s*#+)?$`))
	return
}

func (mp *MarkdownParser) SubHeadersOf(header string) (subHeaders []string) {

	h1RegExps := []string{`(?m)^(.+)\n=+$`, `(?m)^\s*#\s+(.+?)(\s*#+)?$`}
	h2RegExps := []string{`(?m)^(.+)\n-+$`, `(?m)^\s*##\s+(.+?)(\s*#+)?$`}

	reOneStr := fmt.Sprintf(`(?m)^\s*%s\s*\n=+$`, regexp.QuoteMeta(header))
	reOhetrStr := fmt.Sprintf(`(?m)^\s*#\s*%s(\s*#+)?$`, regexp.QuoteMeta(header))

	var regExps [](*regexp.Regexp)
	regExps = append(regExps, regexp.MustCompile(reOneStr))
	regExps = append(regExps, regexp.MustCompile(reOhetrStr))

	for _, regExp := range regExps {
		found := regExp.FindAllStringIndex(mp.rawText, -1)
		if found == nil {
			continue
		}
		for _, loc := range found {
			searchEnd := 0
			textAfterHeader := mp.rawText[loc[1]:]

			for _, h1RegExp := range h1RegExps {
				nextH1 := regexp.MustCompile(h1RegExp).FindStringIndex(textAfterHeader)
				if nextH1 == nil {
					continue
				}
				if searchEnd == 0 || nextH1[0] < searchEnd {
					searchEnd = nextH1[0]
				}
			}

			if searchEnd != 0 {
				textAfterHeader = textAfterHeader[:searchEnd]
			}

			for _, h2RegExp := range h2RegExps {
				subHeaders = merge(subHeaders, getThings(h2RegExp, textAfterHeader))
			}
		}
	}

	return
}

func (mp *MarkdownParser) Names() []string {
	return mp.getThings(`(?s)[^\.!?;]\s+(([А-ЯA-Z][а-яa-z]+[\s-]*){2,})`)
}

func (mp *MarkdownParser) PhoneNumbers() []string {
	return mp.getThings(`(?s)(\+?\s*[\d\(\)-]+[\d \(\)-]+)`)
}

func (mp *MarkdownParser) Links() []string {
	links_re := `([a-zA-Z]+://[a-zA-Z][\w\.\-]+\.[\w\.\-]+[a-zA-Z](:\d+)?(/[/\w\?#_&%]+)?)`
	return mp.getThings(links_re)
}

func (mp *MarkdownParser) Emails() []string {
	return mp.getThings(`([\w][\w\.\+]+@[a-zA-Z][\w\.\-]+\.[\w\.\-]+[a-zA-Z])`)
}

func (mp *MarkdownParser) GenerateTableOfContents() string {
	// Няма време :(
	return ""
}

func merge(one, other []string) []string {
	merged := make([]string, len(one))
	_ = copy(merged, one)
	for _, elem := range other {
		merged = append(merged, elem)
	}
	return merged
}

func getThings(regexpStr, rawText string) (things []string) {
	re := regexp.MustCompile(regexpStr)
	for _, match := range re.FindAllStringSubmatch(rawText, -1) {
		things = append(things, strings.Trim(match[1], "\t\n\v\f\r -"))
	}
	return
}
