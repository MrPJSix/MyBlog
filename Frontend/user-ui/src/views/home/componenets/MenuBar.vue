<template>
    <div class="home-menu-bar">
        <MenuItem v-for="(item, index) in CategoriesRef" :key="index" :item="item" @click.stop="activate(index)">
        </MenuItem>
    </div>
</template>

<script setup lang='ts'>
import { Category, SubCategory } from '@/interface';
import MenuItem from './MenuItem.vue';
import { ref, getCurrentInstance } from 'vue';
import api from '@/api/index';
// const items = reactive([
//     {
//         title: '计算机',
//         icon: 'icon-jisuanjikaifa',
//         activated: true,
//     },
//     {
//         title: '数学',
//         icon: 'icon-shuxue',
//         activated: false,
//     },
//     {
//         title: '英语',
//         icon: 'icon-yingyu',
//         activated: false,
//     },
//     {
//         title: '物理',
//         icon: 'icon-physics-world',
//         activated: false,
//     },
//     {
//         title: '化学',
//         icon: 'icon-shiyanhuaxue',
//         activated: false,
//     },
//     {
//         title: '生物',
//         icon: 'icon-shengwu',
//         activated: false,
//     },
//     {
//         title: '政治',
//         icon: 'icon-zhengfu',
//         activated: false,
//     },
//     {
//         title: '历史',
//         icon: 'icon-king',
//         activated: false,
//     },
//     {
//         title: '地理',
//         icon: 'icon-diqiu',
//         activated: false,
//     },
//     {
//         title: '其他',
//         icon: 'icon-qita',
//         activated: false,
//     }
// ])

let lastActive = 0;

const activate = (index: number) => {
    CategoriesRef.value[lastActive].activated = false;
    CategoriesRef.value[index].activated = true;
    lastActive = index;
    emit1()
}

const CategoriesRef = ref<Category[]>([]);

const instance = getCurrentInstance()

api.category.getPrimary().then(res => {
    CategoriesRef.value.push({
        id: 0,
        title: '查看全部',
        icon: 'icon-quanbu',
        activated: false
    })
    res.data.forEach( (item:any) => {
        CategoriesRef.value.push({
            id: item.id,
            title: item.name,
            icon: item.icon,
            activated: false
        })
        CategoriesRef.value[0].activated = true;
    })
    emit1()
})

const emit1 = () => {
    instance?.proxy?.$Bus.emit('change-category', CategoriesRef.value[lastActive].id as number)
}


</script>

<style scoped lang="less"></style>
