import { createRouter, createWebHistory } from "vue-router";

const Login = () =>  import('../views/Login.vue');
const Admin = () => import('../views/Admin.vue');

const routes = [
    {
        path:'/index',
        name:'Index',
        component: () => import('../views/Index.vue')
    },
    {
        path:'/login',
        name:'Login',
        component: Login
    },
    {
        path:'/admin',
        name:'Admin',
        component: Admin,
        children: [
            {path:'/index2', name:'Index2', component: () => import('../views/Index2.vue')},
            {
                path:'/login',
                name:'Login',
                component: Login
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router