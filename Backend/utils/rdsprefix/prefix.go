package rdsprefix

// 统一规定Redis各类数据类型的前缀

const (
	ArticleLikeSet string = "set:likes:article:" // + articleI
	CommentLikeSet string = ""
)
