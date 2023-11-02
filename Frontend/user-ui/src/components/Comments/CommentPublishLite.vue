<template>
    <div class="c-body-publish-lite" style="display: none;">
        <div class="textarea-container">
            <textarea class="ipt-txt" placeholder="发表你的看法" @select="addFocus" v-model="inputValue"></textarea>
            <button class="comment-submit" @click="postComment">发表评论</button>
        </div>
        <div class="comment-emoji-lite">
            <i class="face"></i>
            <span class="text">表情</span>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api/index'
const router = useRouter()

const articleId = Number(router.currentRoute.value.params.id)
const inputValue = ref('')

let container = null as null | HTMLDivElement
onMounted(() => {
    container = document.querySelector('.textarea-container')
    console.log(container);
})

const addFocus = () => {
    console.log(11);
    container!.classList.add('foucs')
    console.log('focus');
}

const postComment = () => {
    api.comment.postComment({
        article_id: articleId,
        content: inputValue.value
    }).then(res => {
        if(res.status === 200) {
            window.location.reload()
        }
    })
}

</script>

<style scoped>
.foucs {
    height: 104px;
}
</style>