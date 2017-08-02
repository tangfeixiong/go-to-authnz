package hydraexample

import "github.com/spf13/cobra"

func ExampleAppCmd() *cobra.Command {
	return &cobra.Command{
		Use: "example-app",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
}
