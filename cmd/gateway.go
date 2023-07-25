package cmd

import (
	"github.com/spf13/cobra"
)

//
//
// Define the CLI data for the gateway sub-command
//
//

func newKongSubCmd() *cobra.Command {
	kongSubCmd := &cobra.Command{
		Use:   "gateway [sub-command]...",
		Short: "Sub-command to host the decK network operations",
		Long:  `Sub-command to host the decK network operations`,
	}

	return kongSubCmd
}
