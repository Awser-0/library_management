export { request };
export type { Result, PageReq, PageRes };

import axios, { AxiosError } from "axios";
import * as consts from "~/consts";

// 响应结果格式
type Result<Data = never> = {
	code: number;
	msg: string;
	data: Data;
};
// 分页请求格式
type PageReq = {
	page_num?: number;
	page_size?: number;
};
// 分页响应格式
type PageRes<T = any> = {
	list: T[];
	page_num: number;
	page_size: number;
	total: number;
};

const request = axios.create({
	baseURL: import.meta.env.DEV ? "/API" : import.meta.env.VITE_BASE_URL,
});

request.interceptors.request.use((config) => {
	config.headers[consts.REQUEST_HEADER_KEY_TOKEN] = window.localStorage.getItem(
		consts.WINDOW_LOCALSTORAGE_KEY_TOKEN,
	);
	return config;
});

request.interceptors.response.use(
	(data) => {
		return data;
	},
	(err) => {
		if (err instanceof AxiosError) {
			switch (err.response?.status) {
				case 401: {
					ElMessage.error("身份信息无效，请重新登录");
					break;
				}
				case 403: {
					ElMessage.error("无权限操作");
					break;
				}
			}
		}
	},
);
