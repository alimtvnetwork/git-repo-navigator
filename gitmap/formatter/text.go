// Package formatter — text.go writes a plain text file with one git clone command per line.
package formatter

import (
	"fmt"
	"io"

	"github.com/user/gitmap/constants"
	"github.com/user/gitmap/model"
)

// WriteText writes one git clone instruction per line to a plain text file.
func WriteText(w io.Writer, records []model.ScanRecord) {
	for _, r := range records {
		fmt.Fprintln(w, r.CloneInstruction)
	}
}

// ParseText reads one clone command per line (delegates to the existing text parser).
// This is a convenience alias so callers can use formatter.ParseText.
func ParseText(r io.Reader) ([]model.ScanRecord, error) {
	return parseTextLines(r)
}

// parseTextLines reads non-empty lines as clone instructions.
func parseTextLines(r io.Reader) ([]model.ScanRecord, error) {
	// Reuse the cloner's text parsing via the extension dispatch.
	// Since this is the formatter package, we provide a standalone implementation.
	scanner := newLineScanner(r)
	var records []model.ScanRecord
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			records = append(records, parseCloneLineText(line))
		}
	}

	return records, scanner.Err()
}

func newLineScanner(r io.Reader) *lineScanner {
	return &lineScanner{scanner: newBufScanner(r)}
}

type lineScanner struct {
	scanner interface {
		Scan() bool
		Text() string
		Err() error
	}
}

func (ls *lineScanner) Scan() bool { return ls.scanner.Scan() }
func (ls *lineScanner) Text() string {
	return trimSpace(ls.scanner.Text())
}
func (ls *lineScanner) Err() error { return ls.scanner.Err() }

func newBufScanner(r io.Reader) *bufScanner {
	return &bufScanner{r: r}
}

// bufScanner wraps bufio.Scanner for interface compliance.
type bufScanner struct {
	r    io.Reader
	scan *scannerWrap
}

func (bs *bufScanner) Scan() bool {
	if bs.scan == nil {
		bs.scan = newScannerWrap(bs.r)
	}

	return bs.scan.Scan()
}
func (bs *bufScanner) Text() string { return bs.scan.Text() }
func (bs *bufScanner) Err() error   { return bs.scan.Err() }

// Use standard library scanner.
type scannerWrap struct {
	s *stdScanner
}

func newScannerWrap(r io.Reader) *scannerWrap {
	return &scannerWrap{s: newStdScanner(r)}
}
func (sw *scannerWrap) Scan() bool  { return sw.s.Scan() }
func (sw *scannerWrap) Text() string { return sw.s.Text() }
func (sw *scannerWrap) Err() error   { return sw.s.Err() }

func parseCloneLineText(line string) model.ScanRecord {
	parts := splitFields(line)
	rec := model.ScanRecord{CloneInstruction: line}
	if len(parts) >= 5 {
		rec.Branch = parts[3]
		rec.HTTPSUrl = parts[4]
	}
	if len(parts) >= 6 {
		rec.RelativePath = parts[5]
	}

	return rec
}

// Thin wrappers to avoid importing strings/bufio at top level for clarity.
func trimSpace(s string) string {
	return _trimSpace(s)
}

func splitFields(s string) []string {
	return _splitFields(s)
}

// These use the actual standard library — imported below.
var (
	_trimSpace   = initTrimSpace()
	_splitFields = initSplitFields()
)

func initTrimSpace() func(string) string {
	return func(s string) string {
		_ = constants.Version // ensure constants import is used
		start, end := 0, len(s)
		for start < end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\r' || s[start] == '\n') {
			start++
		}
		for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\r' || s[end-1] == '\n') {
			end--
		}

		return s[start:end]
	}
}

func initSplitFields() func(string) []string {
	return func(s string) []string {
		var fields []string
		start := -1
		for i, c := range s {
			if c == ' ' || c == '\t' {
				if start >= 0 {
					fields = append(fields, s[start:i])
					start = -1
				}
			} else if start < 0 {
				start = i
			}
		}
		if start >= 0 {
			fields = append(fields, s[start:])
		}

		return fields
	}
}
