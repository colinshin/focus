// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"focus/app/model/internal"
)

// Category is the golang structure for table gf_category.
type Category internal.Category

type CategoryItem struct {
	Id       uint   `json:"id"`        // 分类ID，自增主键
	ParentId uint   `json:"parent_id"` // 父级分类ID，用于层级管理
	Name     string `json:"name"`      // 分类名称
	Thumb    string `json:"thumb"`     // 封面图
	Brief    string `json:"brief"`     // 简述
}

// API列表请求
type CategoryApiListReq struct {
	ContentType string `v:"required#请输入分类类型"` // 分类类型
	ParentId    int    // 父级ID
}

// API查询详情请求
type CategoryApiItemReq struct {
	Id uint `json:"id" v:"min:1#请输入分类ID"` // 分类ID
}

// DAO查询分类列表
type CategoryDaoGetListReq struct {
	ParentId    uint   // 父级分类ID，用于层级管理
	ContentType string // 内容类型：topic, question, article
}