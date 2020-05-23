package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// @title    CreatePost
// @description   create a post, 新增post
// @auth                     （2020/04/05  20:22）
// @param     post             model.Post
// @return                    error

func CreatePost(post model.Post) (err error) {
	err = global.GVA_DB.Create(&post).Error
	return err
}
