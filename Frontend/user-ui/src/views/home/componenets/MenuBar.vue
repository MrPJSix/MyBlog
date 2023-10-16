<template>
    <div class="home-menu-bar">
        <!-- <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 24 24">
            <g fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M14 10h1c1 0 1 1 2.016 3.527C18 16 18 17 19 17h1"></path>
                <path d="M13 17c1.5 0 3-2 4-3.5s2.5-3.5 4-3.5"></path>
                <path d="M3 19c0 1.5.5 2 2 2s2-4 3-9s1.5-9 3-9s2 .5 2 2"></path>
                <path d="M5 12h6"></path>
            </g>
        </svg> -->
        <div class="home-menu-bar-item list-item--link" v-for="(item, index) in items" :key="index">
            <div class="home-menu-bar-item-prepend">
                <i class="home-menu-bar-item-icon">
                    <SvgIcon :iconName="item.icon"></SvgIcon>
                </i>
                <div class="home-menu-bar-item__spacer"></div>
            </div>
            <div class="home-menu-bar-item-content">
                <div class="home-menu-bar-item-title">{{ item.title }}</div>
            </div>
        </div>

        <!-- <div class="home-menu-bar-item list-item--link">
            <div class="home-menu-bar-item-prepend">
                <i class="home-menu-bar-item-icon">
                    <SvgIcon iconName="icon-jisuanjikaifa"></SvgIcon>
                </i>
                <div class="home-menu-bar-item__spacer"></div>
            </div>
            <div class="home-menu-bar-item-content">
                <div class="home-menu-bar-item-title">计算机</div>
            </div>
        </div> -->
    </div>
</template>

<script setup lang='ts'>
import { reactive } from 'vue';
const items: MenuItem[] = reactive([
    {
        title: '计算机',
        icon: 'icon-jisuanjikaifa'
    },
    {
        title: '数学',
        icon: 'icon-shuxue'
    },
    {
        title: '英语',
        icon: 'icon-yingyu'
    },
    {
        title: '物理',
        icon: 'icon-physics-world'
    },
    {
        title: '化学',
        icon: 'icon-shiyanhuaxue'
    },
    {
        title: '生物',
        icon: 'icon-shengwu'
    },
    {
        title: '政治',
        icon: 'icon-zhengfu'
    },
    {
        title: '历史',
        icon: 'icon-king'
    },
    {
        title: '地理',
        icon: 'icon-diqiu'
    },
    {
        title: '其他',
        icon: 'icon-qita'
    }
])

</script>

<style scoped lang="less"></style>
