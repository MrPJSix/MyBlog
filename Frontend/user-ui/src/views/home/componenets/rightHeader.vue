<template>
    <div class="home-header">
        <div class="home-header-right">
            <el-button class="post-btn" type="primary" @click="toPost">写文章</el-button>
            <div class="home-header-login">
                <div v-if="!IsLogin" class="home-header-login-entry" @click="login">
                    <span>登录</span>
                </div>
                <div v-else class="home-header-avatar-wrap" @click="toUserCenter">
                    <img :src="store.avatar_url" alt="">
                    <!-- <a href=""></a> -->
                </div>
            </div>
            <header-icon>
                <svg class="h-icon__svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" role="img"
                    aria-hidden="true">
                    <path
                        d="M17.75,4.09L15.22,6.03L16.13,9.09L13.5,7.28L10.87,9.09L11.78,6.03L9.25,4.09L12.44,4L13.5,1L14.56,4L17.75,4.09M21.25,11L19.61,12.25L20.2,14.23L18.5,13.06L16.8,14.23L17.39,12.25L15.75,11L17.81,10.95L18.5,9L19.19,10.95L21.25,11M18.97,15.95C19.8,15.87 20.69,17.05 20.16,17.8C19.84,18.25 19.5,18.67 19.08,19.07C15.17,23 8.84,23 4.94,19.07C1.03,15.17 1.03,8.83 4.94,4.93C5.34,4.53 5.76,4.17 6.21,3.85C6.96,3.32 8.14,4.21 8.06,5.04C7.79,7.9 8.75,10.87 10.95,13.06C13.14,15.26 16.1,16.22 18.97,15.95M17.33,17.97C14.5,17.81 11.7,16.64 9.53,14.5C7.36,12.31 6.2,9.5 6.04,6.68C3.23,9.82 3.34,14.64 6.35,17.66C9.37,20.67 14.19,20.78 17.33,17.97Z">
                    </path>
                </svg>
            </header-icon>
            <header-icon>
                <svg class="v-icon__svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" role="img"
                    aria-hidden="true">
                    <path
                        d="M12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,2L14.39,5.42C13.65,5.15 12.84,5 12,5C11.16,5 10.35,5.15 9.61,5.42L12,2M3.34,7L7.5,6.65C6.9,7.16 6.36,7.78 5.94,8.5C5.5,9.24 5.25,10 5.11,10.79L3.34,7M3.36,17L5.12,13.23C5.26,14 5.53,14.78 5.95,15.5C6.37,16.24 6.91,16.86 7.5,17.37L3.36,17M20.65,7L18.88,10.79C18.74,10 18.47,9.23 18.05,8.5C17.63,7.78 17.1,7.15 16.5,6.64L20.65,7M20.64,17L16.5,17.36C17.09,16.85 17.62,16.22 18.04,15.5C18.46,14.77 18.73,14 18.87,13.21L20.64,17M12,22L9.59,18.56C10.33,18.83 11.14,19 12,19C12.82,19 13.63,18.83 14.37,18.56L12,22Z">
                    </path>
                </svg>
            </header-icon>
            <header-icon>
                <svg class="v-icon__svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                    <path
                        d="M12,8A4,4 0 0,1 16,12A4,4 0 0,1 12,16A4,4 0 0,1 8,12A4,4 0 0,1 12,8M12,10A2,2 0 0,0 10,12A2,2 0 0,0 12,14A2,2 0 0,0 14,12A2,2 0 0,0 12,10M10,22C9.75,22 9.54,21.82 9.5,21.58L9.13,18.93C8.5,18.68 7.96,18.34 7.44,17.94L4.95,18.95C4.73,19.03 4.46,18.95 4.34,18.73L2.34,15.27C2.21,15.05 2.27,14.78 2.46,14.63L4.57,12.97L4.5,12L4.57,11L2.46,9.37C2.27,9.22 2.21,8.95 2.34,8.73L4.34,5.27C4.46,5.05 4.73,4.96 4.95,5.05L7.44,6.05C7.96,5.66 8.5,5.32 9.13,5.07L9.5,2.42C9.54,2.18 9.75,2 10,2H14C14.25,2 14.46,2.18 14.5,2.42L14.87,5.07C15.5,5.32 16.04,5.66 16.56,6.05L19.05,5.05C19.27,4.96 19.54,5.05 19.66,5.27L21.66,8.73C21.79,8.95 21.73,9.22 21.54,9.37L19.43,11L19.5,12L19.43,13L21.54,14.63C21.73,14.78 21.79,15.05 21.66,15.27L19.66,18.73C19.54,18.95 19.27,19.04 19.05,18.95L16.56,17.95C16.04,18.34 15.5,18.68 14.87,18.93L14.5,21.58C14.46,21.82 14.25,22 14,22H10M11.25,4L10.88,6.61C9.68,6.86 8.62,7.5 7.85,8.39L5.44,7.35L4.69,8.65L6.8,10.2C6.4,11.37 6.4,12.64 6.8,13.8L4.68,15.36L5.43,16.66L7.86,15.62C8.63,16.5 9.68,17.14 10.87,17.38L11.24,20H12.76L13.13,17.39C14.32,17.14 15.37,16.5 16.14,15.62L18.57,16.66L19.32,15.36L17.2,13.81C17.6,12.64 17.6,11.37 17.2,10.2L19.31,8.65L18.56,7.35L16.15,8.39C15.38,7.5 14.32,6.86 13.12,6.62L12.75,4H11.25Z">
                    </path>
                </svg>
            </header-icon>
            <header-icon>
                <svg class="v-icon__svg" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                    <path
                        d="M12.87,15.07L10.33,12.56L10.36,12.53C12.1,10.59 13.34,8.36 14.07,6H17V4H10V2H8V4H1V6H12.17C11.5,7.92 10.44,9.75 9,11.35C8.07,10.32 7.3,9.19 6.69,8H4.69C5.42,9.63 6.42,11.17 7.67,12.56L2.58,17.58L4,19L9,14L12.11,17.11L12.87,15.07M18.5,10H16.5L12,22H14L15.12,19H19.87L21,22H23L18.5,10M15.88,17L17.5,12.67L19.12,17H15.88Z">
                    </path>
                </svg>
            </header-icon>
            
        </div>
    </div>
