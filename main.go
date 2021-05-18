package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mch1307/vaultlib"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type FlagConfiguration struct {
	EnvFile    string
	SecretPath string
}

var flagConfig FlagConfiguration

func init() {
	pflag.String("env-file", ".env", "Write in a file of environment variables")
	pflag.String("secret-path", "", "Vault secret path")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	flagConfig = FlagConfiguration{
		EnvFile:    viper.GetString("env-file"),
		SecretPath: viper.GetString("secret-path"),
	}
	if flagConfig.SecretPath == "" {
		log.Fatal("please set --secret-path flag")
	}
	if os.Getenv("VAULT_ROLEID") == "" {
		os.Setenv("VAULT_ROLEID", os.Getenv("VAULT_ROLE_ID"))
	}
	if os.Getenv("VAULT_SECRETID") == "" {
		os.Setenv("VAULT_SECRETID", os.Getenv("VAULT_SECRET_ID"))
	}
}

func main() {
	config := vaultlib.NewConfig()
	client, errConfig := vaultlib.NewClient(config)
	if errConfig != nil {
		log.Fatal(errConfig)
	}
	kv, errClient := client.GetSecret(flagConfig.SecretPath)
	if errClient != nil {
		log.Fatal(errClient)
	}
	envs := fmt.Sprintf("# Vault secret: %s\n", flagConfig.SecretPath)
	for k, v := range kv.KV {
		envs += fmt.Sprintf("%v=%q\n", k, v)
	}
	errWriteFile := ioutil.WriteFile(flagConfig.EnvFile, []byte(envs), 0644)
	if errWriteFile != nil {
		log.Fatal(errWriteFile)
	}
	log.Printf("Write vault secret: %s to %s\n", flagConfig.SecretPath, flagConfig.EnvFile)
}
