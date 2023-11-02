<template>
    <div class="a-body" v-html="article_content"></div>
    <hr>
</template>

<script setup lang='ts'>
import { ref, watch, onMounted, getCurrentInstance } from 'vue'
import { marked } from 'marked'
const props = defineProps<{
    content: string
    content_type: string
}>()

const instance = getCurrentInstance()

const article_content = ref('')

const articleBody = ref(null as null | HTMLDivElement)

onMounted(() => {
    articleBody.value = document.querySelector('.a-body')
})

window.addEventListener('scroll',() => {
    console.log(articleBody.value?.getBoundingClientRect().top);
    if(articleBody.value && articleBody.value.getBoundingClientRect().top < -100) {
        console.log(2222);
        instance?.proxy?.$Bus.emit('showToTop')
        // instance?.proxy?.$Bus.off('showToTop')
    }
    if(articleBody.value && articleBody.value.getBoundingClientRect().top > 0) {
        instance?.proxy?.$Bus.emit('hideToTop')
        instance?.proxy?.$Bus.off('hideToTop')
    }
})

watch(()=>props.content + props.content_type, () => {
    if(props.content_type === 'm') {
        article_content.value = marked(props.content)
    } else {
        article_content.value = props.content
    }
})


</script>

<style scoped></style>