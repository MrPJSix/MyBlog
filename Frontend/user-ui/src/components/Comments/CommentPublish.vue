<template>
    <div class="c-body-publish">
        <div class="c-body-publish-container">
            <div class="c-body-publish-container-avartar">
                <img :src="store.avatar_url">
            </div>
            <textarea :placeholder="placeholder" class="c-body-publish-container-input" v-model="inputValue"></textarea>
            <button @click="postComment">
                <span>发表评论</span>
            </button>
        </div>
        <button class="c-body-pulish-emoji">表情</button>
    </div>
</template>

<script setup lang='ts'>
import api from '@/api/index'
import { useRouter } from 'vue-router'
import { useUserDataStore } from '@/stores/modules/userData'
import { ref, computed, watch } from 'vue'
const store = useUserDataStore()

const router = useRouter()

const articleId = Number(router.currentRoute.value.params.id)

const inputValue = ref('')

const postComment = () => {
    if (props.replyTargetId) {
        api.comment.postReply({
            comment_id: props.replyTargetId,
            article_id: articleId,
            content: inputValue.value
        }).then(res => {
            if (res.status === 200) {
                window.location.reload()
            }
        })
    } else {
        api.comment.postComment({
            article_id: articleId,
            content: inputValue.value
        }).then(res => {
            if (res.status === 200) {
                window.location.reload()
            }
        })
    }
}

const props = defineProps<{
    replyTargetId?: number
    replyTargetName?: string
}>()

const placeholder = computed(() => {
    return props.replyTargetName ? 
        `回复${props.replyTargetName}` : 
        '发表一条友善的评论'
})

watch(() => props.replyTargetId, () => {
    
    inputValue.value = ''

})

</script>

<style scoped></style>