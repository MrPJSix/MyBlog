<script setup>

</script>

<template>
  <template>
    <div class="header">
      <div class="header-content">
        <div class="title">{{ title }}</div>
        <img class="avatar" :src="avatarUrl" alt="Avatar">
        <button class="settings-button">设置</button>
      </div>
    </div>
  </template>

  <script>
    export default {
      props: {
        title: String, // 传入标题
        avatarUrl: String, // 传入头像图片的URL
      },
    };
  </script>

  <style scoped>
    .header {
      background-color: #333;
      color: #fff;
      padding: 10px;
    }

    .header-content {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }

    .title {
      font-size: 18px;
      font-weight: bold;
    }

    .avatar {
      width: 40px;
      height: 40px;
      border-radius: 50%;
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
</template>

<style scoped>

</style>