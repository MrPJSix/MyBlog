import { createRouter, createWebHistory } from "vue-router";

const Login = () => import('../views/Login.vue');
const Admin = () => import('../views/Admin.vue');

const Index = () => import('../components/admin/Index.vue');
const UserManagement = () => import('../components/admin/UserManagement.vue');
const CategoryManagement = () => import('../components/admin/CategoryManagement.vue');
const ArticleManagement = () => import('../components/admin/ArticleManagement.vue');
const CommentManagement = () => import('../components/admin/CommentManagement.vue');

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
            { path: 'category', name: 'Category', component: CategoryManagement },
            { path: 'article', name: 'Article', component: ArticleManagement },
            { path: 'comment', name: 'Comment', component: CommentManagement },
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router