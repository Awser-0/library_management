<template>
	<div class="book-page">
		<ul class="books">
			<li v-for="book in books" :key="book.uuid">
				<div class="cover"><img :src="book.cover" /></div>
				<div class="info">
					<div class="title">{{ book.title }}2</div>
					<div class="author">{{ book.author }}</div>
					<div class="introduction">{{ book.introduction }}>2</div>
				</div>
			</li>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
		</ul>
	</div>
</template>

<script setup lang="ts">
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

onMounted(() => {
	bookApi.queryBooks("").then(({ data: result }) => {
		if (result.code == 10200) {
			books.value = result.data.books.map((item) => {
				return {
					uuid: item.uuid,
					title: item.title,
					author: item.author,
					cover: item.cover,
					introduction: item.introduction,
				};
			});
		}
	});
});
</script>

<style>
.book-page {
	min-width: 900px;
	height: 2000px;
}
.books {
	display: flex;
	justify-content: space-between;
	flex-wrap: wrap;

	li {
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
			height: 220px;
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
