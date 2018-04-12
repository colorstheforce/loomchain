package main

import (
	"fmt"
	"os"

	"github.com/loomnetwork/loom/examples/cmd-plugins/create-tx/cmd-plugin"
	"github.com/spf13/cobra"
)

// rootCmd is the entry point for this binary
var rootCmd = &cobra.Command{
	Use:   "ladmin",
	Short: "Loom Admin CLI (Debug Mode)",
}

func main() {
	// NOTE: The VSCode Go plugin will indicate that BuiltinCmdPluginManager is undefined
	//       unless go.buildTags contains BUILTIN_PLUGINS, can just ignore that as long as
	//       the build tag is actually set when the executable is built.
	pm := cli.BuiltinCmdPluginManager{
		RootCmd:         rootCmd,
		CmdPluginSystem: cli.NewCmdPluginSystem(),
	}
	// activate built-in cmd plugins
	createTxCmd := &cmdplugins.CreateTxCmdPlugin{}
	pm.ActivatePlugin(createTxCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
