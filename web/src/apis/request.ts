export { request };
export type { Result };
import axios from "axios";

type Result<Data = unknown> = {
	code: number;
	msg: string;
	data: Data;
};

const request = axios.create({
	baseURL: import.meta.env.DEV ? "/API" : import.meta.env.VITE_BASE_URL,
});
