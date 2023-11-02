import request from "@/utils/request";
import type { CommentPostData, Config, ReplyPostData } from '@/interface/index'
import qs from 'qs'


export const getComments = (id:number) => {
    return request({
        url: `/article/${id}/comments`,
        method: 'get',
    })
}

export const postComment = (commentPostData: CommentPostData) => {
    return request({
        url: `/article/${commentPostData.article_id}/comment`,
        method: 'post',
        data: commentPostData,
    })
}

export const getRepliesOfRootComment = (id:number, config: Config) => {
    return request({
        url: `/comment/${id}/replies?${qs.stringify(config)}`,
        method: 'get',
    })
}

export const postReply = (data: ReplyPostData) => {
    return request({
        url: `/article/${data.article_id}/comment/${data.comment_id}/reply`,
        method: 'post',
        data: {
            content: data.content
        }
    })
}