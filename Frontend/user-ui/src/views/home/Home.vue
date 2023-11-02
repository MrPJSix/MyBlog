<template>
    <div class="home">
        <div class="home-left">
            <Logo></Logo>
            <MenuBar></MenuBar>
        </div>
        <div class="home-right">
            <div class="home-right-top">
                <rightHeader></rightHeader>
            </div>
            <div class="home-right-bottom">
                <div class="home-right-bottom-left">
                    <Swipe></Swipe>
                    <ContentItems></ContentItems>
                    <!-- <EndNoticeBox></EndNoticeBox> -->
                </div>
                <div class="home-right-bottom-right">
                    <Infobox></Infobox>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang='ts'>
import MenuBar from './componenets/MenuBar.vue';
import rightHeader from './componenets/rightHeader.vue';
import Swipe from './componenets/Swipe.vue';
import ContentItems from './componenets/ContentItems.vue';
import Infobox from './componenets/Infobox.vue';
import EndNoticeBox from './componenets/EndNoticeBox.vue';
import Logo from './componenets/Logo.vue';
import { getCurrentInstance, onMounted, ref } from 'vue'
const instance = getCurrentInstance();
const contentBox = ref(null as HTMLElement | null)
onMounted(() => {
    contentBox.value = document.querySelector('.home-right-bottom-left')
})
instance?.proxy?.$Bus.on('scroll-to-top', () => {
    contentBox.value!.scrollTop = 0;
})
</script>

<style lang="less">
@import url("@/assets/less/home.less");
@import url("@/assets/less/common.less");
</style>