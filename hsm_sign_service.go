package hsm_sign_service

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
	"hsm-sign/client"
	service "hsm-sign/service"
	"math/big"
	"os"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cfg, err := NewConfig(ctx)
		if err != nil {
			return err
		}
		log.Info("Initializing hsm sign service")
		hSService, err := NewHsmSignService(cfg)
		if err != nil {
			log.Error("Unable to create hsm sign service", "error", err)
			return err
		}
		log.Info("Starting hsm sign service")
		if err := hSService.Start(); err != nil {
			return err
		}
		defer hSService.Stop()

		log.Info("hsm sign service started")
		<-(chan struct{})(nil)
		return nil
	}
}

type HsmSignService struct {
	ctx           context.Context
	cfg           Config
	driverService *service.Driver
}

func NewHsmSignService(cfg Config) (*HsmSignService, error) {
	ctx := context.Background()
	var logHandler log.Handler
	if cfg.LogTerminal {
		logHandler = log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	} else {
		logHandler = log.StreamHandler(os.Stdout, log.JSONFormat())
	}
	logLevel, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))

	l1Client, err := client.L1EthClientWithTimeout(ctx, cfg.L1EthRpc, false)

	if err != nil {
		return nil, err
	}
	dConfig := &service.DriverConfig{
		HsmAddress:    cfg.HsmAddress,
		HsmAPIName:    cfg.HsmAPIName,
		HsmCreden:     cfg.HsmCreden,
		HsmFeeAPIName: cfg.HsmFeeAPIName,
		HsmFeeAddress: cfg.HsmFeeAddress,
		L1Client:      l1Client,
		L1ChainID:     new(big.Int).SetInt64(1),
	}
	dService, err := service.NewDriver(ctx, dConfig)
	if err != nil {
		log.Error("create hsm sign service fail", "err", err)
		return nil, err
	}

	return &HsmSignService{
		ctx:           ctx,
		cfg:           cfg,
		driverService: dService,
	}, nil
}

func (hss *HsmSignService) Start() error {
	err := hss.driverService.Start()
	if err != nil {
		return err
	}

	return nil
}

func (hss *HsmSignService) Stop() {
	hss.driverService.Stop()
}
