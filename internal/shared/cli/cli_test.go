package cli_test

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wigata-intech/logres/internal/shared/cli"
)

func TestRun(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		build      func(rec *string) *cli.Command
		args       []string
		wantErrIs  error
		wantRec    string
		wantOutHas string
	}

	cases := []testCase{
		{
			name: "dispatches to child",
			build: func(rec *string) *cli.Command {
				child := cli.New("child", "a child")
				child.Exec = func(_ context.Context, args []string) error {
					*rec = "child:" + firstOrEmpty(args)
					return nil
				}
				return cli.New("root", "root").Sub(child)
			},
			args:    []string{"child", "arg1"},
			wantRec: "child:arg1",
		},
		{
			name: "runs leaf exec with parsed flags",
			build: func(rec *string) *cli.Command {
				leaf := cli.New("leaf", "a leaf")
				fs := flag.NewFlagSet("leaf", flag.ContinueOnError)
				name := fs.String("name", "", "a name")
				leaf.Flags = fs
				leaf.Exec = func(_ context.Context, _ []string) error {
					*rec = "name=" + *name
					return nil
				}
				return leaf
			},
			args:    []string{"-name", "logres"},
			wantRec: "name=logres",
		},
		{
			name: "help prints usage without error",
			build: func(*string) *cli.Command {
				return cli.New("root", "root summary").Sub(cli.New("child", "child summary"))
			},
			args:       []string{"help"},
			wantOutHas: "child summary",
		},
		{
			name: "unknown subcommand is a usage error",
			build: func(*string) *cli.Command {
				return cli.New("root", "root").Sub(cli.New("child", "child"))
			},
			args:       []string{"bogus"},
			wantErrIs:  cli.ErrUsage,
			wantOutHas: "unknown command",
		},
		{
			name: "flag parse failure is a usage error",
			build: func(*string) *cli.Command {
				leaf := cli.New("leaf", "a leaf")
				leaf.Flags = flag.NewFlagSet("leaf", flag.ContinueOnError)
				leaf.Exec = func(_ context.Context, _ []string) error { return nil }
				return leaf
			},
			args:      []string{"-nope"},
			wantErrIs: cli.ErrUsage,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var rec string
			cmd := tc.build(&rec)

			var out bytes.Buffer
			cmd.SetOutput(&out)

			err := cmd.Run(context.Background(), tc.args)

			if tc.wantErrIs != nil {
				assert.ErrorIs(t, err, tc.wantErrIs)
			} else {
				require.NoError(t, err)
			}
			if tc.wantRec != "" {
				assert.Equal(t, tc.wantRec, rec)
			}
			if tc.wantOutHas != "" {
				assert.Contains(t, out.String(), tc.wantOutHas)
			}
		})
	}
}

func firstOrEmpty(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return args[0]
}
