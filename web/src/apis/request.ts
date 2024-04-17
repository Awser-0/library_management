export { request };
export type { Result };

import axios, { AxiosError } from "axios";
import * as consts from "~/consts";

type Result<Data = never> = {
	code: number;
	msg: string;
	data: Data;
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
