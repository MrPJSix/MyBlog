import { createRouter, createWebHistory } from "vue-router";

const Login = () => import('../views/Login.vue');
const Admin = () => import('../views/Admin.vue');

const Index = () => import('../components/admin/Index.vue');
const UserManagement = () => import('../components/admin/UserManagement.vue')

const routes = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/admin',
        name: 'Admin',
        component: Admin,
        children: [
            { path: 'index', name: 'Index', component: Index },
            { path: 'user', name: 'User', component: UserManagement },
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router