package cmd

import (
	"flag"
	"math/rand"
	"time"

	"github.com/cesanta/docker_auth/auth_server/server"
	"github.com/cesanta/glog"
	"github.com/facebookgo/httpdown"
	hydracmd "github.com/ory/hydra/cmd"
	"github.com/spf13/cobra"

	dexexample "github.com/tangfeixiong/go-to-authnz/cmd/example-app"
	hydraexample "github.com/tangfeixiong/go-to-authnz/cmd/hydra-consent-app-go"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Use: "go-to-docker",
	}

	hydracmd.RootCmd.AddCommand(hydraexample.ExampleAppCmd()) //hydra
	rootCmd.AddCommand(hydracmd.RootCmd)                      //hydra

	dexcmd := commandRoot()
	dexcmd.AddCommand(dexexample.ExampleAppCmd()) //dex
	rootCmd.AddCommand(dexcmd)                    //dex

	rootCmd.AddCommand(commandMain()) //cesanta docker_auth auth_server main

	return rootCmd.Execute()
}

func commandMain() *cobra.Command {
	mainCmd := &cobra.Command{
		Use: "cesantadas <config file>",
		Run: func(cmd *cobra.Command, args []string) {
			flag.Parse()
			rand.Seed(time.Now().UnixNano())
			glog.CopyStandardLogTo("INFO")

			glog.Infof("docker_auth %s build %s", Version, BuildId)

			cf := flag.Arg(0)
			if cf == "" {
				glog.Exitf("Config file not specified")
			}
			c, err := server.LoadConfig(cf)
			if err != nil {
				glog.Exitf("Failed to load config: %s", err)
			}
			rs := RestartableServer{
				configFile: cf,
				hd:         &httpdown.HTTP{},
			}
			rs.Serve(c)
		},
	}
	mainCmd.Flags().AddGoFlagSet(flag.CommandLine)
	return mainCmd
}
