<template>
  <div class="login-page">
    <div class="square">
      <ul>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
      </ul>
    </div>
    <div class="login-container">
      <div class="login-box">
        <div class="owl" :class="{password: isPasswordFocused}" ref="owlRef">
          <div class="hand"></div>
          <div class="hand hand-r"></div>
          <div class="arms">
            <div class="arm"></div>
            <div class="arm arm-r"></div>
          </div>
        </div>
        <div class="input-box">
          <input type="text" placeholder="账号" v-model="username" @keyup.enter="login">
          <input type="password" placeholder="密码" v-model="password" @focus="handleFocus" @blur="handleBlur" @keyup.enter="login">
          <button @click="login">登录</button>
        </div>
      </div>
    </div>
    <div class="circle">
      <ul>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
      </ul>
    </div>
  </div>
</template>

<script setup>
  import axios from 'axios';
  import { ref } from 'vue';
  import { ElMessage } from "element-plus";

  const username = ref('');
  const password = ref('');

  const isPasswordFocused = ref(false);
  const handleFocus = () => {
    isPasswordFocused.value = true;
  };
  const handleBlur = () => {
    isPasswordFocused.value = false;
  };
  const login = async () => {
    try {
      const response = await axios.post('http://124.220.25.230:9000/admin/login', {
        username: username.value,
        password: password.value
      });
      if (response.data.status === 200) {
        ElMessage({
          message: '登录成功',
          type: 'success',
        })
        localStorage.setItem('token', response.data.token);
      } else {
        ElMessage({
          message: response.data.message,
          type: 'error',
        })
      }
    } catch (error) {
      ElMessage({
        message: 'An error occurred while trying to log in.',
        type: 'warning',
      })
      console.error('An error occurred:', error);
    }
  };
</script>

<style scoped>
@import "../assets/css/login.css";
</style>
