export { useUserStore };

import { defineStore } from "pinia";
import { ref, reactive } from "vue";
import * as consts from "~/consts";

type UserInfo = {
	id: number;
	username: string;
	nickname: string;
	sex: "0" | "1";
	phone: string;
	birth?: string;
	update_time: string;
	create_time: string;
};

const useUserStore = defineStore("user_store", () => {
	const is_login = ref(false);
	const is_admin = ref(false);
	const info = ref<UserInfo>(resetInfo());
	function resetInfo(): UserInfo {
		return {
			id: 0,
			username: "",
			nickname: "",
			sex: "0",
			phone: "",
			birth: "",
			update_time: "",
			create_time: "",
		};
	}
	function signIn(userInfo: UserInfo, isAdmin: boolean, token?: string) {
		info.value = userInfo;
		is_login.value = true;
		is_admin.value = isAdmin;
		if (token) {
			window.localStorage.setItem(consts.WINDOW_LOCALSTORAGE_KEY_TOKEN, token);
		}
	}
	function signOut() {
		is_login.value = false;
		is_admin.value = false;
		info.value = resetInfo();
		window.localStorage.removeItem(consts.WINDOW_LOCALSTORAGE_KEY_TOKEN);
	}
	return {
		is_login,
		is_admin,
		info,
		signIn,
		signOut,
	};
});
