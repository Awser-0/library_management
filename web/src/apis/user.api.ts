export { login };

import { request } from "./request";
import type { Result } from "./request";

type LoginResult = {
	token?: string;
	user: {
		id: number;
		username: string;
		sex: "0" | "1";
		phone: string;
		birth: string;
		isAdmin: boolean;
		update_time: string;
		create_time: string;
	};
};

async function login(username: string, password: string) {
	return request<Result<LoginResult>>({
		url: "/v1/user/login",
		method: "POST",
		data: { username, password },
	});
}
