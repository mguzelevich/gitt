package parser

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	//"strconv"
	"strings"
)

type InitFunc func(p *OutputParser) error

type InitDscrFunc func(dscr *Descriptor) error

type SectionName string

const (
	UNKNOWN SectionName = "UNKNOWN"
)

type sectionDescriptor struct {
	reTitle      []namedRe
	reBody       []namedRe
	handlerTitle sectionHandler
	handlerBody  bodyHandler
}

type sectionHandler func(section SectionName, name string, matches []string, dscr *Descriptor) error

type bodyHandler func(name string, matches []string, dscr *Descriptor) error

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

type matchResult struct {
	reName  string
	matches []string
}

type Section struct {
	name   SectionName
	titles []matchResult
	body   []matchResult
}

type Output []Section

func (o Output) String() string {
	out := bytes.Buffer{}
	for i, s := range o {
		fmt.Fprintf(&out, "[%03d] %s\n", i, s.name)
		for _, st := range s.titles {
			fmt.Fprintf(&out, "t: %s\n", st.reName)
		}
		for _, sb := range s.body {
			fmt.Fprintf(&out, "b: %s\n", sb.reName)
		}
	}
	return out.String()
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

func (p *OutputParser) RegSection(section SectionName,
	reTitle []OutLineRE, reBody []OutLineRE,
	handlerTitle sectionHandler, handlerBody bodyHandler) {

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

func DummyTitleHandler(section SectionName, name string, matches []string, dscr *Descriptor) error {
	return nil
}

func DummyHandler(name string, matches []string, dscr *Descriptor) error {
	return nil
}

func (p *OutputParser) splitOutput(buffer bytes.Buffer) (Output, error) {
	sections := Output{}
	output := bytes.Split([]byte(buffer.String()), []byte{'\n'})

	currentSection := Section{
		name:   UNKNOWN,
		titles: []matchResult{},
		body:   []matchResult{},
	}
	sectionTitlesFinished := false
	for _, line := range output {
		if p.Failed {
			break
		}

		str := string(line)
		trimmed := strings.TrimSpace(str)

		reSectionTitleMatches := p.SectionTitleregexp.FindStringSubmatch(str)
		reSectionHintsMatches := p.SectionHintsRegexp.FindStringSubmatch(str)
		reSectionBodyMatches := p.SectionBodyRegexp.FindStringSubmatch(str)

		switch {
		default:
			p.Failed = true
		case len(reSectionTitleMatches) > 0:
			if sName, reName, matches, err := p.identifyTitleLine(trimmed); err != nil {
				p.Failed = true
				fmt.Fprintf(&p.DebugOutput, "\tunmatched title line [%s]\n", str)
			} else {
				if sectionTitlesFinished {
					if currentSection.name != UNKNOWN {
						sections = append(sections, currentSection)
						currentSection = Section{
							titles: []matchResult{},
							body:   []matchResult{},
						}
					}
					currentSection.name = sName
					currentSection.titles = append(currentSection.titles, matchResult{reName, matches})
					continue
				}
				switch sName {
				case UNKNOWN:
					p.Failed = true
					fmt.Fprintf(&p.DebugOutput, "\tunknown section title [%s]\n", str)
				case currentSection.name: // exist section
					currentSection.titles = append(currentSection.titles, matchResult{reName, matches})
				default: // new section
					if currentSection.name != UNKNOWN {
						sections = append(sections, currentSection)
						currentSection = Section{
							titles: []matchResult{},
							body:   []matchResult{},
						}
					}
					currentSection.name = sName
					currentSection.titles = append(currentSection.titles, matchResult{reName, matches})
				}
			}
		case len(reSectionBodyMatches) > 0:
			if reName, matches, err := p.identifyBodyLine(currentSection.name, trimmed); err != nil {
				p.Failed = true
				fmt.Fprintf(&p.DebugOutput, "\tunmatched body line [%s]\n", str)
			} else {
				currentSection.body = append(currentSection.body, matchResult{reName, matches})
			}
		case trimmed == "":
			sectionTitlesFinished = true
			continue // empty line skipped
		case len(reSectionHintsMatches) > 0:
			sectionTitlesFinished = true
			continue // hint line skipped
		}
	}
	if currentSection.name != UNKNOWN || len(currentSection.titles) > 0 || len(currentSection.body) > 0 {
		sections = append(sections, currentSection)
	}
	if p.Failed {
		return sections, errors.New("split error")
	}
	return sections, nil
}

func (p *OutputParser) identifyTitleLine(title string) (SectionName, string, []string, error) {
	for _, section := range p.sectionsSeq {
		sd := p.sections[section]
		for _, re := range sd.reTitle {
			if matches := re.re.FindStringSubmatch(title); len(matches) > 0 {
				return section, re.name, matches, nil
			}
		}
	}
	fmt.Fprintf(os.Stderr, "ERR\n")
	return UNKNOWN, "", nil, errors.New("not matched title line")
}

func (p *OutputParser) identifyBodyLine(section SectionName, line string) (string, []string, error) {
	sd := p.sections[section]
	for _, re := range sd.reBody {
		if matches := re.re.FindStringSubmatch(line); len(matches) > 0 {
			return re.name, matches, nil
		}
	}
	return "", nil, errors.New("not matched body line")
}

func (p *OutputParser) parseOutput(sections []Section, dscr *Descriptor) error {
	for i, section := range sections {
		fmt.Fprintf(&p.DebugOutput, "[%02d] section: %s\n", i, section.name)
		sd := p.sections[section.name]
		for _, title := range section.titles {
			if err := sd.handlerTitle(section.name, title.reName, title.matches, dscr); err != nil {
				fmt.Fprintf(&p.DebugOutput, "title handler error: %s\nrename=%s, matches=%s", err, title.reName, title.matches)
				return errors.New("title handler error")
			}
		}
		for _, line := range section.body {
			if err := sd.handlerBody(line.reName, line.matches, dscr); err != nil {
				fmt.Fprintf(&p.DebugOutput, "body handler error: %s\nrename=%s, matches=%s", err, line.reName, line.matches)
				return errors.New("title handler error")
			}
		}
	}
	return nil
}

func (p *OutputParser) Process(buffer bytes.Buffer, dscr *Descriptor) (*Descriptor, error) {
	p.Failed = false
	p.DebugOutput = bytes.Buffer{}
	p.DebugOutput.Write([]byte("<<<\n"))

	p.initDscr(dscr)

	out, err := p.splitOutput(buffer)
	if err != nil {
		return dscr, err
	}

	err = p.parseOutput(out, dscr)
	if err != nil {
		return dscr, err
	}
	p.DebugOutput.Write([]byte(">>>\n"))
	return dscr, nil
}

func NewParser(initFunction InitFunc, initDscrFunc InitDscrFunc) *OutputParser {
	p := new(OutputParser)
	p.sections = map[SectionName]sectionDescriptor{}
	p.sectionsSeq = []SectionName{}
	p.initDscr = initDscrFunc

	p.SectionTitleregexp = regexp.MustCompile(`^[a-zA-z]+.*$`)
	p.SectionHintsRegexp = regexp.MustCompile(`^  \(.*\)$`)
	p.SectionBodyRegexp = regexp.MustCompile(`^\t.+$`)

	initFunction(p)
	return p
}
