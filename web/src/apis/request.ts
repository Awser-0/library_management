export { request };
export type { Result };

import axios from "axios";
import * as consts from "~/consts";

type Result<Data = unknown> = {
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
