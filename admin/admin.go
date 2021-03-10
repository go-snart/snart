// Package admin contains a plugin that provides admin-only commands for Snart Bots.
package admin

// Admin is a plugin that provides admin-only commands for Snart Bots.
type Admin struct {
	Errs chan error
}

// String returns the Admin's string representation.
func (a *Admin) String() string {
	return "admin plug"
}
