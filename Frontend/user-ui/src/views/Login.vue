<template>
  <div>
    <div class="login-page">
      <div class="container">
        <div class="form-box" :style="{ transform: isRegister ? 'translateX(82%)' : 'translateX(0)' }">
          <div class="register-box"
               :class="{ hidden: !isRegister }"
          >
            <h1>register</h1>
            <input type="text" placeholder="用户名" v-model="r_username" title="用户名只能是字母和数字，长度为8-25"/>
            <input type="text" placeholder="昵称" v-model="r_full_name"/>
            <input type="password" placeholder="密码" v-model="r_password" title="密码只能是字母、数字和特殊字符，长度为8-25"/>
            <input type="password" placeholder="确认密码" v-model="r_confirm_password"/>
            <button @click="register">注册</button>
          </div>
          <div class="login-box"
               :class="{ hidden: isRegister }"
          >
            <h1>login</h1>
            <input type="text" placeholder="用户名" v-model="l_username"/>
            <input type="password" placeholder="密码" v-model="l_password"/>
            <button @click="login">登录</button>
          </div>
        </div>
        <div class="con-box left">
          <h2>欢迎来到<span>PanBlog</span></h2>
          <p style="font-size: 13px">快来和小伙伴分享<span style="font-weight: bolder">idea</span>吧</p>
          <img src="../assets/images/PanBlog_Logo2.png" alt="" />
          <div
              style="display: flex; justify-content: center; align-items: center; width: 200px;"
          >
            <span style="color: #7a7878">已有账号?</span>
            <span
                id="login"
                style="font-weight: bold; color: pink; cursor: pointer"
                @click="toggleForm"
            >登录</span>
          </div>
        </div>
        <div class="con-box right">
          <h2>欢迎来到<span>PanBlog</span></h2>
          <p style="font-size: 13px">快来和小伙伴分享<span style="font-weight: bolder">idea</span>吧</p>
          <img src="../assets/images/PanBlog_Logo1.png" alt="" />
          <div
              style="display: flex; justify-content: center; align-items: center; width: 200px;"
          >
            <span style="color: #7a7878">没有账号?</span>
            <span
                id="register"
                style="font-weight: bold; color: pink; cursor: pointer"
                @click="toggleForm"
            >注册</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref} from "vue";
// import axios from "axios";
import api from '@/api/index' 
import { ElMessage } from "element-plus";
import { LoginForm, RegisterForm } from "@/interface/index";
import {useRouter} from "vue-router";

const router = useRouter()

/* ---------- 登录框控制-Start ---------- */
const isRegister = ref(false);
const toggleForm = () => {
  isRegister.value = !isRegister.value;
};
/* ---------- 登录框控制-End ---------- */

/* ---------- 注册-Start ---------- */
const r_username = ref('');
const r_full_name = ref('');
const r_password = ref('');
const r_confirm_password = ref('');

// const register = async () => {
//   try {
//     const response = await axios.post(`${BaseURL}/register`, {
//       username: r_username.value,
//       password: r_password.value,
//       confirm_password: r_confirm_password.value,
//       full_name: r_full_name.value
//     });
//     if (response.data.status === 200) {
//       ElMessage({
//         message: '注册成功',
//         type: 'success',
//       });
//       l_username.value = r_username.value;
//       l_password.value = r_password.value;
//     } else {
//       ElMessage({
//         message: response.data.message,
//         type: 'error',
//       });
//     }
//   } catch (error) {
//     ElMessage({
//       message: 'An error occurred while trying to log in.',
//       type: 'warning',
//     });
//     console.error('An error occurred:', error);
//   }
// }

const register = () => {
  const registerForm:RegisterForm = {
    username: r_username.value,
    password: r_password.value,
    confirm_password: r_confirm_password.value,
    full_name: r_full_name.value
  }
  api.user.register(registerForm).then(res => {
    ElMessage({
      message: '注册成功',
      type: 'success',
    })
    toggleForm();
    l_username.value = r_username.value;
    l_password.value = r_password.value;
  }).catch(err => {
    ElMessage({
      message: err.data.message,
      type: 'error',
    })
    console.error('An error occurred:', err)
  })
}

/* ---------- 注册-End ---------- */

/* ---------- 登录-Start ---------- */
const l_username = ref('');
const l_password = ref('');
const login = () => {
  const loginForm: LoginForm = {
    username: l_username.value,
    password: l_password.value
  }
  api.user.login(loginForm).then(res => {
    localStorage.setItem('token', res.token);
    ElMessage({
      message: '登录成功',
      type: 'success',
    })
    router.push('/home')
  }).catch(err => {
    ElMessage({
      message: err,
      type: 'error',
    })
    console.error('An error occurred:', err)
  })
}
// const login = async () => {
//   try {
//     const response = await axios.post(`${BaseURL}/login`, {
//       username: l_username.value,
//       password: l_password.value
//     });
//     if (response.data.status === 200) {
//       ElMessage({
//         message: '登录成功',
//         type: 'success',
//       });
//       localStorage.setItem('token', response.data.token);
//       // router.push("/admin/index")
//     } else {
//       ElMessage({
//         message: response.data.message,
//         type: 'error',
//       });
//     }
//   } catch (error) {
//     ElMessage({
//       message: 'An error occurred while trying to log in.',
//       type: 'warning',
//     });
//     console.error('An error occurred:', error);
//   }
// };
/* ---------- 登录-End ---------- */

</script>

<style scoped>
  @import "../assets/css/login.css";
</style>