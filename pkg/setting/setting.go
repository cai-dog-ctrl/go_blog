package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// NewSetting 返回一个Setting实例
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()           //返回一个viper实例
	vp.SetConfigName("configs") //给配置文件设置一个名子
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml") //设置返回的配置类型
	err := vp.ReadInConfig() //加载配置文件
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
