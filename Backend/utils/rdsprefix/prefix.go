package rdsprefix

// 统一规定Redis各类数据类型的前缀

const (
	ArticleLikeSet  string = "set:likes:article:" // + articleID
	ArticleLikeSync string = "sync:likes:article:"

	UserArticleStarSet  string = "set:stars:art:user:" // + userID
	UserArticleStarSync string = "sync:stars:art:user:"

	UserFollowSet  string = "set:follows:user:"
	UserFollowList string = "list:follows:user:"
	UserFansList   string = "list:fans:user:"
	UserFollowSync string = "sync:follows:user:"

	CommentLikeSet  string = "set:likes:comment:"
	CommentLikeSync string = "sync:likes:comment:"
)
