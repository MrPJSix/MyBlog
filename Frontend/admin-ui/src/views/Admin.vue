<script setup>
import Sidebar from '../components/Sidebar.vue';
</script>

<template>
  <div class="admin">
    <div class="sidebar-page">
      <Sidebar></Sidebar>
    </div>
    <div class="main-content">
      <router-view/>
    </div>
  </div>
</template>

<style scoped>
.admin {
  display: flex;
}
/*.sidebar-page {
  width: 15%;
}*/
.main-content {
  min-width: 85%;
}
</style>