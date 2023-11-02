<template>
    <div class="user-center-config">
        <div class="user-avatar-config">
            <h2>更换头像</h2>
            <div class="avatar-ops">
                <div class="left">
                    <div class="img-preview-wrap">
                        <div class="pre-container">
                            <img src="@/assets/avatar/kasha.jpeg" id="clip_res_img"
                                style="width: 1000x; height: 100px; top: 0; left: 0;">
                        </div>
                        <div class="pre-info">{{ !isUpload ? '原始头像' : '预览头像' }}</div>
                    </div>
                </div>
                <div class="border-line"></div>
                <div class="right">
                    <div class="user-ops" v-show="!isUpload">
                        <input @change="upload" id="uploadAvatar" type="file" accept="image/*" style="display: none;">
                        <label for="uploadAvatar">
                            <span>选择本地图片</span>
                        </label>
                    </div>
                    <div class="container-bg" style="height: 200px;" v-show="isUpload">
                        <div class="img-container">
                            <img id="clip_src_img" src="" alt="">
                            <div class="shadow-box"></div>
                            <div class="crop-wrap">
                                <div class="shadow-box" style="width: 0; height: 0; left: 0; top: 0;">
                                    <img class="shadow-img" src="" alt="" style="width: 0; height: 0; left: 0; top: 0;">
                                </div>
                                <div class="crop-box show" style="width: 0; height: 0; left: 0; top: 0;"
                                    @mousedown="onMouseDown($event)">
                                    <span class="drag-point pointlt"></span>
                                    <span class="drag-point pointlb"></span>
                                    <span class="drag-point pointrt"></span>
                                    <span class="drag-point pointrb"></span>
                                </div>
                            </div>
                        </div>
                        <div class="reset-img">
                            <label for="uploadAvatar">重新选择</label>
                        </div>
                    </div>
                </div>
            </div>
            <button class="save-btn" v-if="isUpload" @click="save">点击上传</button>
        </div>
        <div class="user-info-config">
            <h2>修改资料</h2>
            <div class="info-ops"></div>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { ref, onMounted } from 'vue'
import api from '@/api/index'
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';

const router = useRouter()

const isUpload = ref(false);

let startX = 0;
let startY = 0;
let cropBoxWidth = 0;
let cropBoxHeight = 0;
let imageWidth = 0;
let imageHeight = 0;
let mouseDownPoint: string = '';
let previewImagewidth = 100;
let previewImageHeight = 0;
let ratio = 0;
let rawImgWidth = 0;
let rawImgHeight = 0;

const getImageNewInfo = (imgSrc: string): Promise<{ width: number, height: number }> => {
    return new Promise((resolve, reject) => {
        let img = new Image();
        img.src = imgSrc;
        let width: number = 0;
        let height: number = 0;
        img.onload = function () {
            rawImgWidth = img.width
            rawImgHeight = img.height
            if (img.height > img.width) {
                height = 200
                width = img.width * 200 / img.height;
                previewImagewidth = 100
                previewImageHeight = height * 100 / width
                ratio = 100 / width
            } else {
                width = 200
                height = img.height * 200 / img.width
                previewImageHeight = 100
                previewImagewidth = width * 100 / height
                ratio = 100 / height
            }
            resolve({ width, height });
        }
    })
}

const getImgSrc = async (file: File): Promise<string> => {
    return new Promise((resolve, reject) => {
        let fileReader = new FileReader();
        fileReader.readAsDataURL(file);
        let data = '';
        fileReader.onload = function () {
            data = this.result as string;
            resolve(data);
        }
    })
}

const imgContainer = ref(null as unknown as HTMLElement)
const shadowBox = ref(null as unknown as HTMLElement)
const clipImg = ref(null as unknown as HTMLImageElement)
const cropBox = ref(null as unknown as HTMLElement)
const shadowImg = ref(null as unknown as HTMLImageElement)
const dragPoints = ref(null as unknown as HTMLCollectionOf<HTMLSpanElement>)
const resImg = ref(null as unknown as HTMLImageElement)

onMounted(() => {
    imgContainer.value = document.getElementsByClassName('img-container')[0] as HTMLDivElement;
    clipImg.value = document.getElementById('clip_src_img') as HTMLImageElement;
    shadowBox.value = document.getElementsByClassName('shadow-box')[1] as HTMLDivElement;
    cropBox.value = document.getElementsByClassName('crop-box')[0] as HTMLDivElement;
    shadowImg.value = document.getElementsByClassName('shadow-img')[0] as HTMLImageElement;
    dragPoints.value = document.getElementsByClassName('drag-point') as HTMLCollectionOf<HTMLSpanElement>
    for (let i = 0; i < dragPoints.value.length; i++) {
        dragPoints.value[i].addEventListener('mousedown', onMouseDownForPoints)
    }
    resImg.value = document.getElementById('clip_res_img') as HTMLImageElement;
})

