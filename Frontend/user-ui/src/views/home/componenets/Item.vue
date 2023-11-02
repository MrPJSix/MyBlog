<template>
    <div class="home-content-item">
        <img class="home-content-item-img" src="../../../assets//images/3.png" alt="">
        <div class="home-content-item-info">
            <div class="home-content-item-info--title">
                <h2 class="single-line-ellipsis">{{ article.title }}</h2>
                <p>
                    <el-icon style="margin-right: 5px;" color="gray">
                        <EditPen />
                    </el-icon>
                    <span>{{ article.author.full_name }}</span>
                    <span>&nbsp;&nbsp;</span>
                    <el-icon style="margin-right: 5px;" color="gray">
                        <Clock />
                    </el-icon>
                    <span>{{ timeDisplay(article.created_at) }}</span>
                </p>
            </div>
            <p class="home-content-item-info--main">
                {{ article.content }}
            </p>
            <div class="home-content-item-info--other">
                <el-icon style="margin-right: 5px;" color="gray">
                    <Star />
                </el-icon>
                <span>{{ article.likes }}</span>
                <span>&nbsp;&nbsp;</span>
                <el-icon style="margin-right: 5px;" color="gray">
                    <View />
                </el-icon>
                <span>{{ article.read_count }}</span>
                <span>&nbsp;&nbsp;</span>
                <el-icon style="margin-right: 5px;" color="gray">
                    <ChatLineSquare />
                </el-icon>
                <span>{{ article.comment_count }}</span>
            </div>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { ArticleReciveData } from '@/types/index'
import { timeDisplay } from '@/utils/time'
import { toRef, watch, onUpdated, PropType } from 'vue'
const props = defineProps({
    article: {
        type: Object as PropType<ArticleReciveData>,
        required: true
    },
})

// const newArticle = toRef(props, 'article')

// watch(props.article, (val) => {
//     console.log(val);
//     newArticle.value = val
// })

</script>

<style scoped></style>