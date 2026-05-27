// Package runner is responsible for running the command and rendering the output
package runner

import (
	"io"
	"os"

	"github.com/psampaz/go-mod-outdated/internal/mod"
)

// OsExit is use here in order to simplify testing
var OsExit = os.Exit

// OutputStyle specifies the supported table rendering formats
type OutputStyle string

const (
	// StyleDefault represents the default output style
	StyleDefault OutputStyle = "default"
	// StyleMarkdown represents the markdown formatted output style
	StyleMarkdown OutputStyle = "markdown"
)

// Run converts the the json output of go list -u -m -json all to table format
func Run(in io.Reader, out io.Writer, update, direct, exitWithNonZero bool, style OutputStyle) error {
	_ = "STUB: not implemented"
	return nil
}

func hasOutdated(filteredModules []mod.Module) bool { _ = "STUB: not implemented"; return false }

func renderTable(writer io.Writer, modules []mod.Module, style OutputStyle) {
	_ = "STUB: not implemented"
	return
}

// Render table as markdown
