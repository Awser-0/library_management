<template>
	<div class="home">
		<div class="home_left">
			<div class="home_left-title">图书馆</div>
			<ul class="home_left-menu">
				<li v-for="(menu, i) in menuList" :key="i">
					<div class="label">
						<el-icon><Menu /></el-icon>
						<span>{{ menu.label }}</span>
					</div>
					<div class="links">
						<router-link v-for="(link, i) in menu.links" :key="i" :to="{ name: link.linkName }">
							{{ link.label }}
						</router-link>
					</div>
				</li>
			</ul>
		</div>
		<div class="home_right">
			<div class="home_right_header">
				<div class="header_right">
					<div class="user">欢迎你，{{ userStore.info.username }}</div>
					<el-button
						class="signOut"
						type="danger"
						circle
						title="退出登录"
						:icon="Delete"
						@click="signOut"
					/>
				</div>
			</div>
			<div class="home_right_body">
				<router-view />
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import * as stores from "~/stores";
import { userApi } from "~/apis";
import { router, RouterName } from "~/router";
import { Delete, Menu } from "@element-plus/icons-vue";

const userStore = stores.useUserStore();

function signOut() {
	userStore.signOut();
	router.push({ name: RouterName.Login });
}

const menuList = reactive<
	{
		label: string;
		links: {
			label: string;
			linkName: string;
		}[];
	}[]
>([
	{
		label: "书籍",
		links: [
			{
				label: "书籍列表",
				linkName: RouterName.BookPage,
			},
			{
				label: "我的借阅",
				linkName: RouterName.MyBorrow,
			},
		],
	},
	{
		label: "统计",
		links: [
			{
				label: "借阅统计",
				linkName: RouterName.StatisPage,
			},
		],
	},
	{
		label: "用户",
		links: [
			{
				label: "用户管理",
				linkName: RouterName.UserPage,
			},
		],
	},
]);
</script>

<style>
.home {
	--home_left-bg-color: #001529;
	--home-bg-color: #f2f3f5;
	--base-border-color: #dcdfe6;
	/* --menu-color: #ebedf0; */
	/* --menu-color--active: #e6e8eb; */
	--menu-color: #e6e8eb;
	--menu-color--active: #ffffff;
	min-width: fit-content;
	height: 100vh;
	display: flex;
	background: var(--home-bg-color);

	.home_left {
		flex-shrink: 0;
		width: 220px;
		background: var(--home_left-bg-color);
		.home_left-title {
			font-size: 28px;
			color: #fff;
			text-align: center;
			line-height: 2em;
		}
		.home_left-menu {
			font-size: 18px;
			color: var(--menu-color);
			li {
				.label {
					height: 40px;
					display: flex;
					align-items: center;
					padding-left: 20px;
					span {
						margin-left: 4px;
					}
				}
				.links {
					font-size: 16px;
					background: rgb(15, 36, 56);
					a {
						height: 40px;
						display: flex;
						align-items: center;
						padding-left: 40px;
					}
					a:hover {
						color: var(--menu-color--active);
					}
					a.router-link-active {
						color: var(--menu-color--active);
						background: rgb(64, 158, 255);
					}
				}
			}
			li:has(.router-link-active) {
				.label {
					color: var(--menu-color--active);
				}
			}
		}
	}

	.home_right {
		flex-grow: 1;
		.home_right_header {
			height: 60px;
			display: flex;
			border-bottom: 1px solid #dcdfe6;
			background: #ffffff;
			.header_right {
				display: flex;
				justify-content: center;
				align-items: center;
				margin-left: auto;
				.user {
					padding: 4px;
				}
				.signOut {
					margin: 0 20px 0 10px;
				}
			}
		}
		.home_right_body {
			height: calc(100% - 61px);
			box-sizing: border-box;
			padding: 20px;
			overflow-y: auto;
		}
	}
}
</style>