const upload = () => {
    const file = document.getElementById('uploadAvatar') as HTMLInputElement;
    if (file.files) {
        getImgSrc(file.files[0]).then(imgSrc => {
            clipImg.value.src = imgSrc
            shadowImg.value.src = imgSrc
            resImg.value.src = imgSrc
            getImageNewInfo(imgSrc).then(({ width, height }) => {
                imageWidth = width
                imageHeight = height
                imgContainer.value.style.width = imageWidth + 'px';
                imgContainer.value.style.height = imageHeight + 'px';
                clipImg.value.style.width = imageWidth + 'px';
                clipImg.value.style.height = imageHeight + 'px';
                resImg.value.style.width = previewImagewidth + 'px';
                resImg.value.style.height = previewImageHeight + 'px';
                // reset
                cropBox.value.style.left = '0px';
                cropBox.value.style.top = '0px';
                shadowBox.value.style.left = '0px';
                shadowBox.value.style.top = '0px';
                shadowImg.value.style.left = '0px';
                shadowImg.value.style.top = '0px';
                resImg.value.style.top = '0px';
                resImg.value.style.left = '0px';

                if (imageHeight > imageWidth) {
                    shadowBox.value.style.width = imageWidth + 'px'
                    shadowBox.value.style.height = imageWidth + 'px'
                    shadowImg.value.style.width = imageWidth + 'px'
                    shadowImg.value.style.height = imageHeight + 'px'
                    cropBox.value.style.width = imageWidth + 'px';
                    cropBox.value.style.height = imageWidth + 'px';
                } else {
                    shadowBox.value.style.height = imageHeight + 'px'
                    shadowBox.value.style.width = imageHeight + 'px'
                    shadowImg.value.style.width = imageWidth + 'px'
                    shadowImg.value.style.height = imageHeight + 'px'
                    cropBox.value.style.height = imageHeight + 'px'
                    cropBox.value.style.width = imageHeight + 'px'
                }
            })
            isUpload.value = true;
        });

    }
}

// 节流
function throttle<T extends (...args: any[]) => void>(fn: T, delay: number): (...args: Parameters<T>) => void {
    let timerId: number | undefined
    let lastExecTime = 0
    return function (this: ThisParameterType<T>, ...args: any[]) {
        const currentTimeStamp = Date.now()
        if (currentTimeStamp - lastExecTime >= delay) {
            fn.apply(this, args)
            lastExecTime = currentTimeStamp
        } else {
            clearTimeout(timerId)
            timerId = setTimeout(() => {
                fn.apply(this, args)
                lastExecTime = currentTimeStamp
            }, delay)
        }
    }
}

const handleMouseMoveForBox = (e: MouseEvent) => {
    const curX = e.clientX;
    const curY = e.clientY;
    const curLeft = parseInt(cropBox.value.style.left);
    const curTop = parseInt(cropBox.value.style.top);
    let moveX = 0
    let moveY = 0
    const target = e.target as HTMLElement;

    if (target.tagName === 'DIV') {
        if (curLeft + curX - startX < 0 || curLeft + curX - startX > imageWidth - cropBoxWidth) {
            moveX = 0
        } else {
            moveX = curX - startX
        }
        if (curTop + curY - startY < 0 || curTop + curY - startY > imageHeight - cropBoxHeight) {
            moveY = 0
        } else {
            moveY = curY - startY
        }
        cropBox.value.style.left = curLeft + moveX + 'px';
        cropBox.value.style.top = curTop + moveY + 'px';
        shadowBox.value.style.left = curLeft + moveX + 'px';
        shadowBox.value.style.top = curTop + moveY + 'px';
        shadowImg.value.style.left = -(curLeft + moveX) + 'px';
        shadowImg.value.style.top = -(curTop + moveY) + 'px';

        resImg.value.style.top = -(curTop + moveY) * ratio + 'px';
        resImg.value.style.left = -(curLeft + moveX) * ratio + 'px';
    }
    startX = curX;
    startY = curY;
}

