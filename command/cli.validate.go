package command

import "github.com/go-test-challenge/errors"

func (*CLI) validateCommand(cmd *Command) error {
	if cmd == nil {
		return errors.CommandEmpty
	}

	if cmd.Action == "" && cmd.Object == "" {
		return errors.CommandSyntaxError
	}

	if _, ok := ActionValue[cmd.Action]; !ok {
		return errors.ErrorWithMessage(errors.CommandActionNotFound, "'"+cmd.Action+"'", "invalid")
	}

	if _, ok := ObjectValue[cmd.Object]; !ok && cmd.Object != "" {
		return errors.ErrorWithMessage(errors.CommandObjectNotFound, "'"+cmd.Object+"'", "is not an object")
	}

	return nil
}
