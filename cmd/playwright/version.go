package main

import (
	"github.com/asalbers/porter-playwright/pkg/playwright"
	"github.com/asalbers/porter-playwright/pkg/version"
	"github.com/spf13/cobra"
)

func buildVersionCommand(m *playwright.Mixin) *cobra.Command {
	opts := version.Options{}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the mixin verison",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.PrintVersion(opts)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&opts.RawFormat, "output", "o", string(version.DefaultVersionFormat),
		"Specify an output format.  Allowed values: json, plaintext")

	return cmd
}
