<script setup>
import Sidebar from '../components/admin/Sidebar.vue';
import PageHeader from "../components/admin/PageHeader.vue";
</script>

<template>
  <el-container class="admin">
      <el-aside width="200px">
        <Sidebar></Sidebar>
      </el-aside>
    <el-container>
      <el-header height="10vh">
        <PageHeader></PageHeader>
      </el-header>
      <el-main class="main">
        <router-view/>
      </el-main>
    </el-container>
  </el-container>
</template>

<style scoped>
.admin {
  display: flex;
}
.main {
  margin: 0;
  padding: 0;
  height: 90vh;
}
</style>