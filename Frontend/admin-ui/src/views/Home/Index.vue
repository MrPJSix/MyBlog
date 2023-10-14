<template>
    <div class="main">
        <SideBar />
        <div class="main-right">
            <TopBar />
            <div class="main-right-box">
                <div class="main-right-left">
                <Swipe />
                <MainContent />
            </div>
            <div class="main-right-right">
                <Info />
            </div>
            </div>
            
        </div>
    </div>
</template>

<script setup name="Main">
import SideBar from './SideBar.vue';
import TopBar from './TopBar.vue';
import Swipe from './Swipe.vue';
import MainContent from './Main/Index.vue'
import Info from './Info.vue';
</script>

<style scoped>
.main {
    height: 100%;
    width: 100%;
    display: flex;
}
.main-right {
    height: 100%;
    width: 80%;
    display: flex;
    flex-direction: column;
}

.main-right-box{
    display: flex;
    height: calc(100% - 50px);
}
.main-right-left{
    height: 100%;
    width: 80%;
    border: 1px solid black;
    overflow: auto;
}

::-webkit-scrollbar {
    display: none;
}

.main-right-right {
    height: 100%;
    width: 20%;
    border: 1px solid black;
    display: flex;
    flex-direction: column;
    justify-content: center;
}
</style>