package main

import (
	"ecos/edge-node/alaya"
	moonConfig "ecos/edge-node/config"
	"ecos/edge-node/moon"
	"ecos/edge-node/node"
	"ecos/messenger"
	"ecos/utils/config"
	"ecos/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"strconv"
	"time"
)

var Moon *moon.Moon

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", getGroupInfo)
	return router
}

func getGroupInfo(c *gin.Context) {
	if Moon == nil {
		c.String(http.StatusBadGateway, "Edge Node Not Ready")
	}
	info := Moon.InfoStorage.ListAllNodeInfo()
	c.JSONP(http.StatusOK, info)
}

func action(c *cli.Context) error {
	// read config
	confPath := c.Path("config")
	conf := moonConfig.DefaultConfig
	conf.Moon.SunAddr = c.String("sun")
	config.Register(conf, confPath)
	config.ReadAll()
	_ = config.GetConf(conf)
	err := moonConfig.InitConfig(conf)
	if err != nil {
		logger.Errorf("init config fail: %v", err)
	}

	// init moon node
	logger.Infof("Start init moon node ...")
	nodeInfoDir := "./NodeInfoStorage"
	infoStorage := node.NewStableNodeInfoStorage(nodeInfoDir)
	selfInfo := conf.SelfInfo
	rpcServer := messenger.NewRpcServer(conf.RpcPort)
	m := moon.NewMoon(selfInfo, conf.Moon.SunAddr, nil, nil,
		rpcServer)
	//init Alaya
	metaStorage := alaya.NewStableMetaStorage(nodeInfoDir)
	a := alaya.NewAlaya(selfInfo)

	go func() {
		err := rpcServer.Run()
		if err != nil {
			logger.Errorf("RPC server run error: %v", err)
		}
	}()
	go m.Run()
	Moon = m
	logger.Infof("moon node init success")

	// init Gin
	router := NewRouter()
	port := strconv.FormatUint(conf.HttpPort, 10)
	_ = router.Run(":" + port)
	return nil
}

func main() {
	// init cli
	flags := []cli.Flag{
		&cli.PathFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "set config file",
			Value:   "./edge_node.json",
		},
		&cli.StringFlag{
			Name:    "sun",
			Aliases: []string{"s"},
			Usage:   "set ecos-sun addr",
			Value:   "",
		},
	}
	app := cli.App{
		Name:                 "ECOS-EdgeNode",
		Usage:                "",
		UsageText:            "",
		ArgsUsage:            "",
		Version:              "",
		Flags:                flags,
		EnableBashCompletion: true,
		BashComplete:         nil,
		Action:               action,
		Compiled:             time.Time{},
		Authors: []*cli.Author{{
			Name:  "Zhang Junhua",
			Email: "zhangjh@mail.act.buaa.edu.cn",
		}},
		Copyright: "Copyright BUAA 2022",
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Errorf("edge node run error: %v", err)
	}

}
