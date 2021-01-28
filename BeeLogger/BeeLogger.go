package BeeLogger

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func IntLoggger() (err error) {
	//直接beego.AppConfig.String获取
	logConf := make(map[string]interface{})
	logConf["filename"] = beego.AppConfig.DefaultString("log::log_path", "logs/app.log")
	logConf["level"] = beego.AppConfig.DefaultInt("log::log_level", 7)
	logConf["maxlines"] = beego.AppConfig.DefaultInt("log::maxlines", 10000)
	//日志文件大小，单位B  ,配置改为单位M
	logConf["maxsize"] = beego.AppConfig.DefaultInt("log::maxsize", 10) * 1024 * 1024
	logConf["daily"] = beego.AppConfig.DefaultBool("log::daily", true)
	logConf["Maxdays"] = beego.AppConfig.DefaultInt("log::maxdays", 365)
	lFuncCall := beego.AppConfig.DefaultBool("log::funcCall", true)

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	beego.SetLogger(logs.AdapterFile, string(confStr))
	beego.SetLogFuncCall(lFuncCall)
	beego.Debug("初始化日志配置")
	return
}
