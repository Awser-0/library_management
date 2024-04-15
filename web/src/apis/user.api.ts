export { register, updateUser, updateOhterUser, queryUser, login, loginByToken };
export type { UserInfo };

import { request } from "./request";
import * as consts from "~/consts";
import type { Result } from "./request";

type LoginResult = {
	token?: string;
	user: UserInfo;
};

type UserInfo = {
	id: number;
	username: string;
	name: string;
	sex: "0" | "1";
	phone: string;
	birth: string;
	isAdmin: boolean;
	update_time: string;
	create_time: string;
};

type UserUpdateForm = {
	username: string;
	name: string;
	sex: "0" | "1";
	phone: string;
	birth: Date | null;
	isAdmin: boolean;
};

// 注册用户
async function register(form: {
	username: string;
	password: string;
	name: string;
	sex: "0" | "1";
	phone: string;
	birth: Date | null;
	isAdmin: boolean;
}) {
	return request<Result<boolean>>({
		url: "/v1/user/register",
		method: "POST",
		data: { ...form },
	});
}

// 更新用户
async function updateUser(form: UserUpdateForm) {
	return request<Result<boolean>>({
		url: "/v1/user/update",
		method: "POST",
		data: { ...form },
	});
}

// 更新其他用户
async function updateOhterUser(id: number, form: UserUpdateForm) {
	return request<Result<boolean>>({
		url: "/v1/user/update",
		method: "POST",
		data: { id, ...form },
	});
}

// 查询用户，查询限制 用户名和姓名
async function queryUser(query_string: string) {
	return request<Result<{ users: UserInfo[] }>>({
		url: "/v1/user/query",
		method: "POST",
		data: { query_string },
	});
}

// 通过账号密码进行登录
async function login(username: string, password: string) {
	return request<Result<LoginResult>>({
		url: "/v1/user/login",
		method: "POST",
		data: { username, password },
	});
}

// 通过token进行登录
async function loginByToken(token: string) {
	return request<Result<LoginResult>>({
		url: "/v1/user/login/token",
		method: "POST",
		headers: {
			[consts.REQUEST_HEADER_KEY_TOKEN]: token,
		},
	});
}
