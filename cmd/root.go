package cmd

import (
	"fmt"

	"github.com/FlorentinDUBOIS/bananas/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd is the root command line interface
var RootCmd = cobra.Command{
	Use:   "bananas",
	Short: "A slack bot who love bananas!",
	Run:   root,
}

func init() {
	cobra.OnInitialize(configure)

	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output")
	RootCmd.Flags().Int32P("port", "p", 8080, "set the port to listen")
	RootCmd.Flags().String("token", "", "set the bearer token to use")

	viper.BindPFlags(RootCmd.PersistentFlags())
	viper.BindPFlags(RootCmd.Flags())
}

func configure() {
	viper.AutomaticEnv()

	if viper.GetBool("verbose") {
		log.SetLevel(log.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		log.SetLevel(log.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}
}

func root(cmd *cobra.Command, arguments []string) {
	r := gin.Default()
	listen := fmt.Sprintf(":%d", viper.GetInt("port"))

	for _, route := range router.Routes {
		r.Handle(route.Method, route.Path, route.Handler)
	}

	r.Run(listen)
}
