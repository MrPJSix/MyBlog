<template>
    <div class="user-center-exact">
        <div class="user-avatar">
            <img class="user-avatar-img" :src="selfProfile.avatar_url" alt="">
        </div>
        <div class="user-profile">
            <div class="user-id">
                <h1>{{selfProfile.full_name}}</h1>
            </div>
            <div class="user-bio">{{ selfProfile.bio }}</div>
            <div class="user-info">
                <p>
                    <span>性别</span>未知
                </p>
                <p>
                    <span>发帖数</span>283
                </p>
                <p>
                    <span>粉丝数</span>99999
                </p>
                <p>
                    <span>点赞数</span>9999
                </p>
            </div>
        </div>
        <div class="user-blogs"></div>
    </div>
</template>

<script setup lang='ts'>
import api from '@/api/index'
import { reactive } from 'vue'
import { SelfProfile } from '@/interface'
const selfProfile = reactive({} as SelfProfile)
api.user.getSelfProfile().then(res => {
    selfProfile.id = res.data.id
    selfProfile.username = res.data.username
    selfProfile.full_name = res.data.full_name
    selfProfile.register_date = res.data.register_date
    selfProfile.bio = res.data.bio
    selfProfile.role = res.data.role
    selfProfile.avatar_url = res.data.avatar_url
})
</script>

<style scoped lang="less">

</style>