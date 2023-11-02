import {createDiscreteApi} from 'naive-ui'
import type { MessageOptions } from 'naive-ui'

const {
    message: _message,
    notification: _notification,
    dialog: _dialog,
    loadingBar: _loadingBar,
} = createDiscreteApi(['message','notification','dialog','loadingBar'])


export const loadingBar = {
    start() {
        _loadingBar.start()
    },
    error() {
        _loadingBar.error()
    },
    finish() {
        _loadingBar.finish()
    },
}

export const message = {
    create(content:string, options?:MessageOptions) {
        return _message.create(content, options)
    },
    error(content:string, options?:MessageOptions) {
        return _message.error(content, options)
    },
    info(content:string, options?:MessageOptions) {
        return _message.info(content, options)
    },
    loading(content:string, options?:MessageOptions) {
        return _message.loading(content, options)
    },
    success(content:string, options?:MessageOptions) {
        return _message.success(content, options)
    },
    warning(content:string, options?:MessageOptions) {
        return _message.warning(content, options)
    },
}