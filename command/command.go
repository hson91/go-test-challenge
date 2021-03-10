package command

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/go-test-challenge/errors"
	"github.com/go-test-challenge/models"
	"github.com/go-test-challenge/serializers"
	"github.com/go-test-challenge/services"
	"github.com/go-test-challenge/utilities"
)

// define object
const (
	ObjectUser         string = "user"
	ObjectOrganization string = "organization"
	ObjectTicket       string = "ticket"
)

// ObjectValue : map[string]string
var ObjectValue = map[string]string{
	"user":         ObjectUser,
	"organization": ObjectOrganization,
	"ticket":       ObjectTicket,
}

// define Action
const (
	ActionFind        string = "find"
	ActionPrintStruct string = "struct"
	ActionHelp        string = "help"
	ActionReload      string = "reload"
	ActionExit        string = "exit"
)

// ActionValue : map[string]string
var ActionValue = map[string]string{
	"find":   ActionFind,
	"struct": ActionPrintStruct,
	"help":   ActionHelp,
	"reload": ActionReload,
	"exit":   ActionExit,
	"quit":   ActionExit,
}

// Command : struct
type Command struct {
	Object  string            `json:"object"`
	Action  string            `json:"action"`
	Filters map[string]string `json:"filters"`
}

// CLI : struct
type CLI struct {
	userSrv         *services.UserSrv
	organizationSrv *services.OrganizationSrv
	ticketSrv       *services.TicketSrv
}

// NewCLI : create new instance CLI
func NewCLI(
	userSrv *services.UserSrv,
	organizationSrv *services.OrganizationSrv,
	ticketSrv *services.TicketSrv) *CLI {
	return &CLI{
		userSrv:         userSrv,
		organizationSrv: organizationSrv,
		ticketSrv:       ticketSrv,
	}
}

// ParseCommand : parse string to Command
func (cli *CLI) ParseCommand(commandStr string) *Command {
	if commandStr == "" {
		return nil
	}

	var cmd = &Command{}

	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommands := strings.SplitN(commandStr, " ", 3)

	if len(arrCommands) == 1 {
		cmd.Action = arrCommands[0]
		return cmd
	}

	cmd.Object = arrCommands[0]

	if len(arrCommands) > 1 {
		cmd.Action = arrCommands[1]
	}

	if len(arrCommands) > 2 {
		cmd.Filters = cli.ParseJSONFromString(arrCommands[2])

	}
	return cmd
}

// ParseJSONFromString : parse string to map[string]string
func (*CLI) ParseJSONFromString(s string) map[string]string {
	results := map[string]string{}

	lastQuote := rune(0)
	f := func(c rune) bool {
		switch {
		case c == lastQuote:
			lastQuote = rune(0)
			return false
		case lastQuote != rune(0):
			return false
		case unicode.In(c, unicode.Quotation_Mark):
			lastQuote = c
			return false
		default:
			return unicode.IsSpace(c)
		}
	}
	args := strings.FieldsFunc(s, f)

	re, err := regexp.Compile(`["']`)
	if err != nil {
		fmt.Println(err)
		return results
	}

	for _, item := range args {
		argParse := strings.Split(item, "=")

		key := strings.Trim(argParse[0], "-")

		key = strings.Trim(key, " ")

		key = strings.ToLower(key)

		value := strings.Trim(argParse[len(argParse)-1], " ")
		results[key] = re.ReplaceAllString(value, "")
	}

	return results
}

// Help :
func (*CLI) Help() {
	fmt.Println(
		"Syntax : <object> <action> <options> \n",
		"object: organization, user, ticket \n",
		"action: struct, find \n",
		"options: attribute=value \n",
		"Example: \n",
		"\t view struct organization ->: organization struct \n",
		"\t Find user by id ->: user find _id=1 \n",
		"\t Find ticket by type -> ticket find type='incident' \n",
		"Type `reload` to reload app \n ",
	)
}

// Run :
func (cli *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type 'help' to see systax or 'quit' to exit ")
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		cmd := cli.ParseCommand(cmdString)
		cli.RunCmd(cmd)
	}
}

// RunCmd :
func (cli *CLI) RunCmd(cmd *Command) {
	if err := cli.validateCommand((cmd)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	object := ObjectValue[cmd.Object]
	action := ActionValue[cmd.Action]

	switch action {
	case ActionReload:
		cli.Run()
		break
	case ActionHelp:
		cli.Help()
	case ActionFind:
		cli.Find(object, cmd.Filters)
	case ActionPrintStruct:
		cli.PrintStruct(object)
	case ActionExit:
		os.Exit(1)
	default:
		fmt.Println(" => Command invalid. type 'help' show all options")
	}
}

// Find :
func (cli *CLI) Find(object string, filters map[string]string) {
	switch object {
	case ObjectOrganization:
		var filtersReq serializers.OrganizationReq
		if err := utilities.Switch2Struct(&filters, &filtersReq); err != nil {
			fmt.Println(err)
			break
		}
		organizations, err := cli.organizationSrv.GetAllOrganization(&filtersReq)
		if err != nil {
			fmt.Println(err)
			break
		}

		msg := cli.organizationSrv.PrintData(organizations)
		fmt.Println(msg)
		break
	case ObjectUser:
		var filtersReq serializers.UserReq
		if err := utilities.Switch2Struct(&filters, &filtersReq); err != nil {
			fmt.Println(err)
			break
		}
		users, err := cli.userSrv.GetAllUser(&filtersReq)
		if err != nil {
			fmt.Println(err)
			break
		}

		msg := cli.userSrv.PrintData(users)
		fmt.Println(msg)
		break
	case ObjectTicket:
		var filtersReq serializers.TicketReq
		if err := utilities.Switch2Struct(&filters, &filtersReq); err != nil {
			fmt.Println(err)
			break
		}
		tickets, err := cli.ticketSrv.GetAllTicket(&filtersReq)
		if err != nil {
			fmt.Println(err)
			break
		}

		msg := cli.ticketSrv.PrintData(tickets)
		fmt.Println(msg)
		break
	default:
		fmt.Println(errors.CommandObjectNotFound)
	}
}

// PrintStruct :
func (cli *CLI) PrintStruct(object string) {
	switch object {
	case ObjectOrganization:
		msg, err := utilities.PrintKeysOfObject(&models.Organization{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
		break
	case ObjectUser:
		msg, err := utilities.PrintKeysOfObject(&models.User{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
		break
	case ObjectTicket:
		msg, err := utilities.PrintKeysOfObject(&models.Ticket{})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
		break
	default:
		fmt.Println(errors.CommandObjectNotFound)
	}
}
