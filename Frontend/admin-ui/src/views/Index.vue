<script>
export default {
  name: "Index",
  components: {},
};
</script>

<template>
  <div class="message-test">
    <h1>Message Test Page</h1>
    <el-button @click="showMessage">Show Message</el-button>
  </div>
</template>

<script setup>
import { ElMessage, ElButton } from "element-plus";

const showMessage = () => {
  ElMessage({
    message: 'Hello, this is a message.',
    type: 'success',
  });
};
</script>

<style scoped>
.message-test {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
}
{
  background-color: #f0fff0
;
}
</style>
