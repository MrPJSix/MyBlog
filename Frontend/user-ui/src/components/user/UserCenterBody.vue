<template>
    <div class="user-center-body">
        <UserCenterNav></UserCenterNav>
        <div class="user-center-router">
            <router-view></router-view>
        </div>
    </div>
</template>

<script setup lang='ts'>
import UserCenterNav from '@/components/user/UserCenterNav.vue';
</script>

<style scoped>
</style>