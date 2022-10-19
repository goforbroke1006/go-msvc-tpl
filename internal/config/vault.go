package config

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func LoadFromVault() error {
	var (
		vaultAddr  = viper.GetString("vault.address")
		vaultToken = viper.GetString("vault.token")
		vaultPath  = viper.GetString("vault.path")
	)

	cfg := &api.Config{
		Address: vaultAddr,
		HttpClient: &http.Client{
			Timeout: 5 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
	client, err := api.NewClient(cfg)
	if err != nil {
		return errors.Wrap(err, "can't create http client")
	}
	client.SetToken(vaultToken)

	data, err := client.Logical().Read(vaultPath)
	if err != nil {
		return errors.Wrap(err, "can't read from vault")
	}

	for vaultKey, vaultValue := range data.Data {
		_ = os.Setenv(vaultKey, fmt.Sprintf("%v", vaultValue))
	}

	return nil
}
