<template>
    <div class="home-sub-menu-bar" v-if="!all">

        <div class="home-sub-menu-bar-item" v-for="(sub, index) in subCategories" :key="index" @click="switchPage(index)">
            <span class="home-sub-menu-bar-item__content">{{ sub.title }}</span>
            <span class="home-sub-menu-item__overlay"></span>
        </div>
        <div id="line" class="home-sub-menu-move-line"></div>
    </div>
    <div class="home-content-items">
        <Item v-for="(article, index) in articles" :article="article" :key="index" @click="toArticle(article.id)">
        </Item>
    </div>
    <div class="content-pagination" style="display: flex; justify-content: center; margin: 10px 0;">
        <el-pagination layout="prev, pager, next" :total="articles_count" :current-page.sync="currentPage"
            @update:current-page="fetchData($event)" />
    </div>
</template>

<script setup lang='ts'>
import Item from './Item.vue';
import { reactive, ref, onMounted, getCurrentInstance, watch, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/api/index'
import { SubCategory, Config, ArticleReciveData } from '@/interface'
import { ElPagination } from 'element-plus';

const all = ref(true)

const articles_count = ref(0)

const currentPage = ref(1)

const categoryId = ref(0)

let lastActive = ref(0); // index
// let interval: number | null = null;
const line = ref(null as HTMLElement | null);
const lastOverLay = ref(null as HTMLElement | null);
const lastContent = ref(null as HTMLElement | null);
const instance = getCurrentInstance();
const subCategories = reactive([] as SubCategory[])
const PrimaryCategoryId = ref(-1)
const PageConfig = reactive({
    pagesize: 10,
    pagenum: 1
} as Config)
const articles = reactive([] as ArticleReciveData[])

const fillInArticles = (data: ArticleReciveData[]) => {

    articles.splice(0, articles.length)
    data.forEach(item => {
        articles.push(item)
    })
}


// 如果id为0，代表子菜单栏选中 【全部】
const getArticles = (index: number) => {
    articles.splice(0, articles.length)
    let id = subCategories[index].id;
    id = id === 0 ? PrimaryCategoryId.value : id;
    api.article.getgetArticlesCountByCategory(id).then(res => {
        articles_count.value = res.data
    })
    api.article.getArticlesByCategory(id, PageConfig).then(res => {
        if (res.data.length !== 0) {
            fillInArticles(res.data)
        }
    })
}

// 接收到的id为0，以为主菜单选中 【查看全部】
instance?.proxy?.$Bus.on('change-category', (id) => {

    if (id === PrimaryCategoryId.value) {
        return;
    } else {
        subCategories.splice(0, subCategories.length)
        if (line.value) {
            line.value.style.left = '0px';
        }
        lastActive.value = 0;
    }

    PrimaryCategoryId.value = id as number;

    if (id === 0) {
        all.value = true
        categoryId.value = 0
        api.article.getAllArticlesCount().then(res => {
            articles_count.value = res.data
        })
        api.article.getAllArticles(PageConfig).then(res => {
            fillInArticles(res.data);
        })
    } else {
        all.value = false
        subCategories.push({
            id: PrimaryCategoryId.value,
            title: '全部',
            parentId: -1,
            activated: true
        })
        api.category.getSecondary(id as number).then(res => {
            res.data.forEach((item: any) => {
                subCategories.push({
                    id: item.id,
                    title: item.name,
                    parentId: id as number,
                    activated: false
                })
            })
        })
        getArticles(0)
    }
})


const switchPage = (index: number) => {
    if (index === lastActive.value) return;
    const target = document.querySelector(`.home-sub-menu-bar-item:nth-child(${index + 1})`) as HTMLElement;
    line.value!.style.left = target.offsetLeft + 'px';
    lastActive.value = index;
    lastOverLay.value?.classList.remove('overlay__active')
    lastContent.value?.classList.remove('content__active')
    const targetOverLay = target.querySelector(`.home-sub-menu-bar-item:nth-child(${index + 1}) .home-sub-menu-item__overlay`) as HTMLElement;
    targetOverLay.classList.add('overlay__active')
    const targetContent = target.querySelector(`.home-sub-menu-bar-item:nth-child(${index + 1}) .home-sub-menu-bar-item__content`) as HTMLElement;
    targetContent.classList.add('content__active')
    lastOverLay.value = targetOverLay;
    lastContent.value = targetContent;
    categoryId.value = subCategories[index].id;
    currentPage.value = 1;
    getArticles(index)
}


onMounted(() => {
    lastOverLay.value = document.querySelector(`.home-sub-menu-bar-item:nth-child(1) .home-sub-menu-item__overlay`) as HTMLElement;
    lastOverLay.value?.classList.add('overlay__active')
    lastContent.value = document.querySelector(`.home-sub-menu-bar-item:nth-child(1) .home-sub-menu-bar-item__content`) as HTMLElement;
    lastContent.value?.classList.add('content__active')
})

const updatePos = () => {
    const target = document.querySelector(`.home-sub-menu-bar-item:nth-child(${lastActive.value + 1})`) as HTMLElement;
    line.value!.style.left = target.offsetLeft + 'px';
}

const debounce = (func: () => void, delay: number) => {
    let timerId: number | null = null;
    return function () {
        if (timerId) {
            clearTimeout(timerId);
        }
        timerId = setTimeout(func, delay);
    }
}

watch(all, async () => {
    console.log(all.value);
    if (all.value === false) {
        nextTick(() => {
            line.value = document.querySelector('#line') as HTMLElement;
            window.addEventListener('resize', debounce(updatePos, 100))
        })
    } else {
        window.removeEventListener('resize', debounce(updatePos, 100))
    }
})

const router = useRouter();

const toArticle = (id: number) => {
    router.push(`/article/${id}`)
}

const fetchData = (xxx: number) => {
    currentPage.value = xxx;
    if (categoryId.value === 0) {
        api.article.getAllArticles({ pagenum: currentPage.value, pagesize: 10 }).then(res => {
            if (res.data.length !== 0) {
                
                fillInArticles(res.data)
            }
        })
    } else {
        api.article.getArticlesByCategory(categoryId.value, { pagenum: currentPage.value, pagesize: 10 }).then(res => {

            if (res.data.length !== 0) {
                fillInArticles(res.data)
            }
        })
    }
    instance?.proxy?.$Bus.emit('scroll-to-top')
}

</script>

<style scoped></style>