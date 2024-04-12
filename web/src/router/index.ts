export { router, RouterName };

import { createRouter, createWebHashHistory } from "vue-router";
import type { RouteRecordRaw, RouteRecordName } from "vue-router";
import * as pages from "~/pages";
import * as stores from "~/stores";
import * as consts from "~/consts";
import { userApi } from "~/apis";

const RouterName = {
	Login: "Login",
	Home: "Home",
	StatisPage: "StatisPage",
	BookPage: "BookPage",
	HomeIndex: "HomeIndex",
};

const routes: RouteRecordRaw[] = [
	{
		path: "/",
		name: RouterName.Home,
		component: pages.Home,
		meta: { isMustLogin: true },
		children: [
			{
				path: "",
				name: RouterName.HomeIndex,
				redirect: "book",
			},
			{
				path: "book",
				name: RouterName.BookPage,
				component: pages.HomeVues.Book,
			},
			{
				path: "statis",
				name: RouterName.StatisPage,
				component: pages.HomeVues.Statis,
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
let wantToName: RouteRecordName | null | undefined = "";
router.beforeEach(async (to, from) => {
	console.log("to", to);
	console.log("from", from);
	console.log(to.name);
	console.log(to.meta);

	const userStore = stores.userUserStore();
	if (to.meta.isMustLogin && !userStore.isLogin) {
		wantToName = to.name;
		if (await autoLogin()) return { name: wantToName as string };
		else return { name: RouterName.Login };
	}
	// await autoLogin();
	return true;
});
