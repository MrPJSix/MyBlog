<script setup>
import Sidebar from '../components/Sidebar.vue';
import PageHeader from "../components/PageHeader.vue";
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
      <el-main>
        <router-view/>
      </el-main>
    </el-container>

  </el-container>
</template>

<style scoped>
.admin {
  display: flex;
}
</style>