<template>
    <svg :class="svgClass" aria-hidden="true">
        <use :xlink:href="iconClassName" :fill="color"></use>
    </svg>
</template>

<script setup lang='ts' name="SvgIcon">
import { computed } from 'vue';
const props = defineProps({
    iconName: {
        type: String,
        required: true
    },
    className: {
        type: String,
        default: ''
    },
    // color: {
    //     type: String,
    //     default: '#2c2c2c'
    // },
    activated: {
        type: Boolean,
        default: false
    }
})

const iconClassName = computed(() => {
    return `#${props.iconName}`
})

const svgClass = computed(() => {
    if (props.className) {
        return `svg-icon ${props.className}`
    }
    return 'svg-icon'
})
const color = computed(() => {
    if (props.activated) {
        return '#1E90FF'
    }
})
</script>

<style scoped>
.svg-icon {
    width: 3rem;
    height: 3rem;
    /* vertical-align: -0.15em; */
    /* fill: currentColor; */
    /* position: relative; */
}
</style>