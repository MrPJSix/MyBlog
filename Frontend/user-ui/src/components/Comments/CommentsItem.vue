<template>
    <div class="c-body-display-item">
        <div class="user-face">
            <a href="">
                <div class="user_avatar">
                    <img v-if="comment.user.avatar_url" :src="comment.user.avatar_url">
                </div>
            </a>
        </div>
        <div class="content">
            <div class="comment">
                <div class="com-header">
                    <a href="">
                        {{ comment.user.full_name }}
                    </a>
                </div>
                <p class="com-body">
                    {{ comment.content }}
                </p>
                <div class="com-footer">
                    <span class="time-location">{{ timeDisplay(comment.create_at) }}</span>
                    <span class="like">
                        <i></i>
                        {{ comment.likes }}
                    </span>
                    <span class="hate">
                        <i></i>
                    </span>
                    <span class="reply-btn" @click="showReply">回复</span>
                   
                </div>
            </div>
            <div class="reply-list">
                <RelpyItem v-for="(reply, index) in replies" 
                    :key="index" 
                    :reply="reply" 
                    @showReplyPostBox="showReplyFromSon"
                ></RelpyItem>
                <div v-if="comment.total_replies > 3 && isFold" class="view-more">共{{ comment.total_replies}}条回复
                    <a class="btn-more" @click="viewAllReplies">点击查看</a>
                </div>
                <ElPagination 
                    v-if="!isFold && comment.total_replies > pagesize"
                    small
                    layout="prev, pager, next"
                    :total="comment.total_replies"
                    v-model:current-page="pagenum"
                    @current-change="handleCurrentChange"   
                ></ElPagination>
            </div>
            <CommentPublish v-if="isShowPublish" 
                :replyTargetId="replyTargetId"
                :replyTargetName="replyTargetName"
            ></CommentPublish>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { CommentReciveData, ReplyReciveData } from '@/interface';
import RelpyItem from './RelpyItem.vue';
import { timeDisplay } from '@/utils/time';
import { ref } from 'vue'
import api from '@/api/index';
import CommentPublish from './CommentPublish.vue';
import { ElPagination } from 'element-plus'

/* 页码  */
const pagenum = ref(1)
/* 每页条数 */
const pagesize = ref(10)

const props = defineProps<{
    comment: CommentReciveData
}>()

const isShowPublish = ref(false) // 是否显示回复框

const isFold = ref(true) // 评论是否折叠

let replies = ref([] as ReplyReciveData[])

// 当前页面渲染的回复列表
replies.value = props.comment.replies



const viewAllReplies = () => {
    api.comment.getRepliesOfRootComment(props.comment.id, {
        pagenum: 1,
        pagesize: 10
    }).then(res => {
        replies.value = res.data
        isFold.value = false
    })
}

// 子组件触发显示回复框事件 参数为 待回复的 回复id
const showReplyFromSon = (id:number, name:string) => {
    isShowPublish.value = true // 显示回复框
    replyTargetId.value = id // 设置回复目标id
    replyTargetName.value = name // 设置回复目标name
}

const showReply = () => {
    isShowPublish.value = true
    replyTargetId.value = props.comment.id // 根评论id
    replyTargetName.value = props.comment.user.full_name // 根评论用户名
}

const replyTargetId = ref(-1)
const replyTargetName = ref('')

const handleCurrentChange = (next:number) => {
    api.comment.getRepliesOfRootComment(props.comment.id, {
        pagenum: next,
        pagesize: pagesize.value
    }).then(res => {
        if(res.status === 200){
            replies.value = res.data
        }
    })
}

</script>

<style scoped></style>