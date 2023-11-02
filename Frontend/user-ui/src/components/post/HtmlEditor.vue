<template>
    <div style="border: 1px solid #ccc" class="editor">
        <Toolbar
            style="border-bottom: 1px solid #ccc;"
            :editor="editorRef"
            :defaultConfig="toolbarConfig"
            :mode="mode"
        >
        </Toolbar>
        <Editor
            class="html-editor"
            v-model="valueHtml"
            :defaultConfig="editorConfig"
            :mode="mode"
            @onCreated="handleCreated"
        ></Editor>
    </div>
</template>

<script setup lang='ts'>
import '@wangeditor/editor/dist/css/style.css'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'

import { onBeforeUnmount, ref, shallowRef, onMounted, getCurrentInstance } from 'vue'
import api from '@/api';

const instance = getCurrentInstance()

instance?.proxy?.$Bus.on('post',(data:any) => {
    console.log("这里是html编辑器");
    
    api.article.postArticle({
        title: data.title,
        category_id: data.category_id,
        content: valueHtml.value,
        content_type: 'h'
    }).then(res => {
        if(res.status === 200) {
            instance?.proxy?.$message({
                type: 'success',
                message: '发布成功'
            })
        }
    })
})

const editorRef = shallowRef(null as any)

const valueHtml = ref('')

const mode = ref('default')

onMounted(() => {
    const editor = document.querySelector('.editor') as HTMLElement
    if(editor == null) return
    editor.style.minHeight = `calc(100vh - ${editor.getBoundingClientRect().top}px - 20px)`
})

const toolbarConfig = ref({})
const editorConfig = ref({ placeholder: '请输入内容......' })

onBeforeUnmount(() => {
    const editor = editorRef.value
    if( editor == null ) return
    editor.destroy()
    instance?.proxy?.$Bus.off('post')
})

const handleCreated = (editor:any) => {
    editorRef.value = editor
}
</script>

<style scoped>
.editor {
    box-sizing: border-box;
}
.html-editor {
    /* height: auto; */
    overflow-y: hidden;
    flex: auto;
}
</style>