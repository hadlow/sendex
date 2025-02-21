package output

import (
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/tidwall/pretty"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"

// clear styling on files as it shows as raw chars
var FileStyle = &pretty.Style{
	Key:      [2]string{"", ""},
	String:   [2]string{"", ""},
	Number:   [2]string{"", ""},
	True:     [2]string{"", ""},
	False:    [2]string{"", ""},
	Null:     [2]string{"", ""},
	Escape:   [2]string{"", ""},
	Brackets: [2]string{"", ""},
}

type Buffer struct {
	buffer strings.Builder
	raw    bool // true if no styling needed
}

func (b *Buffer) Write(text string, style string) {
	if b.raw {
		b.buffer.WriteString(text + "\n")
	} else {
		b.buffer.WriteString(style + text + Reset + "\n")
	}
}

func (b *Buffer) WriteInfo(text string) {
	b.Write(text, Blue)
}

func (b *Buffer) WriteSuccess(text string) {
	b.Write(text, Green)
}

func (b *Buffer) WriteWarning(text string) {
	b.Write(text, Yellow)
}

func (b *Buffer) WriteError(text string) {
	b.Write(text, Red)
}

func (b *Buffer) HeaderItem(header string, value []string) {
	if b.raw {
		b.buffer.WriteString(header + ": " + strings.Join(value, ", ") + "\n")
	} else {
		b.buffer.WriteString(Cyan + header + Reset + ": " + strings.Join(value, ", ") + "\n")
	}
}

func (b *Buffer) Status(response *http.Response) {
	if response.StatusCode < 200 {
		b.WriteWarning(response.Status)
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		b.WriteWarning(response.Status)
	}

	if response.StatusCode >= 300 && response.StatusCode < 400 {
		b.WriteWarning(response.Status)
	}

	if response.StatusCode >= 400 && response.StatusCode < 500 {
		b.WriteError(response.Status)
	}

	if response.StatusCode >= 500 {
		b.WriteError(response.Status)
	}
}

func (b *Buffer) Head(response *http.Response, AllowHeaders []string) error {
	for header, value := range response.Header {
		if len(AllowHeaders) == 0 || slices.Contains(AllowHeaders, header) {
			b.HeaderItem(header, value)
		}
	}

	return nil
}

func (b *Buffer) Body(response *http.Response) error {
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	var prettyJSON []byte

	if b.raw {
		prettyJSON = pretty.Color([]byte(body), FileStyle)
	} else {
		prettyJSON = pretty.Color([]byte(body), nil)
	}

	b.buffer.WriteString(string(prettyJSON[:]) + "\n")

	return nil
}
