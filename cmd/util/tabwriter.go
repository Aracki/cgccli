// Package util provides utility tools for CLI.
package util

import (
	"io"
	"os"
	"text/tabwriter"
)

/*
	minwidth is the minimum column width,

	tabwidth is the number of spaces you want tab to be if you specify tab as the padchar,

	padding is the number of cells you want to add to the contents of the cell before it’s width is computed,
	with this you can prevent data that might take up the entire cell from butting up against the data in an adjacent cell,

	padchar is the ASCII character to fill out the cell with, this can be any byte character, whether that’s a space, ' ', a symbol, '*' or whatever.
	If you specify a tab character here, '\t', then the tabwidth will be used to determine it’s width.

	Lastly there are some flags we can pass in. These flags are represented by uint numbers,
	so if we don’t want any flags we can just pass in 0. There’s six flags, dealing with HTML, stripping escape characters,
 	overriding the default left alignment to align right (no centre alignment that I can see),discarding empty columns,
	overriding padchar to always indent columns with tabs and a handy Debug flag that sticks in a vertical bar, |, to distinguish where your columns are.
 */
const (
	tabwriterMinWidth = 0
	tabwriterTabWidth = 0
	tabwriterPadding  = 2
	tabwriterPadChar  = ' '
	tabwriterFlags    = 1
)

// Change where to spit all outputs of CLI
var whereToPrint io.Writer = os.Stdout

// NewTabWriter creates a new Writer with a given tab writer properties.
func NewTabWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(whereToPrint, tabwriterMinWidth, tabwriterTabWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
}
