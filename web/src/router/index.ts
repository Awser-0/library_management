export { router, RouterName };

import { createRouter, createWebHashHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import * as pages from "~/pages";
import * as stores from "~/stores";
import * as consts from "~/consts";
import { userApi } from "~/apis";

const RouterName = {
	Login: "Login",
	Home: "Home",
	StatisPage: "StatisPage",
	BookPage: "BookPage",
};

const routes: RouteRecordRaw[] = [
	{
		path: "/",
		name: RouterName.Home,
		component: pages.Home,
		children: [
			{
				path: "book",
				component: pages.HomeVues.Book,
				name: RouterName.BookPage,
			},
			{
				path: "statis",
				component: pages.HomeVues.Statis,
				name: RouterName.StatisPage,
			},
		],
	},
	{
		path: "/login",
		name: RouterName.Login,
		component: pages.Login,
	},
];

const router = createRouter({
	history: createWebHashHistory(),
	routes,
});

async function autoLogin() {
	const userStore = stores.userUserStore();
	const token = window.localStorage.getItem(consts.WINDOW_LOCALSTORAGE_KEY_TOKEN);
	if (token) {
		return await userApi
			.loginByToken(token)
			.then(({ data: result }) => {
				if (result.code == 10200) {
					const { user, token } = result.data;
					userStore.signIn({ ...user }, token);
					router.push({ name: RouterName.Home });
					return true;
				}
				return false;
			})
			.catch(() => {
				return false;
			});
	}
	return false;
}

router.beforeEach(async (to, from) => {
	const userStore = stores.userUserStore();
	if (to.name != RouterName.Login && !userStore.isLogin) {
		if (await autoLogin()) return true;
		else return { name: RouterName.Login };
	}
	return true;
});
