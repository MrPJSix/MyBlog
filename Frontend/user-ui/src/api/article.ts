import request from "@/utils/request"
import qs from 'qs'
import type { Config, ArticlePostData } from '@/interface/index'
export const getArticlesByCategory = (id:number, config:Config) => {
    const queryString = qs.stringify(config)
    return request({
        url: `/category/${id}/articles?${queryString}`,
        method: 'get',
    })
}

export const getgetArticlesCountByCategory = (id:number) => {
    return request({
        url: `/category/${id}/articles/count`,
        method: 'get',
    })
}

export const getAllArticles = (config:Config) => {
    const queryString = qs.stringify(config)
    return request({
        url: `/articles?${queryString}`,
        method: 'get',
    })
}

export const getAllArticlesCount = () => {
    return request({
        url: 'articles/count',
        method: 'get',
    })
}


export const getArticleById = (id:number) => {
    return request({
        url: `/article/${id}`,
        method: 'get',
    })
}

export const getArticleByUser = (id:number, config:Config) => {
    const queryString = qs.stringify(config)
    return request({
        url: `/user/${id}/articles?${queryString}`,
        method: 'get',
    })
}

export const postArticle = (articlePostData: ArticlePostData) => {
    return request({
        url: 'article',
        method: 'post',
        data: articlePostData
    })
}

export const deleteArticleById = (id:number) => {
    return request({
        url: `/article/${id}`,
        method: 'delete',
    })
}