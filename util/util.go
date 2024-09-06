package util

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func GetEnv(cmd *cobra.Command, flagName, envName string) string {
	value, _ := cmd.Flags().GetString(flagName)
	if value != "" {
		err := os.Setenv(envName, value)
		if err != nil {
			return value
		}
	}
	return os.Getenv(envName)
}
func GetIntEnv(envName string) int {
	val := os.Getenv(envName)
	if val == "" {
		Fatal("Error %s  is required \n", envName)
	}
	ret, err := strconv.Atoi(val)
	if err != nil {
		Fatal("Error %s  is required \n", envName)

	}
	return ret
}

// CheckEnvVars checks if all the specified environment variables are set
func CheckEnvVars(vars []string) error {
	missingVars := []string{}

	for _, v := range vars {
		if os.Getenv(v) == "" {
			missingVars = append(missingVars, v)
		}
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("missing environment variables: %v", missingVars)
	}

	return nil
}

func LoadEnv(file string) {
	if file != "" {
		godotenv.Load(file)
		return
	} else {
		godotenv.Load()
	}
}
