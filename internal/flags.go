package flags

import (
	"flag"
	"log"
	"os"
)

// Operation struct contains the name of the Operation selected in the CLI and the Parameters chosen for it
type Operation struct {
	Name       string
	Parameters Parameters
}

// Parameters struct contains common parameters for different operations
type Parameters struct {
	Hostname     string
	APIkey       string
	SharedSecret string
	APIURL       string
}

// GetFlags parses flags provided at runtime and returns a pointer to the Operation struct generated
func GetFlags() Operation {

	// sub-command used to Set the IP for a hostname in the DynamoDB table
	setCommand := flag.NewFlagSet("set", flag.ExitOnError)
	hostnamePtrSetCommand := setCommand.String("hostname", "", "Hostname you'd like to update with new IP for its A record in route53")
	apiKeyPtrSetCommand := setCommand.String("api-key", "", "API-Key used to make API calls")
	sharedSecretPtrSetCommand := setCommand.String("shared-secret", "", "Shared Secret used for your hostname")
	apiURLPtrSetCommand := setCommand.String("api-url", "", "Base URL of your Route53 DDNS API")

	// sub-command used to Get the current IP for a hostname in the DynamoDB table
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	apiKeyPtrGetCommand := getCommand.String("api-key", "", "API-Key used to make API calls")
	apiURLPtrGetCommand := getCommand.String("api-url", "", "Base URL of your Route53 DDNS API")

	// os.Args length must be 2 or above to include subcommands
	// if its not then error out
	if len(os.Args) < 2 {
		log.Fatalln("set, get, or list-hosts subcommands are required")
	}

	// switch statement evaluating which sub-command to parse
	switch os.Args[1] {
	case "set":
		setCommand.Parse(os.Args[2:])
	case "get":
		getCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if setCommand.Parsed() {
		if *hostnamePtrSetCommand == "" {
			setCommand.PrintDefaults()
			os.Exit(1)
		}

		if *apiKeyPtrSetCommand == "" {
			setCommand.PrintDefaults()
			os.Exit(1)
		}

		if *sharedSecretPtrSetCommand == "" {
			setCommand.PrintDefaults()
			os.Exit(1)
		}

		if *apiURLPtrSetCommand == "" {
			setCommand.PrintDefaults()
			os.Exit(1)
		}

		return Operation{
			Name: os.Args[1],
			Parameters: Parameters{
				Hostname:     *hostnamePtrSetCommand,
				APIkey:       *apiKeyPtrSetCommand,
				SharedSecret: *sharedSecretPtrSetCommand,
				APIURL:       *apiURLPtrSetCommand,
			},
		}
	}

	if getCommand.Parsed() {
		if *apiKeyPtrGetCommand == "" {
			getCommand.PrintDefaults()
			os.Exit(1)
		}

		if *apiURLPtrGetCommand == "" {
			setCommand.PrintDefaults()
			os.Exit(1)
		}

		return Operation{
			Name: os.Args[1],
			Parameters: Parameters{
				Hostname:     "",
				APIkey:       *apiKeyPtrGetCommand,
				SharedSecret: "",
				APIURL:       *apiURLPtrGetCommand,
			},
		}
	}

	return Operation{
		Name: "",
		Parameters: Parameters{
			Hostname:     "",
			APIkey:       "",
			SharedSecret: "",
			APIURL:       "",
		},
	}
}
