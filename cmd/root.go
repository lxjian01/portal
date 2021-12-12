package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	appConf "portal/global/config"
	globalConf "portal/global/config"
	"portal/global/consul"
	"portal/global/log"
	"portal/global/myorm"
	"portal/web"
)

var cfgFile string

// 定义根命令
var rootCmd = &cobra.Command{
	Use: "portal",
	Run: func(cmd *cobra.Command, args []string) {
		conf := globalConf.GetAppConfig()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Start http server error by ", r)
				os.Exit(1)
			}
		}()
		fmt.Println("Starting init log")
		log.Init(conf.Log)
		fmt.Println("Init log ok")

		fmt.Println("Starting init gorm")
		myorm.InitDB()
		defer myorm.CloseDB()
		fmt.Println("Init gorm ok")

		fmt.Println("Starting init consul client")
		consul.InitConsul()
		fmt.Println("Init consul client ok")

		// init gin server
		log.Info("Starting gin server")
		web.StartServer(conf.Httpd)
		log.Info("Start gin server ok")
	},
}

// Execute方法触发init方法
func init() {
	// 初始化配置文件转化成对应的结构体
	cobra.OnInitialize(initConfig)
	// config file like --config=/opt/code/portal/config/portal.yaml
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "default is $HOME/portal.yaml")
}

// 启动调用的入口方法
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Execute error by ", err)
		os.Exit(1)
	}
}

//通过viper初始化配置文件到结构体
func initConfig() {
	if cfgFile != "" {
		// use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Printf("find home dir error by %s \n", err.Error())
			os.Exit(1)
		}
		// 设置读取的文件路径
		viper.AddConfigPath(home)
		// 设置读取的文件名
		viper.SetConfigName("config")
		// 设置文件的类型
		viper.SetConfigType("yaml")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Read config error by %v \n", err)
		os.Exit(1)
	}
	var appConf appConf.AppConfig
	if err :=viper.Unmarshal(&appConf); err !=nil{
		fmt.Printf("Unmarshal config error by %v \n", err)
		os.Exit(1)
	}
	globalConf.SetAppConfig(&appConf)
}