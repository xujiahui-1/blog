package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

//页面的固定数据，都是写死在配置文件里
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

//系统
type SystemCongfig struct {
	AppName        string
	Version        float32
	CurrentDir     string //目录
	CdnURL         string
	QiniuAccesskey string
	QiniuSecretKey string
	//评论相关
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerKey string
	ValineServerURL string
}

//toml配置文件的映射
type TomlConfig struct {
	Viewer Viewer
	System SystemCongfig
}

var Cfg *TomlConfig

func init() {
	Cfg = new(TomlConfig)
	Cfg.System.AppName = "my0-blog"
	Cfg.System.Version = 1.0
	getwd, _ := os.Getwd()
	Cfg.System.CurrentDir = getwd //目录
	_, err := toml.DecodeFile("config/config.toml", Cfg)
	if err != nil {
		panic(err)
	}
}
