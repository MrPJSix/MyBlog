<template>
    <MdEditor class="md-edior" v-model="text" :showCodeRowNumber="true">
        <template #defToolbars>
            <Emoji>
                <template #trigger>Emoji</template>
            </Emoji>
        </template>
    </MdEditor>
</template>

<script setup lang='ts'>
import { ref, onMounted, getCurrentInstance, onBeforeUnmount } from 'vue'
import { MdEditor } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css';

import { Emoji } from '@vavt/v3-extension'

import '@vavt/v3-extension/lib/asset/Emoji.css'
import api from '@/api';
import { message } from '@/utils/client';

const text = ref('');
// const toolbars = ['bold', 0, 'underline']

const instance = getCurrentInstance()

onMounted(() => {
    const editor = document.querySelector('.md-edior') as HTMLElement
    if(editor == null) return
    editor.style.minHeight = `calc(100vh - ${editor.getBoundingClientRect().top}px - 20px)`
})

instance?.proxy?.$Bus.on('post',(data:any) => {
    console.log("这里是md编辑器");
    console.log(text.value);
    api.article.postArticle({
        title: data.title,
        category_id: data.category_id,
        content: text.value,
        content_type: 'm'
    }).then(res => {
        if(res.status === 200) {
            message.success('发布成功')
        }
    })
})

onBeforeUnmount(() => {
    instance?.proxy?.$Bus.off('post')
})


</script>

<style scoped>
.md-edior{
    border: 1px solid #ccc;
    box-sizing: border-box;
}
</style>