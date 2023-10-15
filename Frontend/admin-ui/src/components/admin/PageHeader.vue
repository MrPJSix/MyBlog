<script setup>
import { Setting } from '@element-plus/icons-vue'

</script>

<template>
  <div class="header header-content">
    <el-page-header :icon="ArrowLeft"></el-page-header>
    <div class="title">欢迎来到博客后台管理系统</div>
    <div class="avatar">
      <el-avatar :size="50" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
      <span class="avatar-title">Admin</span>
    </div>
    <el-button :icon="Setting" color="#dddddd" circle />
  </div>
</template>

<style scoped>
  .header {
    margin: 0;
    padding: 0;
    height: 100%;
    border-bottom: 2px solid #c6d0dc; /* 1px的下边框线，可以根据需要调整线条宽度和颜色 */
    border-left: none; /* 移除左边框 */
    border-right: none; /* 移除右边框 */
    border-top: none; /* 移除上边框 */
  }

  .header-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .title {
    font-size: 18px;
    font-weight: bold;
    color: #00363d;
  }

  .avatar {
    width: 10%;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .avatar-title{
    font-size: 18px;
    font-weight: bold;
    color: #00363d;
  }

  .settings-button {
    background-color: #555;
    color: #fff;
    padding: 5px 10px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
</style>