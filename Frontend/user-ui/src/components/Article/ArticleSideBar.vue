<template>
    <div class="right-side-bar">
        <div class="side-toolbar">
            <div class="toolbar-item" @click="like_change">
                <svg class="icon icon-like" v-if="!like" viewBox="0 0 1024 1024">
                    <path
                        d="M885.9 533.7c16.8-22.2 26.1-49.4 26.1-77.7c0-44.9-25.1-87.4-65.5-111.1a67.67 67.67 0 0 0-34.3-9.3H572.4l6-122.9c1.4-29.7-9.1-57.9-29.5-79.4A106.62 106.62 0 0 0 471 99.9c-52 0-98 35-111.8 85.1l-85.9 311H144c-17.7 0-32 14.3-32 32v364c0 17.7 14.3 32 32 32h601.3c9.2 0 18.2-1.8 26.5-5.4c47.6-20.3 78.3-66.8 78.3-118.4c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c-.2-12.6-2-25.1-5.6-37.1zM184 852V568h81v284h-81zm636.4-353l-21.9 19l13.9 25.4a56.2 56.2 0 0 1 6.9 27.3c0 16.5-7.2 32.2-19.6 43l-21.9 19l13.9 25.4a56.2 56.2 0 0 1 6.9 27.3c0 16.5-7.2 32.2-19.6 43l-21.9 19l13.9 25.4a56.2 56.2 0 0 1 6.9 27.3c0 22.4-13.2 42.6-33.6 51.8H329V564.8l99.5-360.5a44.1 44.1 0 0 1 42.2-32.3c7.6 0 15.1 2.2 21.1 6.7c9.9 7.4 15.2 18.6 14.6 30.5l-9.6 198.4h314.4C829 418.5 840 436.9 840 456c0 16.5-7.2 32.1-19.6 43z">
                    </path>
                </svg>
                <svg class="icon" fill="#32a0d5" v-else viewBox="0 0 1024 1024">
                    <path
                        d="M885.9 533.7c16.8-22.2 26.1-49.4 26.1-77.7c0-44.9-25.1-87.4-65.5-111.1a67.67 67.67 0 0 0-34.3-9.3H572.4l6-122.9c1.4-29.7-9.1-57.9-29.5-79.4A106.62 106.62 0 0 0 471 99.9c-52 0-98 35-111.8 85.1l-85.9 311h-.3v428h472.3c9.2 0 18.2-1.8 26.5-5.4c47.6-20.3 78.3-66.8 78.3-118.4c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c0-12.6-1.8-25-5.4-37c16.8-22.2 26.1-49.4 26.1-77.7c-.2-12.6-2-25.1-5.6-37.1zM112 528v364c0 17.7 14.3 32 32 32h65V496h-65c-17.7 0-32 14.3-32 32z">
                    </path>
                </svg>
                <span v-if="!like" class="toolbar-item__num">1.7万</span>
                <span v-else class="toolbar-item__num" style="color: #32a0d5;">1.7万</span>
            </div>
            <div class="toolbar-item" @click="star_change">
                <svg class="icon icon-star" v-if="!star" viewBox="0 0 1024 1024">
                    <path
                        d="M908.1 353.1l-253.9-36.9L540.7 86.1c-3.1-6.3-8.2-11.4-14.5-14.5c-15.8-7.8-35-1.3-42.9 14.5L369.8 316.2l-253.9 36.9c-7 1-13.4 4.3-18.3 9.3a32.05 32.05 0 0 0 .6 45.3l183.7 179.1l-43.4 252.9a31.95 31.95 0 0 0 46.4 33.7L512 754l227.1 119.4c6.2 3.3 13.4 4.4 20.3 3.2c17.4-3 29.1-19.5 26.1-36.9l-43.4-252.9l183.7-179.1c5-4.9 8.3-11.3 9.3-18.3c2.7-17.5-9.5-33.7-27-36.3zM664.8 561.6l36.1 210.3L512 672.7L323.1 772l36.1-210.3l-152.8-149L417.6 382L512 190.7L606.4 382l211.2 30.7l-152.8 148.9z">
                    </path>
                </svg>
                <svg class="icon" fill="#32a0d5" v-else viewBox="0 0 1024 1024">
                    <path
                        d="M908.1 353.1l-253.9-36.9L540.7 86.1c-3.1-6.3-8.2-11.4-14.5-14.5c-15.8-7.8-35-1.3-42.9 14.5L369.8 316.2l-253.9 36.9c-7 1-13.4 4.3-18.3 9.3a32.05 32.05 0 0 0 .6 45.3l183.7 179.1l-43.4 252.9a31.95 31.95 0 0 0 46.4 33.7L512 754l227.1 119.4c6.2 3.3 13.4 4.4 20.3 3.2c17.4-3 29.1-19.5 26.1-36.9l-43.4-252.9l183.7-179.1c5-4.9 8.3-11.3 9.3-18.3c2.7-17.5-9.5-33.7-27-36.3z">
                    </path>
                </svg>
                <span v-if="!star" class="toolbar-item__num">1.7万</span>
                <span v-else class="toolbar-item__num" style="color: #32a0d5;">1.7万</span>
            </div>
            <div class="toolbar-item">
                <svg class="icon icon-comment" viewBox="0 0 1024 1024">
                    <defs></defs>
                    <path
                        d="M573 421c-23.1 0-41 17.9-41 40s17.9 40 41 40c21.1 0 39-17.9 39-40s-17.9-40-39-40zm-280 0c-23.1 0-41 17.9-41 40s17.9 40 41 40c21.1 0 39-17.9 39-40s-17.9-40-39-40z">
                    </path>
                    <path
                        d="M894 345c-48.1-66-115.3-110.1-189-130v.1c-17.1-19-36.4-36.5-58-52.1c-163.7-119-393.5-82.7-513 81c-96.3 133-92.2 311.9 6 439l.8 132.6c0 3.2.5 6.4 1.5 9.4c5.3 16.9 23.3 26.2 40.1 20.9L309 806c33.5 11.9 68.1 18.7 102.5 20.6l-.5.4c89.1 64.9 205.9 84.4 313 49l127.1 41.4c3.2 1 6.5 1.6 9.9 1.6c17.7 0 32-14.3 32-32V753c88.1-119.6 90.4-284.9 1-408zM323 735l-12-5l-99 31l-1-104l-8-9c-84.6-103.2-90.2-251.9-11-361c96.4-132.2 281.2-161.4 413-66c132.2 96.1 161.5 280.6 66 412c-80.1 109.9-223.5 150.5-348 102zm505-17l-8 10l1 104l-98-33l-12 5c-56 20.8-115.7 22.5-171 7l-.2-.1C613.7 788.2 680.7 742.2 729 676c76.4-105.3 88.8-237.6 44.4-350.4l.6.4c23 16.5 44.1 37.1 62 62c72.6 99.6 68.5 235.2-8 330z">
                    </path>
                    <path d="M433 421c-23.1 0-41 17.9-41 40s17.9 40 41 40c21.1 0 39-17.9 39-40s-17.9-40-39-40z"></path>
                </svg>
                <span class="toolbar-item__num">1.7万</span>
            </div>
        </div>
        <div class="to-top">
            <svg class="icon" viewBox="0 0 1024 1024">
                <path
                    d="M890.5 755.3L537.9 269.2c-12.8-17.6-39-17.6-51.7 0L133.5 755.3A8 8 0 0 0 140 768h75c5.1 0 9.9-2.5 12.9-6.6L512 369.8l284.1 391.6c3 4.1 7.8 6.6 12.9 6.6h75c6.5 0 10.3-7.4 6.5-12.7z">
                </path>
            </svg>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { ref, getCurrentInstance, onMounted } from 'vue'

