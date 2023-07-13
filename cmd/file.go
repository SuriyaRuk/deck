package cmd

import (
	"github.com/spf13/cobra"
)

//
//
// Define the CLI data for the file sub-command
//
//

func newFileSubCmd() *cobra.Command {
	fileCmd := &cobra.Command{
		Use:   "file [sub-command]...",
		Short: "Sub-command to host the decK file operations",
		Long:  `Sub-command to host the decK file operations`,
	}

	return fileCmd
}
