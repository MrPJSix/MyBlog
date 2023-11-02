<template>
    <div class="header">
        <div class="header-right">
            <div class="header-login">
                <div v-if="!IsLogin" class="header-login-entry" @click="login">
                    <span>登录</span>
                </div>
                <div v-else class="header-avatar-wrap" @click="toUserCenter">
                    <img :src=store.avatar_url alt="">
                    <a href=""></a>
                    <!-- 默认头像 如果非登录状态 -->
                    <!-- <div class="blog-avatar"></div> -->
                </div>
            </div>
            <header-icon @click="dayMode = !dayMode">
                <img v-if="dayMode" src="@/assets/svg/sun.svg" alt="sun">
                <img v-else src="@/assets/svg/moon.svg" alt="moon">
            </header-icon>
            <header-icon>
                <img src="@/assets/svg/setting.svg" alt="setting">
            </header-icon>
            <header-icon>
                <img src="@/assets/svg/language.svg" alt="language">
            </header-icon>
            <header-icon>
                <!-- <img src="@/assets/svg/language.svg" alt="language"> -->
                <!-- <img src="@/assets/svg/github-night.svg" alt="github-night"> -->
                <img class="svg-icon" src="@/assets/svg/github.svg" alt="github-day">
            </header-icon>
        </div>
    </div>
</template>

<script setup lang='ts'>
import HeaderIcon from '@/views/home/componenets/HeaderIcon.vue';
// import svg_github-day from '@/assets/svg/github-day.svg';
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import {useUserDataStore} from '@/stores/modules/userData'

const store = useUserDataStore()
store.updateUserData()
const router = useRouter()
const IsLogin = ref(false)
const login = () => {
    router.push('/login')
}
const toUserCenter = () => {
    router.push('/usercenter')
}
if (localStorage.getItem('token')) {
    IsLogin.value = true
}
const dayMode = ref(true)
</script>

<style scoped lang="less">
@import url('@/assets/less/header.less');

.svg-icon {
    width: 36px;
    height: 36px;
}
</style>