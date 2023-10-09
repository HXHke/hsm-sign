package hsm_sign_service

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
	"hsm-sign/flags"
	"math/big"
)

type Config struct {
	HsmAddress    string
	HsmFeeAddress string
	HsmAPIName    string
	HsmFeeAPIName string
	HsmCreden     string
	L1EthRpc      string
	LogLevel      string
	LogTerminal   bool
	L1ChainID     *big.Int
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		HsmCreden:     ctx.GlobalString(flags.HsmCredenFlag.Name),
		HsmAddress:    ctx.GlobalString(flags.HsmAddressFlag.Name),
		HsmAPIName:    ctx.GlobalString(flags.HsmAPINameFlag.Name),
		HsmFeeAddress: ctx.GlobalString(flags.HsmFeeAddressFlag.Name),
		HsmFeeAPIName: ctx.GlobalString(flags.HsmFeeAPINameFlag.Name),
		L1EthRpc:      ctx.GlobalString(flags.L1EthRpcFlag.Name),
		LogLevel:      ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:   ctx.GlobalBool(flags.LogTerminalFlag.Name),
	}
	err := ValidateConfig(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func ValidateConfig(cfg *Config) error {
	if cfg.LogLevel == "" {
		cfg.LogLevel = "debug"
	}
	_, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return err
	}
	return nil
}
