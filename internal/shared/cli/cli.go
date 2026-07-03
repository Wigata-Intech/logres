// Package cli is a tiny, dependency-free command router in the spirit of cobra:
// a tree of named commands, each with an optional flag set and an executor. It
// keeps Logres a single binary without a third-party CLI framework.
package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func defaultOutput() io.Writer { return os.Stderr }

// ErrUsage signals that the command was invoked incorrectly; callers should
// treat it as a usage error (exit code 2) rather than a runtime failure.
var ErrUsage = errors.New("cli: usage error")

type ExecFunc func(ctx context.Context, args []string) error

// Command is a node in the command tree. A node with children dispatches to
// them; a node with an Exec runs it. A node may have both: children take
// precedence when the next argument names one.
type Command struct {
	Name    string
	Summary string
	Flags   *flag.FlagSet
	Exec    ExecFunc

	out      io.Writer
	subs     map[string]*Command
	subOrder []string
}

func New(name, summary string) *Command {
	return &Command{
		Name:    name,
		Summary: summary,
		subs:    make(map[string]*Command),
	}
}

func (c *Command) Sub(child *Command) *Command {
	if _, exists := c.subs[child.Name]; !exists {
		c.subOrder = append(c.subOrder, child.Name)
	}
	c.subs[child.Name] = child
	return c
}

// SetOutput directs help/usage text; it propagates to children. Defaults to
// os.Stderr at the point Run is called if never set.
func (c *Command) SetOutput(w io.Writer) {
	c.out = w
	for _, s := range c.subs {
		s.SetOutput(w)
	}
}

// Run parses flags and dispatches. args are the arguments after the command's
// own name (i.e. os.Args[1:] for the root).
func (c *Command) Run(ctx context.Context, args []string) error {
	if c.out == nil {
		c.SetOutput(defaultOutput())
	}

	// Route to a child before consuming flags so children own their flags.
	if len(args) > 0 {
		if child, ok := c.subs[args[0]]; ok {
			return child.Run(ctx, args[1:])
		}
		if isHelp(args[0]) {
			c.usage()
			return nil
		}
	}

	if c.Flags != nil {
		c.Flags.SetOutput(c.out)
		c.Flags.Usage = c.usage
		if err := c.Flags.Parse(args); err != nil {
			return ErrUsage
		}
		args = c.Flags.Args()
	}

	if c.Exec != nil {
		return c.Exec(ctx, args)
	}

	// A group node reached with no matching child.
	if len(args) > 0 {
		fmt.Fprintf(c.out, "%s: unknown command %q\n\n", c.Name, args[0])
	}
	c.usage()
	return ErrUsage
}

func (c *Command) usage() {
	var b strings.Builder
	fmt.Fprintf(&b, "%s — %s\n\n", c.Name, c.Summary)

	if len(c.subOrder) > 0 {
		fmt.Fprintf(&b, "Usage:\n  %s <command> [flags]\n\nCommands:\n", c.Name)
		names := append([]string(nil), c.subOrder...)
		sort.Strings(names)
		width := 0
		for _, n := range names {
			if len(n) > width {
				width = len(n)
			}
		}
		for _, n := range names {
			fmt.Fprintf(&b, "  %-*s  %s\n", width, n, c.subs[n].Summary)
		}
	} else {
		fmt.Fprintf(&b, "Usage:\n  %s [flags]\n", c.Name)
	}

	fmt.Fprint(c.out, b.String())
	if c.Flags != nil {
		fmt.Fprintln(c.out, "\nFlags:")
		c.Flags.PrintDefaults()
	}
}

func isHelp(arg string) bool {
	switch arg {
	case "help", "-h", "--help":
		return true
	default:
		return false
	}
}
