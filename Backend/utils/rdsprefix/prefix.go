package rdsprefix

// 统一规定Redis各类数据类型的前缀

const (
	ArticleLikeSet  string = "set:likes:article:" // + articleID
	ArticleLikeSync string = "is_syncing:set:likes:article:"

	CommentLikeSet  string = "set:likes:comment:"
	CommentLikeSync string = "is_syncing:set:likes:comment:"
)
