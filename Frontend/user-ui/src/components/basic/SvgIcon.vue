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
    color: {
        type: String,
        default: '#2c2c2c'
    }
})

const iconClassName = computed(() => {
    console.log(props.iconName);
    return `#${props.iconName}`
})

const svgClass = computed(() => {
    if (props.className) {
        return `svg-icon ${props.className}`
    }
    return 'svg-icon'
})
</script>

<style scoped>
.svg-icon {
    width: 3rem;
    height: 3rem;
    /* vertical-align: -0.15em; */
    fill: currentColor;
    /* position: relative; */
}
</style>