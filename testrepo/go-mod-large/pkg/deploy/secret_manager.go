package deploy

import (
	"os"
	"path/filepath"

	"github.com/flant/logboek"

	"github.com/flant/werf/pkg/deploy/secret"
	"github.com/flant/werf/pkg/deploy/werf_chart"
)

func GetSafeSecretManager(projectDir string, secretValues []string, ignoreSecretKey bool) (secret.Manager, error) {
	isSecretsExists := false
	if _, err := os.Stat(filepath.Join(projectDir, werf_chart.ProjectSecretDir)); !os.IsNotExist(err) {
		isSecretsExists = true
	}
	if _, err := os.Stat(filepath.Join(projectDir, werf_chart.ProjectDefaultSecretValuesFile)); !os.IsNotExist(err) {
		isSecretsExists = true
	}
	if len(secretValues) > 0 {
		isSecretsExists = true
	}

	if isSecretsExists {
		if ignoreSecretKey {
			logboek.LogInfoLn("Secrets decryption disabled")
			return secret.NewSafeManager()
		}

		key, err := secret.GetSecretKey(projectDir)
		if err != nil {
			return nil, err
		}

		return secret.NewManager(key, secret.NewManagerOptions{})
	} else {
		return secret.NewSafeManager()
	}
}
