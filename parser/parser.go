package parser

import (
	"bytes"
	"fmt"
	//	"os"
	"errors"
	"regexp"
	//"strconv"
	"strings"
)

type InitFunc func(p *OutputParser) error

type InitDscrFunc func(dscr *Descriptor) error

type sectionHandler func(name string, matches []string, dscr *Descriptor) error

type SectionName string

const (
	UNKNOWN SectionName = "UNKNOWN"
)

type sectionDescriptor struct {
	reTitle      []namedRe
	reBody       []namedRe
	handlerTitle sectionHandler
	handlerBody  sectionHandler
}

type OutputParser struct {
	Failed      bool
	DebugOutput bytes.Buffer

	initDscr InitDscrFunc

	SectionTitleregexp *regexp.Regexp
	SectionHintsRegexp *regexp.Regexp
	SectionBodyRegexp  *regexp.Regexp

	sectionsSeq []SectionName
	sections    map[SectionName]sectionDescriptor
}

type OutSection struct {
	titles []string
	body   []string
}

type OutLineRE struct {
	name string
	re   string
}

type namedRe struct {
	name string
	re   *regexp.Regexp
}

func NewRE(name string, re string) OutLineRE {
	return OutLineRE{name, re}
}

func (p *OutputParser) RegSection(section SectionName, reTitle []OutLineRE, reBody []OutLineRE, handlerTitle sectionHandler, handlerBody sectionHandler) {
	sd := sectionDescriptor{
		reTitle:      []namedRe{},
		reBody:       []namedRe{},
		handlerTitle: handlerTitle,
		handlerBody:  handlerBody,
	}
	for _, l := range reTitle {
		sd.reTitle = append(sd.reTitle, namedRe{l.name, regexp.MustCompile(l.re)})
	}
	for _, l := range reBody {
		sd.reBody = append(sd.reBody, namedRe{l.name, regexp.MustCompile(l.re)})
	}
	p.sections[section] = sd
	p.sectionsSeq = append(p.sectionsSeq, section)
}

func DummyHandler(name string, matches []string, dscr *Descriptor) error {
	return nil
}

func (p *OutputParser) splitOutput(buffer bytes.Buffer) ([]OutSection, error) {
	outSections := []OutSection{}
	output := bytes.Split([]byte(buffer.String()), []byte{'\n'})

	outSection := OutSection{
		titles: []string{},
		body:   []string{},
	}

	titlesClosed := false
	for _, line := range output {
		str := string(line)
		trimmed := strings.TrimSpace(str)

		reSectionTitleMatches := p.SectionTitleregexp.FindStringSubmatch(str)
		reSectionHintsMatches := p.SectionHintsRegexp.FindStringSubmatch(str)
		reSectionBodyMatches := p.SectionBodyRegexp.FindStringSubmatch(str)

		switch {
		default:
			p.Failed = true
			fmt.Fprintf(&p.DebugOutput, "\tincorrect body line [%s]\n", str)
		case len(reSectionTitleMatches) > 0:
			if titlesClosed {
				outSections = append(outSections, outSection)
				outSection = OutSection{
					titles: []string{},
					body:   []string{},
				}
				titlesClosed = false
			}
			classes := p.classifyTitleLine(trimmed)
			cnt := len(classes)
			if cnt == 0 {
				p.Failed = true
				fmt.Fprintf(&p.DebugOutput, "\tsection's title not matched to any section [%s]\n", str)
			} else if cnt == 1 {
				outSection.titles = append(outSection.titles, trimmed)
			} else if cnt > 1 {
				outSection.titles = append(outSection.titles, trimmed)
			}
			// // new section start matched
			// if len(sectionStrings) > 0 {
			// 	outSections = append(outSections, sectionStrings)
			// 	sectionStrings = []string{}
			// }
			// sectionStrings = append(sectionStrings, trimmed)
		case trimmed == "":
			titlesClosed = true
			continue // empty line skipped
		case len(reSectionHintsMatches) > 0:
			titlesClosed = true
			continue // hint line skipped
		case len(reSectionBodyMatches) > 0:
			outSection.body = append(outSection.body, trimmed)
		}
	}
	if len(outSection.titles) > 0 || len(outSection.body) > 0 {
		outSections = append(outSections, outSection)
	}
	if p.Failed {
		//fmt.Fprintf(&p.DebugOutput, "\tsection [%s]\n", section)
		return outSections, errors.New("split error")
	}
	return outSections, nil
}

