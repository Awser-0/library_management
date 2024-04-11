export { userUserStore };

import { defineStore } from "pinia";
import { ref, reactive } from "vue";
import * as consts from "~/consts";

type UserInfo = {
	id: number;
	username: string;
	sex: "0" | "1";
	phone: string;
	birth?: string;
	update_time: string;
	create_time: string;
};

const userUserStore = defineStore("user_store", () => {
	const isLogin = ref(false);
	const info = ref<UserInfo>(resetInfo());
	function resetInfo(): UserInfo {
		return {
			id: 0,
			username: "",
			sex: "0",
			phone: "",
			birth: "",
			update_time: "",
			create_time: "",
		};
	}
	function signIn(userInfo: UserInfo, token?: string) {
		info.value = userInfo;
		isLogin.value = true;
		if (token) {
			window.localStorage.setItem(consts.WINDOW_LOCALSTORAGE_KEY_TOKEN, token);
		}
	}
	function signOut() {
		isLogin.value = false;
		info.value = resetInfo();
		window.localStorage.removeItem(consts.WINDOW_LOCALSTORAGE_KEY_TOKEN);
	}
	return {
		isLogin,
		info,
		signIn,
		signOut,
	};
});
