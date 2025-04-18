package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 评论

// CommentHandler 创建评论
func CommentHandler(c *gin.Context) {
	var comment models.Comment
	if err := c.BindJSON(&comment); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 生成评论ID
	commentID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	comment.CommentID = commentID
	comment.AuthorID = userID

	// 创建评论
	if err := mysql.CreateComment(&comment); err != nil {
		zap.L().Error("mysql.CreateComment(&comment) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// CommentListHandler 评论列表
func CommentListHandler(c *gin.Context) {
	// 获取帖子ID
	postIDStr := c.Query("post_id")
	if postIDStr == "" {
		ResponseError(c, CodeInvalidParams)
		return
	}
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 获取帖子下的所有评论
	comments, err := mysql.GetCommentListByPostID(postID)
	if err != nil {
		zap.L().Error("mysql.GetCommentListByPostID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, comments)
}
