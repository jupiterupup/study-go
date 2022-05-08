package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("----web service start ----")
		// TODO: 其他指令
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		_ = fmt.Errorf("RootCmd execute failed %s \n", err)
		os.Exit(-1)
	}
}
