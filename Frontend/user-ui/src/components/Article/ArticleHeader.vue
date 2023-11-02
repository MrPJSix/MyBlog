<template>
    <div class="a-header">
        <h1>
            {{ title }}
        </h1>
        <div class="a-header-info">
            <span class="a-header-info-creted_time">{{ timeDisplay(created_at) }}</span>
            <!-- <span>{{ formatTime(updated_at) }}</span> -->
            <span class="a-header-info-views_count">{{ views_count }}浏览</span>
            <span class="a-header-info-likes_count">{{ likes_count }}喜欢</span>
            <span class="a-header-info-comments_count">{{ comments_count }}评论</span>
            <!-- <span>{{ tags.toString() }}</span> -->
        </div>
        <div class="a-header-writer">
            <div class="a-header-writer-avatar">
                <img :src="author.avatar_url">
            </div>
            <div class="a-header-writer-info">
                <div class="a-header-writer-info-name">{{ author.full_name }}</div>
                <span>粉丝: 99999&nbsp;&nbsp;文章: 999</span>
            </div>
            <button>+关注</button>
        </div>
    </div>
    <hr>
</template>

<script setup lang='ts'>
import { Author } from '@/interface/index'
import { timeDisplay } from '@/utils/time'
import { watch, reactive } from 'vue'

const author = reactive({} as Author)
const props = defineProps<{
    title: string
    author: {
        user_id: number
        full_name: string
        avatar_url: string
    }
    created_at: string
    updated_at: string
    comments_count: number
    likes_count: number
    views_count: number
    tags: string[]
}>()
watch(() => props.author, (val) => {
    author.user_id = val.user_id
    author.full_name = val.full_name
    author.avatar_url = val.avatar_url
})
</script>

<style scoped></style>