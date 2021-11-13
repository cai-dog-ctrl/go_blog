package model

import (
	"blog/pkg/app"
	"github.com/jinzhu/gorm"
)

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}

// Count 返回Tag记录数
func(t Tag)Count(db *gorm.DB)(int ,error){
	var count int
	if t.Name!="" {
		db=db.Where("name = ?",t.Name)
	}
	db=db.Where("state = ? ",t.State)
	if err:=db.Model(&t).Where("is_del = ?",0).Count(&count).Error;err!=nil{
		return 0,err
	}
	return count ,nil
}

// List 分页查询Tags
func (t Tag)List(db *gorm.DB,pageOffset ,pageSize int )([]*Tag,error){
	var tags[]*Tag
	var err error
	if pageOffset >=0&&pageSize>0{
		db=db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name!=""{
		db=db.Where("name = ?",t.Name)
	}
	db=db.Where("state = ?",t.State)
	if err=db.Where("is_del = ?",0).Find(&tags).Error;err!=nil{
		return nil,err
	}
	return tags,nil
}
func (t Tag)Create(db *gorm.DB)error{
	return db.Create(&t).Error
}

func (t Tag)Update(db *gorm.DB)error{
	return db.Model(&Tag{}).Where("id = ? AND is_del = ?",t.ID,0).Update(t).Error
}

func (t Tag)Delete(db *gorm.DB)error{
	return db.Where("id = ? AND is_del = ?",t.Model.ID,0).Delete(&t).Error
}


