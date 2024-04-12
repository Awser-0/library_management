<template>
	<div class="book-page">
		<div class="book-operate">
			<el-button type="primary" size="large" @click="bookAddVisible = true">添加</el-button>
		</div>
		<ul class="book-list">
			<li class="book" v-for="book in books" :key="book.uuid">
				<div class="cover"><img :src="book.cover" /></div>
				<div class="info">
					<div class="title">{{ book.title }}</div>
					<div class="author">作者：{{ book.author }}</div>
					<div class="introduction">简介：{{ book.introduction }}</div>
				</div>
			</li>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
		</ul>
		<BookAddVue :visible="bookAddVisible" @close="bookAddVisible = false" />
	</div>
</template>

<script setup lang="ts">
import BookAddVue from "./BookAdd.vue";
import { ref, reactive, onMounted } from "vue";
import type { FormInstance, FormRules } from "element-plus";
import * as stores from "~/stores";
import { bookApi } from "~/apis";
import { router, RouterName } from "~/router";

type BookInfo = {
	uuid: number;
	title: string;
	author: string;
	cover: string;
	introduction: string;
};

const books = ref<BookInfo[]>([]);
const bookAddVisible = ref(false);
bookAddVisible.value = true;

onMounted(() => {
	bookApi.queryBooks("").then(({ data: result }) => {
		if (result.code == 10200) {
			books.value = result.data.books.map((item) => {
				const cover = item.cover != "" ? import.meta.env.VITE_IMAGE_BASE_URL + item.cover : "";
				return {
					uuid: item.uuid,
					title: item.title,
					author: item.author,
					cover: cover,
					introduction: item.introduction,
				};
			});
		}
	});
});
</script>

<style scoped>
.book-page {
	min-width: 900px;
	height: 2000px;
}
.book-operate {
	text-align: right;
}
.book-operate :deep(.el-button) {
	text-align: right;
	font-size: 16px;
}
.book-list {
	display: flex;
	justify-content: space-between;
	flex-wrap: wrap;

	.book {
		display: flex;
		width: 460px;
		border-radius: 20px;
		margin: 30px;
		box-sizing: border-box;
		padding: 20px;
		background: #ffffff;
		flex-grow: 1;
		.cover {
			flex-shrink: 0;
			width: 160px;
			height: 200px;
			margin-right: 20px;
			background: #cdcdcd;
			img {
				width: 100%;
				height: 100%;
			}
		}
		.info {
			/* width: calc(100% - 160px- 20px - 780px); */
			flex-grow: 1;
			font-size: 18px;
			.title {
				font-weight: bold;
				font-size: 22px;
				line-height: 1.8em;
				white-space: pre-wrap;
				word-break: break-all;
				display: -webkit-box;
				overflow: hidden;
				-webkit-line-clamp: 2;
				-webkit-box-orient: vertical;
			}
			.author {
				line-height: 1.5em;
				white-space: pre-wrap;
				word-break: break-all;
				display: -webkit-box;
				overflow: hidden;
				-webkit-line-clamp: 2;
				-webkit-box-orient: vertical;
			}
			.introduction {
				padding: 4px 0;
				line-height: 1.2em;
				white-space: pre-wrap;
				word-break: break-all;
				display: -webkit-box;
				overflow: hidden;
				-webkit-line-clamp: 4;
				-webkit-box-orient: vertical;
			}
		}
	}
	i {
		width: 460px;
		flex-grow: 1;
		margin: 30px;
	}
}
</style>
