<template>
    <div class="c-body">
        <div class="c-body-bar">
            <span class="active">按热度排序</span>
            <span>按时间排序</span>
        </div>
        <CommentPublish></CommentPublish>
        <CommentsDisplay></CommentsDisplay>
        <CommentPublishLite></CommentPublishLite>
    </div>
</template>

<script setup lang='ts'>
import CommentPublish from './CommentPublish.vue';
import CommentPublishLite from './CommentPublishLite.vue';
import CommentsDisplay from './CommentsDisplay.vue';
import { onMounted } from 'vue'

let cb: null | HTMLDivElement = null
let c1: null | HTMLDivElement = null
let c2: null | HTMLDivElement = null

onMounted(() => {

    c1 = document.querySelector('.c-body-publish')
    
    c2 = document.querySelector('.c-body-publish-lite')
    
})

window.addEventListener('scroll', () => {
    // 上方评论框 滚动到超出屏幕时 显示下方评论框
    if((c1?.getBoundingClientRect().top as number) + (c1?.offsetHeight as number) < 0){
        c2!.style.display = 'block'
    } else {
        c2!.style.display = 'none'
    }
})

</script>

<style scoped>
</style>