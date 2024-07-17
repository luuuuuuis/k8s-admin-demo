package dao

import (
	"errors"
	"k8s-server/db"
	"k8s-server/model"

	"k8s-server/utils"
)

var Workflow workflow

type workflow struct{}

// 定义列表的返回内容，Items是workflow元素列表，Total为workflow元素数量
type WorkflowResp struct {
	Items []*model.Workflow `json:"items"`
	Total int               `json:"total"`
}

// 获取列表分页查询
func (w *workflow) GetList(name string, page, limit int) (data *WorkflowResp, err error) {
	//定义分页数据的起始位置
	startSet := (page - 1) * limit

	//定义数据库查询返回内容
	var (
		workflowList []*model.Workflow
		total        int
	)

	//数据库查询，Limit方法用于限制条数，Offset方法设置起始位置
	tx := db.GORM.
		Model(&model.Workflow{}).
		Where("name like ?", "%"+name+"%").
		Unscoped().
		Count(&total).
		Limit(limit).
		Offset(startSet).
		Order("id desc").
		Find(&workflowList)
	//gorm会默认把空数据也放到err中，故这里要排除空数据的情况
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		utils.Logger.Error().Stack().Err(errors.New("获取Workflow列表失败, ")).Msg(tx.Error.Error())
		return nil, errors.New("获取Workflow列表失败, " + tx.Error.Error())
	}

	return &WorkflowResp{
		Items: workflowList,
		Total: total,
	}, nil
}

// 查询workflow单条数据
func (w *workflow) GetById(id int) (workflow *model.Workflow, err error) {
	workflow = &model.Workflow{}
	tx := db.GORM.Where("id = ?", id).First(&workflow)
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		utils.Logger.Error().Stack().Err(errors.New("获取Workflow单条数据失败, ")).Msg(tx.Error.Error())
		return nil, errors.New("获取Workflow单条数据失败, " + tx.Error.Error())
	}
	return
}

// 新增workflow
func (w *workflow) Add(workflow *model.Workflow) (err error) {
	tx := db.GORM.Create(&workflow)
	if tx.Error != nil {
		utils.Logger.Error().Stack().Err(errors.New("添加Workflow失败, ")).Msg(tx.Error.Error())
		return errors.New("添加Workflow失败, " + tx.Error.Error())
	}
	return nil
}

// 删除workflow
// 软删除 db.GORM.Delete("id = ?", id)
// 软删除执行的是UPDATE语句，将deleted_at字段设置为时间即可，gorm 默认就是软删。
// 实际执行语句 UPDATE `workflow` SET `deleted_at` = '2021-03-01 08:32:11' WHERE `id` IN ('1'
// 硬删除 db.GORM.Unscoped().Delete("id = ?", id)) 直接从表中删除这条数据
// 实际执行语句 DELETE FROM `workflow` WHERE `id` IN ('1');
func (w *workflow) DelById(id int) (err error) {
	tx := db.GORM.Where("id = ?", id).Delete(&model.Workflow{})
	if tx.Error != nil {
		utils.Logger.Error().Stack().Err(errors.New("删除Workflow失败, ")).Msg(tx.Error.Error())
		return errors.New("删除Workflow失败, " + tx.Error.Error())
	}
	return nil
}
