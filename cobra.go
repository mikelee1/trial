package main

import (
	"fmt"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/peer/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const (
	nodeFuncName = "node"
	shortDes     = "Operate a peer node: start|status."
	longDes      = "Operate a peer node: start|status."
)

var versionFlag bool

//usage ./cobra node start
var mainCmd = &cobra.Command{
	Use: "peer",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// check for --logging-level pflag first, which should override all other
		// log settings. if --logging-level is not set, use CORE_LOGGING_LEVEL
		// (environment variable takes priority; otherwise, the value set in
		// core.yaml)
		var loggingSpec string
		if viper.GetString("logging_level") != "" {
			loggingSpec = viper.GetString("logging_level")
		} else {
			loggingSpec = viper.GetString("logging.level")
		}
		flogging.InitFromSpec(loggingSpec)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Print(version.GetInfo())
		} else {
			cmd.HelpFunc()(cmd, args)
		}
	},
}

func main() {
	mainCmd.AddCommand(Cmd())
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}

// Cmd returns the cobra command for Node
func Cmd() *cobra.Command {
	nodeCmd.AddCommand(startCmd())
	nodeCmd.AddCommand(statusCmd())

	return nodeCmd
}

var nodeCmd = &cobra.Command{
	Use:   nodeFuncName,
	Short: fmt.Sprint(shortDes),
	Long:  fmt.Sprint(longDes),
}

var chaincodeDevMode bool
var orderingEndpoint string

var nodeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the node.",
	Long:  `Starts a node that interacts with the network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return start()
	},
}

var nodeStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Starts the node.",
	Long:  `Starts a node that interacts with the network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return status()
	},
}

func startCmd() *cobra.Command {
	// Set the flags on the node start command.
	flags := nodeStartCmd.Flags()
	flags.BoolVarP(&chaincodeDevMode, "peer-chaincodedev", "", false,
		"Whether peer in chaincode development mode")
	flags.StringVarP(&orderingEndpoint, "orderer", "o", "orderer:7050", "Ordering service endpoint")

	return nodeStartCmd
}

func statusCmd() *cobra.Command {
	// Set the flags on the node start command.
	flags := nodeStatusCmd.Flags()
	flags.BoolVarP(&chaincodeDevMode, "peer-chaincodedev", "", false,
		"Whether peer in chaincode development mode")
	flags.StringVarP(&orderingEndpoint, "orderer", "o", "orderer:7050", "Ordering service endpoint")

	return nodeStatusCmd
}

func start() error {
	fmt.Println("start")
	return nil
}

func status() error {
	fmt.Println("status")
	return nil
}
