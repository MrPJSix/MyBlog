import request from "@/utils/request";
import type {RegisterForm, LoginForm, Profile} from '@/interface/index'

export const register = (registerForm: RegisterForm) => {
    return request({
        url: '/register',
        method: 'post',
        data: registerForm,
    })
}

export const login = (loginForm: LoginForm) => {
    return request({
        url: '/login',
        method: 'post',
        data: loginForm,
    })
}

export const getSelfProfile = () => {
    return request({
        url: '/user/self/profile',
        method: 'get',
    })
}

export const updateProfile = (profile: Profile) => {
    return request({
        url: 'user/self/profile',
        method: 'put',
        data: profile,
    })
}

export const postAvatar = (avatar: File) => {
    const formData = new FormData()
    formData.append('avatar', avatar)
    return request({
        url: '/user/avatar',
        method: 'post',
        data: formData,
    })
}