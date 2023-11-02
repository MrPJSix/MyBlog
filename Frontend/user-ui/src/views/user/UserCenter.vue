<template>
    <div class="user-center">
        <div class="user-center-content">
            <UserCenterHead></UserCenterHead>
            <UserCenterBody></UserCenterBody>
        </div>
    </div>
</template>

<script setup lang='ts'>
import UserCenterHead from '@/components/user/UserCenterHead.vue'
import UserCenterBody from '@/components/user/UserCenterBody.vue'
</script>

<style lang="less">
@import url('@/assets/less/user-center.less');
</style>