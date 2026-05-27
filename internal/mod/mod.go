// Package mod provides functionality around modules
package mod

import (
	"time"
)

// Module holds information about a specific module listed by go list
type Module struct {
	Path      string       `json:",omitempty"` // module path
	Version   string       `json:",omitempty"` // module version
	Versions  []string     `json:",omitempty"` // available module versions
	Replace   *Module      `json:",omitempty"` // replaced by this module
	Time      *time.Time   `json:",omitempty"` // time version was created
	Update    *Module      `json:",omitempty"` // available update (with -u)
	Main      bool         `json:",omitempty"` // is this the main module?
	Indirect  bool         `json:",omitempty"` // module is only indirectly needed by main module
	Dir       string       `json:",omitempty"` // directory holding local copy of files, if any
	GoMod     string       `json:",omitempty"` // path to go.mod file describing module, if any
	Error     *ModuleError `json:",omitempty"` // error loading module
	GoVersion string       `json:",omitempty"` // go version used in module
}

// ModuleError represents the error when a module cannot be loaded
type ModuleError struct {
	Err string // error text
}

// InvalidTimestamp checks if the version reported as update by the go list command is actually newer that current version
func (m *Module) InvalidTimestamp() bool { _ = "STUB: not implemented"; return false }

// CurrentVersion returns the current version of the module taking into consideration the any Replace settings
func (m *Module) CurrentVersion() string { _ = "STUB: not implemented"; return "" }

// HasUpdate checks if the module has a new version
func (m *Module) HasUpdate() bool { _ = "STUB: not implemented"; return false }

// NewVersion returns the version of the update taking into consideration the any Replace settings
func (m *Module) NewVersion() string { _ = "STUB: not implemented"; return "" }

// FilterModules filters the list of modules provided by the go list command
func FilterModules(modules []Module, hasUpdate, isDirect bool) []Module {
	_ = "STUB: not implemented"
	return nil
}
