package server

import (
	"go_crud/common/constants"
	"go_crud/common/models"
	"go_crud/common/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	var newPost models.Post
	err := c.ShouldBindJSON(&newPost)
	if utils.CreateHttpError(c, err, http.StatusBadRequest, "Invalid request body") {
		return
	}

	newPost.Username = c.MustGet(constants.GinContextE.UserName).(string)

	postModel, _ := models.ModelMetaMap[constants.ModelsE.PostModel].(*models.PostModelMeta)
	record := postModel.CreatePost(newPost).Scan(&newPost)

	if record.Error != nil {
		utils.CreateHttpError(c, record.Error, http.StatusBadRequest, "Error creating post")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
		"title":   newPost.Title,
		"slug":    newPost.Slug,
	})

}

func GetPostBySlugHandler(c *gin.Context) {
	slug := c.Param("slug")
	postModel, _ := models.ModelMetaMap[constants.ModelsE.PostModel].(*models.PostModelMeta)
	var post models.Post
	record := postModel.GetPostBySlug(slug, &post)
	if utils.CreateHttpError(c, record.Error, http.StatusBadRequest, "Error fetching post") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Post fetched successfully",
		"title":    post.Title,
		"body":     post.Body,
		"slug":     post.Slug,
		"username": post.Username,
	})
}

func FetchPostsBulk(c *gin.Context) {
	start := c.Query("start")
	limit := c.Query("limit")
	if start == "" {
		start = "0"
	}
	if limit == "" {
		limit = "10"
	}

	iStart, sErr := strconv.Atoi(start)
	iLimit, lErr := strconv.Atoi(limit)
	if sErr != nil || lErr != nil {
		utils.CreateHttpError(c, nil, http.StatusBadRequest, "Invalid query params")
		return
	}

	var posts = []models.Post{}
	postModel, _ := models.ModelMetaMap[constants.ModelsE.PostModel].(*models.PostModelMeta)
	postModel.FetchPostsBulk(iStart, iLimit, posts).Scan(&posts)

	c.JSON(http.StatusOK, gin.H{
		"message": "Posts fetched successfully",
		"posts":   posts,
	})
}

func UpdatePostHandler(c *gin.Context) {}