const handleMouseMoveForPoints = (e: MouseEvent) => {

    const curX = e.clientX;
    const curY = e.clientY;

    const curLeft = parseInt(cropBox.value.style.left);
    const curTop = parseInt(cropBox.value.style.top);
    if (true) {
        cropBoxWidth = cropBox.value.offsetWidth;
        cropBoxHeight = cropBoxWidth
        let temp = curX - startX;
        let flagX = 0  // left 
        let flagY = 0  // top
        let flag = 0 // 放大 缩小

        // 左上角
        if (mouseDownPoint.includes('pointlt')) {
            flagX = 1
            flagY = 1
            if (temp > 0 && temp < cropBoxWidth) {
                flag = -1
            } else if (temp < 0 && curLeft > -temp && curTop > -temp) {
                flag = 1
            }
        }

        // 左下角
        if (mouseDownPoint.includes('pointlb')) {
            flagX = 1
            if (temp > 0 && temp < cropBoxWidth) {
                flag = -1
            } else if (temp < 0 && curLeft < -temp && -temp + cropBoxWidth + curTop < imageHeight) {
                flag = 1
            }
        }
        // 右上角
        if (mouseDownPoint.includes('pointrt')) {
            flagY = 1
            if (temp > 0 && curTop > temp && curLeft + cropBoxWidth + temp < imageWidth) {
                flag = 1
            } else if (temp < 0 && -temp < cropBoxWidth) {
                flag = -1
            }
        }

        // 右下角
        if (mouseDownPoint.includes('pointrb')) {
            if (temp > 0 && curTop + cropBoxWidth + temp < imageHeight && curLeft + cropBoxWidth + temp < imageWidth) {
                flag = 1
            } else if (temp < 0 && -temp < cropBoxWidth) {
                flag = -1
            }
        }
        temp = Math.abs(temp)
        cropBox.value.style.width = cropBoxWidth + temp * flag + 'px'
        cropBox.value.style.height = cropBox.value.style.width
        cropBox.value.style.left = curLeft - temp * flag * flagX + 'px'
        cropBox.value.style.top = curTop - temp * flag * flagY + 'px'
        shadowBox.value.style.width = cropBox.value.style.width
        shadowBox.value.style.height = cropBox.value.style.width
        shadowBox.value.style.top = cropBox.value.style.width
        shadowBox.value.style.left = cropBox.value.style.left
        shadowBox.value.style.top = cropBox.value.style.top
        shadowImg.value.style.left = '-' + cropBox.value.style.left
        shadowImg.value.style.top = '-' + cropBox.value.style.top

        resImg.value.style.width = 100 * imageWidth / parseInt(cropBox.value.style.width) + 'px'
        resImg.value.style.height = 100 * imageHeight / parseInt(cropBox.value.style.height) + 'px'
        ratio = parseInt(resImg.value.style.width) / imageWidth // 预览图宽度 / 原图宽度

        resImg.value.style.top = -(parseInt(cropBox.value.style.top) * ratio) + 'px'
        resImg.value.style.left = -(parseInt(cropBox.value.style.left) * ratio) + 'px'

        previewImagewidth = parseInt(resImg.value.style.width)
        previewImageHeight = parseInt(resImg.value.style.height)
    }
    startX = curX;
    startY = curY;
}

const throttledHandleMouseMoveForBox = throttle(handleMouseMoveForBox, 10)
const throttledHandleMouseMoveForPoints = throttle(handleMouseMoveForPoints, 10)
const onMouseDown = (e: MouseEvent) => {
    cropBoxWidth = cropBox.value.offsetWidth;
    cropBoxHeight = cropBoxWidth
    startX = e.clientX;
    startY = e.clientY;
    document.addEventListener('mousemove', throttledHandleMouseMoveForBox)
    document.addEventListener('mouseup', () => {
        document.removeEventListener('mousemove', throttledHandleMouseMoveForBox)
    })
}


const onMouseDownForPoints = (e: MouseEvent) => {
    e.stopPropagation()
    const target = e.target as HTMLElement;
    mouseDownPoint = target.className
    startX = e.clientX;
    startY = e.clientY;
    document.addEventListener('mousemove', throttledHandleMouseMoveForPoints)
    document.addEventListener('mouseup', () => {
        document.removeEventListener('mousemove', throttledHandleMouseMoveForPoints)
    })
}

const dataURLtoFile = (dataUrl:string, fileName:string):File => {
    let arr = dataUrl.split(',')
    let mime = arr[0].match(/:(.*?);/)![1]
    let bstr = atob(arr[1])
    let n = bstr.length
    let u8arr = new Uint8Array(n)
    while(n--){
        u8arr[n] = bstr.charCodeAt(n)
    }
    return new File([u8arr],fileName,{type:mime})
}


const save = () => {
    const canvas = document.createElement('canvas');
    const ctx = canvas.getContext('2d');
    let ratio = 200 / cropBox.value.offsetWidth
    canvas.width = imageWidth * ratio
    canvas.height = imageHeight * ratio
    ctx?.drawImage(clipImg.value, 0, 0, imageWidth * ratio, imageHeight * ratio)
    const sourceX = parseInt(cropBox.value.style.left) * ratio
    const sourceY = parseInt(cropBox.value.style.top) * ratio

    const imageData = ctx?.getImageData(sourceX, sourceY, 200, 200)
    const croppedCanvas = document.createElement('canvas');
    const croppedctx = croppedCanvas.getContext('2d')
    croppedCanvas.width = 200
    croppedCanvas.height = 200
    croppedctx?.putImageData(imageData as ImageData, 0, 0)
    const dataURL = croppedCanvas.toDataURL('image/jpeg')
    // const link = document.createElement('a')
    // link.href = dataURL
    // link.download = "cropped_image.jpeg"
    // const a = link.click()
    const file = dataURLtoFile(dataURL, 'avatar.jpeg')
    api.user.postAvatar(file).then(res => {
        if (res.status === 200) {
            ElMessage({
                message: '头像上传成功',
                type: 'success'
            })
            router.push('/usercenter')
        }
    })
}

</script>

<style scoped></style>