package service

import (
	"blog/global"
	"blog/internal/dao"
	"context"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine)) //SetSpanToGorm 将 span 设置为 gorm 设置，返回克隆的 DB
	return svc
}
