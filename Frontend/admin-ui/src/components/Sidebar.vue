<template>
  <div class="sidebar">
    <el-menu
        background-color="#0A3946FF"
        text-color="#ffffff"
        active-text-color="#ffffff"
        :collapse="isCollapse"
        @open="handleOpen"
        @close="handleClose"
        class="menu-title"
    >
      <el-menu-item index="0" class="title-font">
        <el-icon size="25px"><element-plus /></el-icon>
        <template #title>博客后台管理</template>
      </el-menu-item>
    </el-menu>
    <el-menu
        default-active="/admin"
        background-color="#545c64"
        text-color="#fff"
        :collapse="isCollapse"
        @open="handleOpen"
        @close="handleClose"
        class="menu"
        router="true"
    >
      <el-menu-item index="1">
        <el-icon><location /></el-icon>
        <template #title>首页</template>
      </el-menu-item>
      <el-menu-item index="/index2">
        <el-icon><icon-menu /></el-icon>
        <template #title>Navigator Two</template>
      </el-menu-item>
      <el-menu-item index="3">
        <el-icon><document /></el-icon>
        <template #title>Navigator Three</template>
      </el-menu-item>
      <el-menu-item index="4">
        <el-icon><setting /></el-icon>
        <template #title>Navigator Four</template>
      </el-menu-item>
      <el-button @click="toggleCollapse" class="collapse-toggle-button"  link text>
        <el-icon v-if="isCollapse" size="15px" color="#fff">
          <CaretRight />
        </el-icon>
        <el-icon v-else size="15px" color="#fff">
          <CaretLeft />
        </el-icon>
      </el-button>
    </el-menu>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import {
    Document,
    Menu as IconMenu,
    Location,
    Setting,
    CaretLeft,
    CaretRight,
    ElementPlus,
  } from '@element-plus/icons-vue'

  const isCollapse = ref(true)
  const handleOpen = (key, keyPath) => {
    console.log(key, keyPath)
  }
  const handleClose = (key, keyPath) => {
    console.log(key, keyPath)
  }
  const toggleCollapse = () => {
    isCollapse.value = !isCollapse.value
  }
</script>

<style scoped>
  .sidebar {
    //position: relative;
    //width: 200px;  /* Or your desired width */
    //min-height: 100%;
  }
  .menu-title{
    height: 15vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
  .title-font{
    font-size: 20px;
    color: #DAA520FF;
    font-family: 微软雅黑;
    font-weight: bold;
  }
  .menu {
    min-height: 85vh;
  }
  .collapse-toggle-button {
    position: absolute;
    top: 50%;
    right: 2px;  /* Adjust according to your needs */
    transform: translateY(-50%);
    width: 15px;  /* Or your desired width */
    height: 15px; /* Or your desired height */
  }
</style>
