<template>
    <button class="h-btn h-btn--icon text-middle-emphasis">
        <span class="h-btn__overlay"></span>
        <span class="h_btn__content">
            <i class="h-icon">
                <slot></slot>
            </i>
        </span>
    </button>
</template>

<script setup lang='ts'>
</script>

<style scoped lang="less">
@import url('@/assets/less/common.less');
</style>