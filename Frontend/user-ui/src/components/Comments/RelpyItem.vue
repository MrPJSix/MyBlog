<template>
    <div class="r-l-item">
        <a class="reply-face" href="">
            <div class="reply-avatar">
                <img :src="reply.user.avatar_url" alt="">
            </div>
        </a>
        <div class="reply-con-wrap">
            <div class="reply-con">
                <span class="user">{{ reply.user.full_name }}</span>
                <span class="text">
                    <span v-if="reply.replied_user">回复</span>
                    <a href="">@{{ reply.replied_user.full_name }}</a>
                    <span>:</span>
                    {{ reply.content }}
                </span>
            </div>
        </div>

        <div class="reply-info">
            <span class="time-location">{{ timeDisplay(reply.create_at) }}</span>
            <span class="like">
                <i></i>
                {{ reply.likes }}
            </span>
            <span class="hate">
                <i></i>
            </span>
            <span class="reply-btn" @click="showReplyPostBox">回复</span>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { ReplyReciveData } from '@/interface/index'
import { timeDisplay } from '@/utils/time' 

const props = defineProps<{
    reply: ReplyReciveData
}>()

const emit = defineEmits<{
    (e: 'showReplyPostBox', id: number, name: string): void
}>()


function showReplyPostBox() {
    emit('showReplyPostBox', props.reply.id, props.reply.user.full_name)
}


</script>

<style scoped></style>