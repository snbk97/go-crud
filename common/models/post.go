package models

import (
	zlog "go_crud/common/logger"
	"strings"

	"github.com/gosimple/slug"

	"gorm.io/gorm"
)

type PostModelMeta struct {
	DB *gorm.DB
}

var logger = zlog.CreateLogger("PostModelMeta")

func (meta *PostModelMeta) Init(db *gorm.DB) interface{} {
	return &PostModelMeta{DB: db}
}

type Post struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null"`
	Body     string `json:"body" gorm:"not null"`
	ImageUrl string `json:"image_url"`
	Slug     string `json:"slug" gorm:"unique;not null"`
	Username string `json:"username" gorm:"not null"`
	User     User   `json:"user" gorm:"foreignKey:Username;references:Username"`
}

func (meta *PostModelMeta) CreatePost(post Post) *gorm.DB {
	var u User
	_slug := strings.Join([]string{post.Username, post.Title}, "-")

	post.Slug = slug.Make(_slug)
	foundUser := meta.DB.Model(&User{}).Where("username = ?", post.Username).First(&u)
	if foundUser.Error != nil {
		panic("User not found")
	}
	post.Username = u.Username
	logger.LogInfo("Creating post with slug: ", post.Slug)
	return meta.DB.Model(&Post{}).Create(&post)
}

func (meta *PostModelMeta) GetPostBySlug(slug string, dest interface{}) *gorm.DB {
	return meta.DB.Model(&Post{}).Where("slug = ?", slug).Where("deleted_at IS NULL").First(&dest)
}

func (meta *PostModelMeta) GetPostByUserName(username string, dest interface{}) *gorm.DB {
	return meta.DB.Model(&Post{}).Where("username = ?", username).Where("deleted_at IS NULL").Scan(&dest)
}

func (meta *PostModelMeta) FetchPostsBulk(start, limit int, dest interface{}) *gorm.DB {
	return meta.DB.Model(&Post{}).Where("deleted_at IS NULL").Offset(start).Limit(limit).Find(&dest)
}
