import dayjs from 'dayjs'
export const formatTime = (time: string) => {
  return dayjs.unix(Number(time)).format('YYYY年MM月DD日 HH:mm')
}



const LESSTHANONEMINUTE = "刚刚"
const LESSTHANONEHOUR = "分钟前"
const LESSTHANONEDAY = "小时前"
const LESSTHANONEWEEK = "天前"
const LESSTHANONEYEAR = ''


export const timeDisplay = (time: string) => {
  const now = dayjs()
  const tar = dayjs.unix(Number(time))
  const secDiff = now.diff(tar, 'second')
  if(secDiff < 60) return LESSTHANONEMINUTE
  const minDiff = now.diff(tar, 'minute')
  if(minDiff < 90) return `${minDiff}${LESSTHANONEHOUR}`
  const hourDiff = now.diff(tar, 'hour')
  if(hourDiff < 24) return `${hourDiff}${LESSTHANONEDAY}`
  const dayDiff = now.diff(tar, 'day')
  if(dayDiff < 7) return `${dayDiff}${LESSTHANONEWEEK}`
  const yearDiff = now.diff(tar, 'year')
  if(yearDiff < 1) return tar.format('MM月DD日')
  return tar.format('YYYY年MM月DD日')
}