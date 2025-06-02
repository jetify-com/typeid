package cli

import (
	"strings"

	"github.com/spf13/cobra"
	"go.jetify.com/typeid"
)

func NewCmd() *cobra.Command {
	command := &cobra.Command{
		Use:           "new [<type_prefix>]",
		Args:          cobra.MaximumNArgs(1),
		Short:         "Generate a new TypeID using the given type prefix",
		RunE:          newCmd,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return command
}

func newCmd(cmd *cobra.Command, args []string) error {
	prefix := ""
	if len(args) > 0 {
		prefix = strings.ToLower(args[0])
	}
	tid, err := typeid.Generate(prefix)
	if err != nil {
		return err
	}
	cmd.Println(tid)
	return nil
}
