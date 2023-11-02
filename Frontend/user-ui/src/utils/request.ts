import axios from "axios";
import { message } from "@/utils/client";
import { useRouter } from 'vue-router'

const baseURL = 'http://124.220.25.230:9000'

const router = useRouter()

let request = axios.create({
    baseURL,
    timeout: 8000,
})

request.interceptors.request.use(
    function(config){
        let token = localStorage.getItem('token')
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`
        }
        return config
    },
    function(error){
        return Promise.reject(error)
    }
)

request.interceptors.response.use(
    function (res) {
        if(res.data.status === 200) {
            return res.data
        } else if (res.data.status !== undefined) {
            message.error(res.data.message)
            return Promise.reject(new Error(res.data.message))
        } else {
            return res
        }
    },
    function (error) {
        message.error(error.response?.data.message)
        console.log("error:", error.response);
        if(localStorage.getItem('token')){
            localStorage.removeItem('token')
        }
        router.push('/login')
        return Promise.reject(error)
    }
)

declare module 'axios' {
    interface AxiosInstance {
        (config: AxiosRequestConfig): Promise<any>
    }
}

export default request