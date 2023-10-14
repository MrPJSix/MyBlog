<template>
    <div class="main-content">
        <div class="main-content-bar">
            <div class="main-content-bar-item">全部</div>
            <div class="main-content-bar-item">Python</div>
            <div class="main-content-bar-item">Golang</div>
            <div class="main-content-bar-item">Vue</div>
            <div class="main-content-bar-item">计算机网络</div>
            <div class="main-content-bar-item">操作系统</div>
            <div class="main-content-bar-item">更多</div>
        </div>
        <div class="main-content-items">
            <Item></Item>
            <Item></Item>
            <Item></Item>
        </div>
        <BottomBox></BottomBox>
    </div>
    
</template>

<script setup name="MainContent">
import Item from './Item.vue';
import BottomBox from './BottomBox.vue';
</script>

<style scoped>
.main-content {
    width: 850px;
    /* border: 1px solid black; */
    margin: 30px auto;
    margin-bottom: 5px;
}
.main-content-bar {
    width: 100%;
    height: 50px;
    border-bottom: 2px solid black;
    font: 15px 'sans-serif';
    line-height: 50px;
    display: flex;
    justify-content: space-around;
    margin-bottom: 50px;
}
.main-content-items {
    width: 100%;
    display: flex;
    flex-direction: column;
}
</style>