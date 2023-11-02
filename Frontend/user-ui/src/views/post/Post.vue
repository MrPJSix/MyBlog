<template>
    <div class="post-article">
        <h1>发布文章</h1>
        <div class="category-selector">
            <ElSelect v-model="selecedPrimary" placeholder="选择分类" @change="getSecondary" class="select-box">
                <ElOption v-for="(c, i) in primary" 
                    :key=i 
                    :value="c.id"
                    :label="c.title"
                ></ElOption>
            </ElSelect>
            <ElSelect v-model="selecedSecondary" placeholder="选择分类" @change="getId" class="select-box">
                <ElOption v-for="(c, i) in secondary" 
                    :key=i 
                    :value="c.id"
                    :label="c.title"
                ></ElOption>
            </ElSelect>
            <ElButton type="primary" @click="switchMode">{{ editorMode === 'html' ? '点击切换至Markdown格式' : '点击切换至Html模式' }}</ElButton>
            <ElButton class="post-btn" type="primary" @click="post">POST</ElButton>
        </div>

        <ElInput 
            v-model="title" 
            maxlength="30" 
            placeholder="请输入标题" 
            show-word-limit 
            type="textarea"
            rows="2"
            resize="none"
            class="title-input"
        ></ElInput>
        <ElDivider></ElDivider>
        <HtmlEditor v-if="editorMode === 'html'"></HtmlEditor>
        <MdEditor v-else></MdEditor>
    </div>
</template>

<script setup lang='ts'>

import { ElSelect, ElOption, ElInput, ElDivider, ElButton } from 'element-plus'
import { ref, getCurrentInstance } from 'vue'
import api from '@/api/index'
import HtmlEditor from '@/components/post/HtmlEditor.vue';
import MdEditor from '@/components/post/MdEditor.vue';

const instance = getCurrentInstance()

const editorMode = ref('html')

const title = ref('')

const primary = ref<{ id: number, title: string }[]>([])
const secondary = ref<{ id: number, title: string }[]>([])

const selecedPrimary = ref('')
const selecedSecondary = ref('')

const selectedCategory = ref<{primaryId:number,secondaryId:number}>({
    primaryId: 0,
    secondaryId: 0
})

api.category.getPrimary().then(res => {
    res.data.forEach((item: any) => {
        primary.value.push({
            id: item.id,
            title: item.name
        })
    })
})

const getSecondary = (value:number) => {
    secondary.value = []
    selecedSecondary.value = ''
    selectedCategory.value.primaryId = value
    api.category.getSecondary(value).then(res => {
        res.data.forEach((item: any) => {
            secondary.value.push({
                id: item.id,
                title: item.name
            })
        })
    })
}

const getId = (value:number) => {
    selectedCategory.value.secondaryId = value
}


const post = () => {
    instance?.proxy?.$Bus.emit('post', {
        title: title.value,
        category_id: selectedCategory.value.secondaryId,
    })
}

const switchMode = () => {
    editorMode.value = editorMode.value === 'html' ? 'markdown' : 'html'
}

</script>

<style scoped lang="less">
.post-article {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 1200px;
    // min-height: calc(100vh - 80px);
    margin: 10px auto;
    margin-bottom: 0;
    box-sizing: border-box;
    box-shadow: 0 0 10px rgba(0, 0, 0, .1);
    padding: 20px 50px;
    h1 {
        height: 45px;
    }
}
.category-selector {
    margin: 20px 0;
    box-sizing: border-box;
    width: 100%;
    position: relative;
    .select-box {
        width: 200px;
        margin-right: 20px;
    }
    .post-btn {
        position: absolute;
        right: 0;
    }
}

.title-input {
    width: 100%;
}
</style>