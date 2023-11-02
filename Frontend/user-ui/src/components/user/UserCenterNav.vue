<template>
    <div class="user-center-nav">
        <div class="nav-item" v-for="(item, index) in nav_items" :key="index">
            <a :href="item.path">
                <div class="nav-item-span"></div>
                <div class="nav-item-con">
                    <span>{{ item.title }}</span>
                </div>
            </a>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
const base_path = '/usercenter';
const nav_items = reactive([
    {
        title: '个人中心',
        path: base_path,
    },
    {
        title: '修改资料',
        path: base_path + '/config',
    },
    {
        title: '我的主题',
        path: base_path + '/mytopics',
    },
    {
        title: '我的回复',
        path: base_path + '/myreplies',
    },
    {
        title: '我的收藏',
        path: base_path + '/myfavorites',
    },
    {
        title: '关注用户',
        path: base_path + '/myfollowings',
    },
    {
        title: '我的粉丝',
        path: base_path + '/myfans',
    }
])
</script>

<style scoped></style>