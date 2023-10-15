export function getCurrentTime() {
    // 获取当前时间
    const currentDate = new Date();

    // 获取年月份
    const year = currentDate.getFullYear();
    const month = currentDate.getMonth() + 1; // 月份从 0 开始，所以需要加 1
    const day = currentDate.getDate();

    // 获取时分秒
    const hours = currentDate.getHours();
    const minutes = currentDate.getMinutes();
    const seconds = currentDate.getSeconds();

    return {
        day,
        year,
        month,
        hours,
        minutes,
        seconds
    };
}