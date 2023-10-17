<template>
  <el-row>
    <el-col :span="2"></el-col>
    <el-col :span="20">
      <div class="time-content">
        <span class="time-title">当前时间：</span>
        <span style="margin-right: 10px" class="time-numfont">{{ currentTime.year }}-{{ currentTime.month}}-{{ currentTime.day }}</span>
        <span class="time-numfont">{{currentTime.hours}}:{{currentTime.minutes}}:{{currentTime.seconds}}</span>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="8">
      <div class="statistic-card">
        <div class="statistic-title">总用户量</div>
        <div class="statistics">
          <span class="user-stats">{{ usersCount }}</span>
          <el-icon size="20px"><User /></el-icon>
        </div>
      </div>
    </el-col>
    <el-col :span="8">
      <div class="statistic-card">
        <div class="statistic-title">总文章数</div>
        <div class="statistics">
          <span class="article-stats">{{ articlesCount }}</span>
          <el-icon size="20px"><Reading /></el-icon>
        </div>
      </div>
    </el-col>
    <el-col :span="8">
      <div class="statistic-card">
        <div class="statistic-title">总评论数</div>
        <div class="statistics">
          <span class="comment-stats">{{ commentsCount }}</span>
          <el-icon size="20px"><ChatLineRound /></el-icon>
        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script setup>
  import {
    User,
    Reading,
    ChatLineRound,
  } from "@element-plus/icons-vue";
  import { ref, onMounted, onUnmounted } from 'vue';
  import { getCurrentTime } from "../../utils/current-time.js";
  import axios from "axios";

  /* --------- 时间显示 Start --------- */
  const currentTime = ref(getCurrentTime());
  // 在组件挂载时开始计时器
  onMounted(() => {
    startTimer();
  });
  // 在组件卸载时停止计时器
  onUnmounted(() => {
    stopTimer();
  });
  // 定义计时器和更新时间的函数
  let timer;
  function startTimer() {
    timer = setInterval(() => {
      currentTime.value = getCurrentTime();
    }, 1000);
  }
  function stopTimer() {
    clearInterval(timer);
  }
  /* --------- 时间显示 End --------- */

  // 在请求拦截器中设置 Authorization 头部
  axios.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  });

  const usersCount = ref(0);
  const articlesCount = ref(0);
  const commentsCount = ref(0);

  onMounted(async () => {
    fetchData();
  });

  async function fetchData() {
    try {
      usersCount.value = await getUsersCount();
      articlesCount.value = await getArticlesCount();
      commentsCount.value = await getCommentsCount();
    } catch (error) {
      console.error(error);
    }
  }

  async function getUsersCount() {
    try {
      const response = await axios.get('http://124.220.25.230:9000/admin/users/count');
      return response.data.data;
    } catch (error) {
      throw error;
    }
  }

  async function getArticlesCount() {
    try {
      const response = await axios.get('http://124.220.25.230:9000/admin/articles/count');
      return response.data.data;
    } catch (error) {
      throw error;
    }
  }

  async function getCommentsCount() {
    try {
      const response = await axios.get('http://124.220.25.230:9000/admin/comments/count');
      return response.data.data;
    } catch (error) {
      throw error;
    }
  }
  const dialogTableVisible = ref(false)
  const gridData = [
    {
      date: '2016-05-02',
      name: 'John Smith',
      address: 'No.1518,  Jinshajiang Road, Putuo District',
    },
    {
      date: '2016-05-04',
      name: 'John Smith',
      address: 'No.1518,  Jinshajiang Road, Putuo District',
    },
    {
      date: '2016-05-01',
      name: 'John Smith',
      address: 'No.1518,  Jinshajiang Road, Putuo District',
    },
    {
      date: '2016-05-03',
      name: 'John Smith',
      address: 'No.1518,  Jinshajiang Road, Putuo District',
    },
  ]
</script>

<style scoped>
@import "../../assets/css/index.css";
</style>
