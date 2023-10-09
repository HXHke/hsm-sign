package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "HSM_SIGN"

func prefixEnvVar(prefix, suffix string) string {
	return prefix + "_" + suffix
}

var (
	HsmAPINameFlag = cli.StringFlag{
		Name:     "hsm-api-name",
		Usage:    "the api name of hsm for mt-batcher",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "HSM_API_NAME"),
	}
	HsmFeeAPINameFlag = cli.StringFlag{
		Name:     "hsm-fee-api-name",
		Usage:    "the api name of hsm for mt-batcher fee address",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "HSM_FEE_API_NAME"),
	}
	HsmAddressFlag = cli.StringFlag{
		Name:     "hsm-address",
		Usage:    "the address of hsm key for mt-batcher",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "HSM_ADDRESS"),
	}
	HsmFeeAddressFlag = cli.StringFlag{
		Name:     "hsm-fee-address",
		Usage:    "the address of hsm key for mt-batcher fee",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "HSM_FEE_ADDRESS"),
	}
	HsmCredenFlag = cli.StringFlag{
		Name:     "hsm-creden",
		Usage:    "the creden of hsm key for mt-batcher",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "HSM_CREDEN"),
	}
	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   prefixEnvVar(envVarPrefix, "L1_ETH_RPC"),
	}
	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar(envVarPrefix, "LOG_LEVEL"),
	}
	LogTerminalFlag = cli.BoolFlag{
		Name: "log-terminal",
		Usage: "If true, outputs logs in terminal format, otherwise prints " +
			"in JSON format. If SENTRY_ENABLE is set to true, this flag is " +
			"ignored and logs are printed using JSON",
		EnvVar: prefixEnvVar(envVarPrefix, "LOG_TERMINAL"),
	}
)

var requiredFlags = []cli.Flag{
	HsmCredenFlag,
	HsmAddressFlag,
	HsmAPINameFlag,
	HsmFeeAddressFlag,
	HsmFeeAPINameFlag,
	L1EthRpcFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	LogTerminalFlag,
}

var Flags = append(requiredFlags, optionalFlags...)
