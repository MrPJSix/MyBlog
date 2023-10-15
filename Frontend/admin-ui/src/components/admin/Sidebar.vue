<template>
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
      <el-icon size="25px"><ElementPlus /></el-icon>
      <template #title>PBlog-Admin</template>
    </el-menu-item>
  </el-menu>
  <el-menu
      default-active="/admin/index"
      background-color="#545c64"
      text-color="#ffffff"
      @open="handleOpen"
      @close="handleClose"
      class="menu"
      router="true"
  >
    <el-menu-item index="/admin/index">
      <el-icon><Promotion /></el-icon>
      <template #title>首页</template>
    </el-menu-item>
    <el-menu-item index="/admin/user">
      <el-icon><User /></el-icon>
      <template #title>用户管理</template>
    </el-menu-item>
    <el-menu-item index="3">
      <el-icon><Memo /></el-icon>
      <template #title>分类管理</template>
    </el-menu-item>
    <el-menu-item index="4">
      <el-icon><Reading /></el-icon>
      <template #title>文章管理</template>
    </el-menu-item>
    <el-menu-item index="5">
      <el-icon><ChatLineRound /></el-icon>
      <template #title>评论管理</template>
    </el-menu-item>
  </el-menu>
</template>

<script setup>
  import {
    Promotion,
    User,
    Memo,
    Reading,
    ChatLineRound,
    ElementPlus,
  } from '@element-plus/icons-vue'

  const handleOpen = (key, keyPath) => {
    console.log(key, keyPath)
  }
  const handleClose = (key, keyPath) => {
    console.log(key, keyPath)
  }

</script>

<style scoped>
  .menu-title{
    height: 10vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
  .title-font{
    font-size: 15px;
    color: #c5e0e5;
    font-family: 微软雅黑;
    font-weight: bold;
  }
  .menu {
    min-height: 90vh;
  }
</style>
