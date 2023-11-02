<template>
    <div class="home-swipe">
        <n-carousel draggable>
            <img class="carousel-img" src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel1.jpeg">
            <img class="carousel-img" src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel2.jpeg">
            <img class="carousel-img" src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel3.jpeg">
            <img class="carousel-img" src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel4.jpeg">
        </n-carousel>
    </div>
</template>

<script setup lang='ts'>
import { NCarousel } from 'naive-ui'
</script>

<style scoped>
.carousel-img {
    width: 100%;
    height: 240px;
    object-fit: cover;
}
</style>