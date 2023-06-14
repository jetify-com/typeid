package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "typeid",
		Short: "Type-safe, K-sortable, globally unique identifiers",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	command.AddCommand(NewCmd())
	command.AddCommand(EncodeCmd())
	command.AddCommand(DecodeCmd())

	return command
}

func Execute(ctx context.Context, args []string) int {
	cmd := RootCmd()
	cmd.SetArgs(args)
	err := cmd.ExecuteContext(ctx)
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		return 1
	}
	return 0
}

func Main() {
	code := Execute(context.Background(), os.Args[1:])
	os.Exit(code)
}
