package dao

import (
	"blog/internal/model"
	"blog/pkg/app"
)

func (d *Dao)CountTag(name string,state uint8)(int ,error){
	tag:=model.Tag{Name: name,State: state}   //将Tag的name和state封装为结构体
	return tag.Count(d.engine)    //查询Tag的记录数
}

func(d *Dao)GetTagList(name string,state uint8,page,pageSize int)([]*model.Tag,error){
	tag:=model.Tag{Name: name,State: state}   //封装一个Tag作为参数
	pageOffset:=app.GeyPageOffset(page,pageSize)
	return tag.List(d.engine,pageOffset,pageSize)   //根据pageOffset和pageSize查询Tag列表
}

func(d *Dao)CreateTag(name string,state uint8,createdBy string)error{
	tag:=model.Tag{
		Name: name,
		State: state,
		Model:&model.Model{CreatedBy: createdBy},
	}                 //封装tag

	return tag.Create(d.engine)   //创建Tag
}

func (d *Dao)UpdateTag(id uint32,name string,state uint8,modifiedBy string)error{
	tag:=model.Tag{
		Name: name,
		State: state,
		Model:&model.Model{ID: id,ModifiedBy: modifiedBy},
	}      //这些基本都一样
	return tag.Update(d.engine)
}

func (d *Dao)DeleteTag(id uint32)error{
	tag:=model.Tag{Model:&model.Model{
		ID: id,
	}}
	return tag.Delete(d.engine)
}
