package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/leishi1313/Shifter/route"
	"github.com/leishi1313/Shifter/scheduler"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Shifter",
	Short: "一个用来刷流/追剧的RSS管理软件",
	Long:  `一个用来刷流/追剧的RSS管理软件`,
	Run: func(cmd *cobra.Command, args []string) {
		scheduler.Run()

		addr := viper.GetString("web.bind")
		port := viper.GetString("web.port")

		router := route.Init()
		router.Logger.Fatal(router.Start(fmt.Sprintf("%s:%s", addr, port)))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	flags := rootCmd.Flags()
	persistent := rootCmd.PersistentFlags()

	persistent.StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")

	flags.BoolP("toggle", "t", false, "Help message for toggle")

	flags.StringP("web.bind", "b", "127.0.0.1", "address to listen on")
	flags.StringP("web.port", "p", "8000", "port to listen on")

	viper.BindPFlags(flags)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find current directory.
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("Shifter")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
