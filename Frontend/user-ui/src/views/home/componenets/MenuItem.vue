<template>
    <div class="home-menu-bar-item list-item--link">
        <div class="home-menu-bar-item-prepend">
            <i class="home-menu-bar-item-icon">
                <SvgIcon :iconName="item.icon" :activated="item.activated"></SvgIcon>
            </i>
            <div class="home-menu-bar-item__spacer"></div>
        </div>
        <div class="home-menu-bar-item-content">
            <div class="home-menu-bar-item-title" :class="content">{{ item.title }}</div>
        </div>
        <span class="home-menu-bar-item__overlay" :class="overlay"></span>
    </div>
</template>

<script setup lang='ts'>
import SvgIcon from '@/components/basic/SvgIcon.vue';
import { computed } from 'vue';
const props = defineProps({
    item: {
        type: Object,
        required: true
    }
})
const content = computed(() => {
    if (props.item.activated) {
        return 'content__active'
    }
    return ''
})
const overlay = computed(() => {
    if (props.item.activated) {
        return 'overlay__active'
    }
    return ''
})
</script>

<style scoped></style>