package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func loadTheReadme() string {
	content, err := ioutil.ReadFile("./README.md")
	if err != nil {
		return ""
	}
	return string(content)
}

func TestHeaders(t *testing.T) {
	mdParser := NewMarkdownParser(loadTheReadme())
	headers := mdParser.Headers()

	if len(headers) < 1 {
		t.Fatalf("Parser found no headers where there were some")
	}

	if headers[0] != "MarkdownParser" {
		t.Fatalf("Expected first header to be MarkdownParser but it was %s", headers[0])
	}

	mdParser = NewMarkdownParser("Lalala\nSomething\n==-===\nother")
	headers = mdParser.Headers()

	if len(headers) != 0 {
		t.Errorf("There should have been no parsed headers")
	}

	valid_h1_syntaxes := []string{
		"Lalala\nSomething\n======\nother",
		"Lalala\nSomething\n=\nother",
		"Something\n======",
		"Something\n======\n",
		"# Something",
		"# Something #",
		"# Something\n",
		"# Something #\n",
		"# Something ####\n",
		"# Something ####\n",
		"#Something",
		"lala\n# Something",
		"lala\n# Something\nlala",
	}

	for _, mdText := range valid_h1_syntaxes {
		mdParser = NewMarkdownParser(mdText)
		headers = mdParser.Headers()

		if len(headers) != 1 {
			t.Errorf("Did not parse any headers when parsing %s: %d", mdText,
				len(headers))
			continue
		}

		if headers[0] != "Something" {
			t.Errorf("Wrong parsed headers when parsing `%s`: `%s`", mdText,
				headers[0])
		}
	}

	mdParser = NewMarkdownParser("Lalala\nSomething\n-\nother")
	headers = mdParser.Headers()

	if len(headers) != 0 {
		t.Errorf("Parsed headers where it should haven't")
	}

	mdParser = NewMarkdownParser("")
	headers = mdParser.Headers()

	if len(headers) != 0 {
		t.Errorf("There should have been no parsed headers")
	}
}

func TestSubHeadersOf(t *testing.T) {
	mdParser := NewMarkdownParser(loadTheReadme())
	subHeaders := mdParser.SubHeadersOf("MarkdownParser")

	if len(subHeaders) < 1 {
		t.Fatalf("Parser found no sub headers where there were some")
	}

	if subHeaders[0] != "type MarkdownParser" {
		t.Fail()
	}
}

func TestTableOfContents(t *testing.T) {
	mdParser := NewMarkdownParser(loadTheReadme())
	tableOfContents := mdParser.GenerateTableOfContents()

	splitted := strings.Split(tableOfContents, "\n")
	expected := "1.1.3 `func (mp *MarkdownParser) SubHeadersOf(header string) []string`"

	if len(splitted) < 5 {
		t.Fatalf("Expected to parse at least 5 entries in table of contents")
	}

	if splitted[4] != expected {
		t.Fail()
	}
}

func TestNames(t *testing.T) {
	mdParser := NewMarkdownParser(`Super Meat Boy. В това изречение ще намерите името
Иван Попов.
Димитър Иванов не е име! Димитър също! Но пък Иван Павлов е хубаво име, Георги Кранев
също. Хуави неща може да се кажат за Иван Ковачев Павлов, но и за
Едсон Арантес Ду Насименто - Пеле! Mozilla Firefox не е име, но Mozilla Firefox е! Ходи
ги разбери! Ааа не, еднакви са. Честно! Не е като да има Кирилица из между тях, като
в Mоzillа Firеfox да кажем. Абе, Димитър беше ли име? Кажете ми дали Димитър е име, моля!
Друго интересно е дали как ще се с Жан Пиер - и подобни. Ами Ц М не е име, пък!
`)
	names := mdParser.Names()

	expectedNames := []string{
		"Meat Boy",
		"Иван Попов",
		"Иван Павлов",
		"Георги Кранев",
		"Иван Ковачев Павлов",
		"Едсон Арантес Ду Насименто - Пеле",
		"Mozilla Firefox",
		"Mоzillа Firеfox",
		"Жан Пиер",
	}

	if len(names) != len(expectedNames) {
		t.Errorf("Number of names (%d) differ than expected (%d)", len(names),
			len(expectedNames))
	}

	for _, name := range expectedNames {
		if contains(names, name) {
			continue
		}
		t.Errorf("`%s` was not among the found names as it should have been", name)
	}
}

func contains(haystack []string, needle string) bool {
	for _, elem := range haystack {
		if elem == needle {
			return true
		}
	}
	return false
}

func TestPhones(t *testing.T) {
	mdParser := NewMarkdownParser(`Някакви телефонн нормера 0889123456.
Още един +359889123456. Тук имаме още два: (089) 123-456 и 0 (889) 123 - 456.
Не може да пропуснем +4531223 2332 123, както и 123 3456 621.
`)
	phones := mdParser.PhoneNumbers()

	expectedPhones := []string{
		"0889123456",
		"+359889123456",
		"(089) 123-456",
		"0 (889) 123 - 456",
		"+4531223 2332 123",
		"123 3456 621",
	}

	if len(phones) != len(expectedPhones) {
		t.Errorf("Number of phones (%d) differ than expected (%d)", len(phones),
			len(expectedPhones))
	}

	for _, phone := range expectedPhones {
		if contains(phones, phone) {
			continue
		}
		t.Errorf("`%s` was not among the found phones as it should have been", phone)
	}
}
