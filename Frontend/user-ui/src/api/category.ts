import request from "@/utils/request";

export const getPrimary = () => {
    return request({
        url: '/categories/primary',
        method: 'get',
    })
}

export const getSecondary = (primaryId: number) => {
    return request({
        url: `/category/${primaryId}/subs`,
        method: 'get',
    })
}

export const getCategory = (catogoryId: number) => {
    return request({
        url: `/category/${catogoryId}`,
        method: 'get',
    })
}