<template>
    <div class="c-body-display">
        <CommentsItem v-for="(comment,index) in comments" :key="index" :comment="comment"></CommentsItem>
    </div>
</template>

<script setup lang='ts'>
import CommentsItem from './CommentsItem.vue';
import { useRouter } from 'vue-router'
import { reactive } from 'vue'
import api from '@/api/index'
import { CommentReciveData } from '@/interface/index'

const router = useRouter()

const ArticleId = Number(router.currentRoute.value.params.id)

const comments = reactive([] as CommentReciveData[])

api.comment.getComments(ArticleId).then(res => {
    if(res.data && res.data.length !== 0){
        res.data.forEach((item: CommentReciveData) => {
            comments.push(item)
        })
    }
})

</script>

<style scoped>
</style>