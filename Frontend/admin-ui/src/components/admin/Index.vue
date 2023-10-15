<template>
  <el-row>
    <el-col :span="2"></el-col>
    <el-col :span="20">
      <div class="time-content">
        <span>当前时间：</span>
        <span style="margin-right: 5px">{{ currentTime.year }}-{{ currentTime.month}}-{{ currentTime.day }}</span>
        <span>{{currentTime.hours}}:{{currentTime.minutes}}:{{currentTime.seconds}}</span>
      </div>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="8">
      <div class="statistic-card">
        <div class="statistic-title">总用户量</div>
        <div class="statistics">
          <span class="user-stats">10,000</span>
          <el-icon size="20px"><User /></el-icon>
        </div>
      </div>
    </el-col>
    <el-col :span="8">
      <div class="statistic-card">
        <div class="statistic-title">总文章数</div>
        <div class="statistics">
          <span class="article-stats">10,000</span>
          <el-icon size="20px"><Reading /></el-icon>
        </div>
      </div>
    </el-col>
    <el-col :span="8">
      <div class="statistic-card">
        <div class="statistic-title">总评论数</div>
        <div class="statistics">
          <span class="comment-stats">10,000</span>
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

  /* --------- 时间显示 Start --------- */
  import {getCurrentTime} from "../../utils/current-time.js";
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
</script>

<style scoped>
.time-content {
  margin: 10px;
  min-height: 36px;
}
.statistic-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  border: 1px solid lightgray;
  padding: 10px;
  margin-left: 10px;
  margin-right: 10px;
  border-radius: 10px;
}
.statistic-title{
  font-size: 22px;
  font-weight: bold;
  color: #555555;
}
.statistics {
  margin-top: 10px;
  padding: 10px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  text-align: center;
}
.statistics span {
  margin-right: 10px;
}
.user-stats {
  font-family: "Times New Roman";
  font-size: 20px;
  color: lightcoral;
}
.article-stats {
  font-family: "Times New Roman";
  font-size: 20px;
  color: dodgerblue;
}
.comment-stats {
  font-family: "Times New Roman";
  font-size: 20px;
  color: green;
}
</style>