func (p *OutputParser) classifyTitleLine(title string) []SectionName {
	matchedSections := []SectionName{}
	for _, name := range p.sectionsSeq {
		sd := p.sections[name]
		for _, re := range sd.reTitle {
			if matches := re.re.FindStringSubmatch(title); len(matches) > 0 {
				matchedSections = append(matchedSections, name)
			}
		}
	}
	return matchedSections
}

func (p *OutputParser) identifySection(titles []string) (SectionName, []string) {
	matchedSection := UNKNOWN
	for _, title := range titles {
		for _, name := range p.sectionsSeq {
			sd := p.sections[name]
			matchingStarted := false
			for _, re := range sd.reTitle {
				if matches := re.re.FindStringSubmatch(title); len(matches) > 0 {
					matchingStarted = true
					matchedSection = name
					return name, []string{}
				}
			}
			if matchingStarted {

			}
		}
	}
	return matchedSection, []string{}
}

func (p *OutputParser) parseOutput(outputSections []OutSection, dscr *Descriptor) error {
	for i, section := range outputSections {
		sName, matches := p.identifySection(section.titles)
		fmt.Fprintf(&p.DebugOutput, "[%02d] section: %s\n", i, sName)
		if sName == UNKNOWN {
			p.Failed = true
			fmt.Fprintf(&p.DebugOutput, "\tunknown section: %s\n", section.titles)
			continue
		}
		s := p.sections[sName]
		s.handlerTitle(string(sName), matches, dscr)
		for _, bodyLine := range section.body {
			matched := false
			if s.reBody == nil {
				break
			}
			for _, re := range s.reBody {
				if matches := re.re.FindStringSubmatch(bodyLine); len(matches) > 0 {
					s.handlerBody(re.name, matches, dscr)
					matched = true
					break
				}
			}
			if !matched {
				p.Failed = true
				fmt.Fprintf(&p.DebugOutput, "\tincorrect body line [%s]\n", bodyLine)
			}
		}
		if p.Failed {
			fmt.Fprintf(&p.DebugOutput, "     section titles:\n")
			for _, l := range section.body {
				fmt.Fprintf(&p.DebugOutput, "\t%s\n", l)
			}
			fmt.Fprintf(&p.DebugOutput, "     section body:\n")
			for _, l := range section.body {
				fmt.Fprintf(&p.DebugOutput, "\t%s\n", l)
			}
		}
	}
	return nil
}
func (p *OutputParser) Process(buffer bytes.Buffer, dscr *Descriptor) (*Descriptor, error) {
	p.DebugOutput = bytes.Buffer{}
	p.DebugOutput.Write([]byte("<<<\n"))

	p.initDscr(dscr)

	out, err := p.splitOutput(buffer)
	if err != nil {
		fmt.Printf(p.DebugOutput.String())
		panic(err)
	}
	err = p.parseOutput(out, dscr)
	if err != nil {
		fmt.Printf(p.DebugOutput.String())
		panic(err)
	}
	p.DebugOutput.Write([]byte(">>>\n"))
	return dscr, nil
}

func NewParser(initFunction InitFunc, initDscrFunc InitDscrFunc) *OutputParser {
	p := new(OutputParser)
	p.Failed = false
	p.sections = map[SectionName]sectionDescriptor{}
	p.sectionsSeq = []SectionName{}
	p.initDscr = initDscrFunc

	p.SectionTitleregexp = regexp.MustCompile(`^[a-zA-z]+.*$`)
	p.SectionHintsRegexp = regexp.MustCompile(`^  \(.*\)$`)
	p.SectionBodyRegexp = regexp.MustCompile(`^\t.+$`)

	initFunction(p)
	return p
}
