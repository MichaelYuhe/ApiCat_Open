import { UserInfo } from '@/typings/user'
import Ajax from './Ajax'
import useApi from '@/hooks/useApi'

// 用户登录
export const userEmailLogin = (data: Partial<UserInfo>) => Ajax.post('/account/login/email', { ...data })
// 用户注册
export const userRegister = (data: Partial<UserInfo>) => Ajax.post('/account/register/email', { ...data })
// 退出
export const logout = () => Ajax.post('/logout')

export const modifyPassword = (data = {}) => Ajax.put('/user/password', { ...data })

export const modifyUserInfo = (data = {}) => Ajax.put('/user', { ...data })
