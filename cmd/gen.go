/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ushinohama966/pass-gen/internal"
)

const (
	characterLengthFlag              = "length"
	excludeSpecialCharFlag           = "exclude-special-character"
	defaultPassLength         uint16 = 12
	defaultExcludeSpecialChar        = false
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate password",
	Long:  `generate password`,
	Run: func(cmd *cobra.Command, _ []string) {
		lengthString := cmd.Flag(characterLengthFlag).Value.String()
		n, _ := strconv.Atoi(lengthString)

		exclude, _ := strconv.ParseBool(cmd.Flag(excludeSpecialCharFlag).Value.String())

		internal.Gen(uint16(n), exclude)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().Uint16("length", defaultPassLength, "specify password length")
	genCmd.PersistentFlags().Uint16P(
		characterLengthFlag,
		"l",
		defaultPassLength,
		"specify password length",
	)
	genCmd.PersistentFlags().BoolP(
		excludeSpecialCharFlag,
		"e",
		defaultExcludeSpecialChar,
		"exclude special characters",
	)
	genCmd.PersistentFlags().Lookup(excludeSpecialCharFlag).NoOptDefVal = "true"

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
