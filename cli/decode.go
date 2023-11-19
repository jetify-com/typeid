package cli

import (
	"github.com/spf13/cobra"
	"go.jetpack.io/typeid"
)

func DecodeCmd() *cobra.Command {
	command := &cobra.Command{
		Use:           "decode <type_id>",
		Args:          cobra.ExactArgs(1),
		Short:         "Decode the given TypeID into a UUID",
		RunE:          decodeCmd,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return command
}

func decodeCmd(cmd *cobra.Command, args []string) error {
	tid, err := typeid.FromString(args[0])
	if err != nil {
		return err
	}
	cmd.Printf("type: %s\n", tid.Prefix())
	cmd.Printf("uuid: %s\n", tid.UUID())
	return nil
}
