package service

import (
	kms "cloud.google.com/go/kms/apiv1"
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"google.golang.org/api/option"
	"math/big"
	"sync"
	"time"
)

var DefaultTimeout = 15 * time.Second

type DriverConfig struct {
	L1Client      *ethclient.Client
	HsmAddress    string
	HsmFeeAddress string
	HsmAPIName    string
	HsmFeeAPIName string
	HsmCreden     string
	L1ChainID     *big.Int
}

type Driver struct {
	Ctx    context.Context
	Cfg    *DriverConfig
	cancel func()
	wg     sync.WaitGroup
}

func NewDriver(ctx context.Context, cfg *DriverConfig) (*Driver, error) {
	_, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	return &Driver{
		Cfg:    cfg,
		Ctx:    ctx,
		cancel: cancel,
	}, nil
}

func (d *Driver) Start() error {
	d.wg.Add(1)
	go d.Work()
	return nil
}

func (d *Driver) Stop() {
	d.cancel()
	d.wg.Wait()
}

func (d *Driver) Work() {
	var err error
	var opts *bind.TransactOpts
	opts, err = NewHSMTransactOpts(d.Ctx, d.Cfg.HsmFeeAPIName,
		d.Cfg.HsmFeeAddress, d.Cfg.L1ChainID, d.Cfg.HsmCreden)
	if err != nil {
		return
	}

	nonce64, err := d.Cfg.L1Client.NonceAt(
		d.Ctx, common.HexToAddress(d.Cfg.HsmAddress), nil,
	)

	opts.Context = d.Ctx
	opts.Nonce = new(big.Int).SetUint64(nonce64)
	opts.NoSend = true

	//todo add bind contract to send tx

}

func NewHSMTransactOpts(ctx context.Context, hsmAPIName string, hsmAddress string, chainID *big.Int, hsmCreden string) (*bind.TransactOpts, error) {
	proBytes, err := hex.DecodeString(hsmCreden)
	apikey := option.WithCredentialsJSON(proBytes)
	client, err := kms.NewKeyManagementClient(ctx, apikey)
	if err != nil {
		return nil, err
	}
	mk := &bsscore.ManagedKey{
		KeyName:      hsmAPIName,
		EthereumAddr: common.HexToAddress(hsmAddress),
		Gclient:      client,
	}
	opts, err := mk.NewEthereumTransactorrWithChainID(ctx, chainID)
	return opts, nil
}
