package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	appConf "portal/config"
	globalConf "portal/global/config"
	"portal/global/gorm"
	"portal/global/log"
	"portal/httpd"
)

// 定义根命令
var rootCmd = &cobra.Command{
	Use: "kubernetes-demo-go",
	Run: func(cmd *cobra.Command, args []string) {
		conf := globalConf.GetAppConfig()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Start http server error by ", r)
				os.Exit(1)
			}
		}()
		fmt.Println("Starting init sysmgr log")
		log.Init(conf.Log)
		fmt.Println("Init sysmgr log ok")

		fmt.Println("Starting init gorm")
		gorm.InitDB()
		defer gorm.CloseDB()
		fmt.Println("Init gorm ok")

		// init gin server
		log.Info("Starting init gin server")
		httpd.StartHttpdServer(conf.Httpd)
		log.Info("Start gin server ok")
	},
}

// Execute方法触发init方法
func init() {
	// 初始化配置文件转化成对应的结构体
	cobra.OnInitialize(initConfig)
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
	dir,_ := os.Getwd()
	env := os.Getenv("ENV")
	if env == ""{
		env = "dev"
	}
	configPath := filepath.Join(dir,"config/"+env)
	// 设置读取的文件路径
	viper.AddConfigPath(configPath)
	// 设置读取的文件名
	viper.SetConfigName("config")
	// 设置文件的类型
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Read config error by %v \n",err))
	}
	var appConf appConf.AppConfig
	if err :=viper.Unmarshal(&appConf); err !=nil{
		panic(fmt.Sprintf("Unmarshal config error by %v \n",err))
	}
	globalConf.SetAppConfig(&appConf)
}