</template>

<script setup lang='ts'>
import HeaderIcon from '@/views/home/componenets/HeaderIcon.vue'
import { ref, onBeforeMount } from 'vue'
import {useRouter} from 'vue-router'
import { useUserDataStore } from '@/stores/modules/userData'
import { ElButton } from 'element-plus'
const IsLogin = ref(false)
const router = useRouter()
const login = () => {
    router.push('/login')
}
const toUserCenter = () => {
    router.push('/usercenter')
}

const toPost = () => {
    router.push('/post')
}

onBeforeMount(() => {
    if (localStorage.getItem('token')) {
        IsLogin.value = true
    }
})

const store = useUserDataStore()

</script>

<style scoped lang="less">
.home-header {
    display: flex;
    align-items: center;
    height: 100%;
    font-size: 20px;
    position: relative;
}

.home-header-right {
    display: flex;
    position: absolute;
    right: 8px;
    align-items: center;

    .post-btn {
        height: 40px;
        line-height: 40px;
        padding: 5px 20px;
        font-size: 16px;
        margin-right: 20px;
        box-sizing: border-box;
    }

    .home-header-login {
        width: 50px;
        height: 50px;
        margin-right: 10px;
        cursor: pointer;
        display: flex;
        align-items: center;

        .home-header-login-entry {
            width: 40px;
            height: 40px;
            border-radius: 50%;
            color: #fff;
            background-color: #00AEEC;
            text-align: center;
            font-size: 15px;
            line-height: 40px;
            cursor: pointer;
            border: 2px solid #fff;
        }

        .home-header-avatar-wrap {
            display: flex;
            align-items: center;
            justify-content: center;
            img {
                width: 45px;
                height: 45px;
                // background-image: url("@/assets/avatar/1.jpg");
                // background-size: cover;
                border-radius: 50%;
                box-sizing: border-box;
                border: 2px solid #fff;
                transition: all .3s ease-in-out;
            }

            img:hover {
                transform: scale(1.1);
                box-shadow: 0 0 10px rgba(0, 0, 0, .3);
            }
        }
    }
}
</style>