import { defineStore } from 'pinia'
import api from '@/api'

export const useUserDataStore = defineStore({
    id: 'user-data',
    state: () => ({
        user_name: '',
        full_name: '',
        avatar_url: '',
    }),
    actions: {
        updateUserData() {
            api.user.getSelfProfile().then(res => {
                this.user_name = res.data.username
                this.full_name = res.data.full_name
                this.avatar_url = res.data.avatar_url
            })
        }
    }
})