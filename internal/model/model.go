package model

import (
	"blog/global"
	"blog/pkg/setting"
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// Model 公共Model
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"` //ID
	CreatedBy  string `json:"created_by"`            //创建者
	ModifiedBy string `json:"modified_by"`           //修改者
	CreateOn   uint32 `json:"create_on"`             //创建时间
	ModifiedOn uint32 `json:"modified_on"`           //修改时间
	DeleteOn   uint32 `json:"delete_on"`             //删除时间
	IsDel      uint8  `json:"is_del"`                //是否删除   0为删除  1为未删除
}

// Tag  标签Model
type Tag struct {
	*Model
	Name  string `json:"name"`  //标签名
	State uint8  `json:"state"` //标签状态   0为禁用  1为启用
}

// Article 文章Model
type Article struct {
	*Model
	Title         string `json:"title"`           //文章标题
	Desc          string `json:"desc"`            //文章简述
	Content       string `json:"content"`         //文章内容
	CoverImageUrl string `json:"cover_image_url"` //封面图片地址
	State         uint8  `json:"state"`           //文章状态	0为禁用  1为启用
}

// ArticleTag 文章标签
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)) //连接数据库
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" { //设置数据库的运行模式
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback) //注册回调函数
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns) //设置空闲连接池中的最大连接数。
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns) //设置到数据库的最大打开连接数。
	otgorm.AddGormCallbacks(db)                           //SQL追踪
	return db, nil
}

//新增函数行为的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	//如果操作数据库中没有错误
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeFiled, ok := scope.FieldByName("CreateOn"); ok { //找到 数据库 中“CreateOn”的字段
			if createTimeFiled.IsBlank { //如果创建时间为空
				_ = createTimeFiled.Set(nowTime) //设置为当前时间
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok { //找到“ModifiedOn”字段
			if modifyTimeField.IsBlank { //如果修改时间为空
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

//更新行为的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok { //设置“gorm:update_column”的字段属性
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//删除行为的回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() { //检查是否有错误
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deleteOnField, hasDeleteOnField := scope.FieldByName("DeleteOn")    //找到“DeleteOn”字段，是否删除
		isDelField, hasIsDeleteField := scope.FieldByName("IsDel")          //与上同理
		if !scope.Search.Unscoped && hasDeleteOnField && hasIsDeleteField { //判断是否存在 DeletedOn 和 IsDel 字段，若存在则调整为执行 UPDATE 操作进行软删除
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(), //获取当前表名
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return "" + str
	}
	return ""
}
