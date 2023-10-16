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
                    <!-- <Swipe></Swipe> -->
                    <subMenuBar></subMenuBar>
                    <ContentItems></ContentItems>
                    <EndNoticeBox></EndNoticeBox>
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
import subMenuBar from './componenets/subMenuBar.vue';
import ContentItems from './componenets/ContentItems.vue';
import Infobox from './componenets/Infobox.vue';
import EndNoticeBox from './componenets/EndNoticeBox.vue';
import Logo from './componenets/Logo.vue';
</script>

<style lang="less">
@import url("@/assets/less/home.less");

.home {
    display: flex;
    // background-color: aquamarine;
    width: 100%;
    height: 100vh;

    .home-left {
        width: @home-left-width;
        // background-color: red;
        border-right: 2px solid rgba(0, 0, 0, 0.2);
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }

    .home-right {
        width: 100%-@home-left-width;
        // background-color: rgb(12, 12, 163);
        display: flex;
        flex-direction: column;

        .home-right-top {
            height: @home-right-top-height;
            border-bottom: 2px solid rgba(0, 0, 0, 0.2);
            // background-color: yellow;
        }

        .home-right-bottom {
            height: 100%-@home-right-top-height;
            // background-color: green;
            display: flex;

            ::-webkit-scrollbar {
                display: none;
            }

            .home-right-bottom-left {
                width: @home-right-bottom-left-width;
                // background-color: pink;
                overflow: auto;
                border-right: 2px solid rgba(0, 0, 0, 0.2);
            }

            .home-right-bottom-right {
                width: 100%-@home-right-bottom-left-width;
                // background-color: purple;
                display: flex;
                flex-direction: column;
                justify-content: space-around;
            }
        }
    }
}
</style>