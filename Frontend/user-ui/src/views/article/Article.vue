<template>
  <div class="article">
    <ArticleHeader 
        :title="article.title" 
        :author="article.author" 
        :created_at="article.created_at"
        :updated_at="article.updated_at"
        :comments_count="article.comment_count"
        :likes_count="article.likes"
        :views_count="article.read_count"
        :tags="['Python', 'Java']"
    ></ArticleHeader>
    <ActicleBody
        :content="article.content" 
        :content_type="article.content_type"   
    ></ActicleBody>
    <ArticleFooter></ArticleFooter>
    <ArticleSideBar></ArticleSideBar>
  </div>
</template>

<script setup lang="ts">
import ArticleHeader from '@/components/Article/ArticleHeader.vue'
import ActicleBody from '@/components/Article/ActicleBody.vue'
import ArticleFooter from '@/components/Article/ArticleFooter.vue'
import ArticleSideBar from '@/components/Article/ArticleSideBar.vue'
import { reactive, ref } from 'vue'
import api from '@/api/index'
import { ArticleReciveData } from '@/interface'
import { useRouter } from 'vue-router'

const router = useRouter()

const articleId = Number(router.currentRoute.value.params.id)

const article = ref({} as ArticleReciveData)

api.article.getArticleById(articleId).then(res => {
  article.value.id = res.data.id
  article.value.created_at = res.data.created_at
  article.value.updated_at = res.data.updated_at
  article.value.title = res.data.title
  article.value.content = res.data.content
  article.value.img = res.data.img
  article.value.comment_count = res.data.comment_count
  article.value.read_count = res.data.read_count
  article.value.likes = res.data.likes
  article.value.category = res.data.category
  article.value.author = res.data.author
  article.value.content_type = res.data.content_type
})

const demo = reactive({
  id: 1,
  created_at: '1695797349',
  updated_at: '1695797349',
  title: '深入浅出Python编程',
  desc: '一篇介绍Python的核心概念和应用的文章。',
  content: 'Python是一种广泛使用的高级编程语言，著名的简洁明了...',
  img: '',
  comment_count: 1,
  read_count: 2,
  category: {
    category_id: 1,
    category_name: 'Python',
  },
  author: {
    id: 3,
    name: 'User',
    avatar: 'http://117.72.17.185:9199/myblog-img/3_1696962640545445297.png',
  },
})
</script>

<style lang="less" scoped>
</style>
