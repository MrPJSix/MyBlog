<template>
    <div class="main">
        <Article></Article>
        <Comments></Comments>
    </div>
</template>

<script setup lang='ts'>
import Article from '@/views/article/Article.vue'
import Comments from '@/views/Comments/Comments.vue';
</script>

<style lang="less">
@import url('@/assets/less/main.less');

</style>