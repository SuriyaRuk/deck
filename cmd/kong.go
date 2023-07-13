/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

//
//
// Define the CLI data for the kong sub-command
//
//

func newKongSubCmd() *cobra.Command {
	kongSubCmd := &cobra.Command{
		Use:   "kong [sub-command]...",
		Short: "Sub-command to host the decK network operations",
		Long:  `Sub-command to host the decK network operations`,
	}

	return kongSubCmd
}