const like = ref(false)
const star = ref(false)

const like_change = () => {
    like.value = !like.value
}
const star_change = () => {
    star.value = !star.value
}

const instance = getCurrentInstance()

const toTop = ref(null as null | HTMLDivElement)

onMounted(() => {
    toTop.value = document.querySelector('.to-top')
})

instance?.proxy?.$Bus.on('showToTop', () => {
    console.log(111);
    toTop.value!.style.transform = 'translateY(-20)'
    toTop.value!.style.opacity = '1'
})

instance?.proxy?.$Bus.on('hideToTop', () => {
    toTop.value!.style.transform = 'translateY(0)'
    toTop.value!.style.opacity = '0'
})

</script>

<style scoped lang="less">
.right-side-bar {
    position: fixed;
    margin-left: 1026px;
    bottom: 60px;
    transition: 0.5s ease-in-out;
    margin-bottom: 50px;
    z-index: 10;
}

.side-toolbar {
    background-color: #fff;
    box-shadow: 0 2px 11px 0 rgba(0, 0, 0, 0.03);
    border-radius: 6px;
    padding: 10px 0;
}

.toolbar-item {
    display: flex;
    position: relative;
    color: #505050;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100px;
    margin-top: 8px;
    margin-bottom: 8px;
    cursor: pointer;

    .icon {
        margin: 0 20px;
        width: 48px;
        transition: fill 0.3s ease-in-out;
    }

    .icon-like {
        fill: #505050;
    }

    .icon-star {
        fill: #505050;
    }

    .icon-comment {
        fill: #505050;
    }

    .toolbar-item__num {
        font-size: 14px;
        transition: 0.3s;
        margin-top: 8px;
        line-height: 8px;
    }

    &:hover {
        .icon-like {
            fill: #32a0d5;
        }

        .icon-star {
            fill: #32a0d5;
        }

        .icon-comment {
            fill: #32a0d5;
        }

        .toolbar-item__num {
            color: #32a0d5;
        }
    }
}
.to-top {
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #fff;
    width: 88px;
    height: 88px;
    border-radius: 6px;
    transition: transform .3s, opacity .3s;
    box-shadow: 0 2px 11px 0 rgba(0,0,0,.03);
    margin-top: 15px;
    opacity: 0;
    // transform: translateY(0);
    .icon {
        width: 36px;
        fill: #505050;
        &:hover {
            fill: #32a0d5;
        }
    }
}
</style>