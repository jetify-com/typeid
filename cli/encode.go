package cli

import (
	"github.com/spf13/cobra"
	"go.jetify.com/typeid/v2"
)

func EncodeCmd() *cobra.Command {
	command := &cobra.Command{
		Use:           "encode [<type_prefix>] <uuid>",
		Args:          cobra.RangeArgs(1, 2),
		Short:         "Encode the given UUID into a TypeID using the given type prefix",
		RunE:          encodeCmd,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return command
}

func encodeCmd(cmd *cobra.Command, args []string) error {
	prefix := ""
	uuid := ""
	if len(args) == 1 {
		uuid = args[0]
	} else {
		prefix = args[0]
		uuid = args[1]
	}
	tid, err := typeid.FromUUID(prefix, uuid)
	if err != nil {
		return err
	}
	cmd.Println(tid)
	return nil
}
