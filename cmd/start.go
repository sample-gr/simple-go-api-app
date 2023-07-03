package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cloudbees-test/simple-go-api-app/config"
	"github.com/cloudbees-test/simple-go-api-app/server"
)

func init() {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start server",
		Long:  `Start the api-server`,
		Run: func(cmd *cobra.Command, args []string) {
			var config config.Configuration
			viper.Unmarshal(&config)
			server.Start(&config)
		},
	}

	// port flag
	cmd.Flags().IntP("port", "p", viper.GetInt("port"), "Port to run the server on")
	viper.BindPFlag("port", cmd.Flags().Lookup("port"))

	rootCmd.AddCommand(cmd)
}
