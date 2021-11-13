package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// NewSetting 返回一个Setting实例
func NewSetting()(*Setting,error){
	vp:=viper.New()//返回一个viper实例
	vp.SetConfigName("configs") //给配置文件设置一个名子
	vp.AddConfigPath("config/") //为 Viper 添加路径以在其中搜索配置文件
	vp.SetConfigType("yaml")  //设置返回的配置类型
	err:=vp.ReadInConfig() //加载配置文件
	if err!=nil{
		return nil,err
	}
	return &Setting{vp},err
}
