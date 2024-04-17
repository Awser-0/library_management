<template>
	<div class="book-page">
		<div class="book-operate">
			<div class="operate_left">
				<el-input
					v-model="form.queryString"
					style="max-width: 300px"
					placeholder="请输入查询字符"
					class="input-query"
				>
					<template #append>
						<el-button :icon="Search" @click="query" />
					</template>
				</el-input>
			</div>
			<div class="operate_right">
				<el-button type="primary" size="large" @click="bookAddVisible = true">添加</el-button>
			</div>
		</div>
		<ul class="book-list">
			<li class="book" v-for="book in books" :key="book.uuid">
				<div class="cover">
					<img :src="book.cover" />
					<el-icon class="editIcon" @click="showBookEdit(book.uuid)"><Edit /></el-icon>
				</div>
				<div class="info">
					<div class="title">{{ book.title }}</div>
					<div class="author">作者：{{ book.author }}</div>
					<div class="introduction" :title="book.introduction">简介：{{ book.introduction }}</div>
				</div>
				<el-button
					type="success"
					:icon="Check"
					circle
					@click="showBorrowBookDialogForm(book.uuid)"
				/>
			</li>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
			<i></i>
		</ul>
		<BookAddVue :visible="bookAddVisible" @close="bookAddVisible = false" @add-after="queryBooks" />
		<BookEditVue
			:visible="bookEditForm.visible"
			:uuid="bookEditForm.uuid"
			@close="bookEditForm.visible = false"
			@edit-after="queryBooks"
		/>
		<el-dialog v-model="borrowBookDialogForm.visible" title="申请理由" width="400">
			<el-form :model="form">
				<el-form-item>
					<el-input
						v-model="borrowBookDialogForm.desc"
						autocomplete="off"
						placeholder="比如：想看"
					/>
				</el-form-item>
			</el-form>
			<template #footer>
				<div class="dialog-footer">
					<el-button @click="borrowBookDialogForm.visible = false">取消</el-button>
					<el-button type="primary" @click="applyBorrowBook">确定</el-button>
				</div>
			</template>
		</el-dialog>
	</div>
</template>

<script setup lang="ts">
import BookAddVue from "./BookAdd.vue";
import BookEditVue from "./BookEdit.vue";
import { ref, reactive, onMounted } from "vue";
import { ElMessageBox } from "element-plus";
import type { FormInstance, FormRules } from "element-plus";
import { Edit, Search, Check } from "@element-plus/icons-vue";
import * as stores from "~/stores";
import { bookApi } from "~/apis";
import { router, RouterName } from "~/router";

const form = reactive({ queryString: "" });

type BookInfo = {
	uuid: number;
	title: string;
	author: string;
	cover: string;
	introduction: string;
};

const books = ref<BookInfo[]>([]);
const bookAddVisible = ref(false);
// bookAddVisible.value = true;
const bookEditForm = reactive({
	visible: false,
	uuid: 0,
});

function showBookEdit(uuid: number) {
	bookEditForm.uuid = uuid;
	bookEditForm.visible = true;
}

async function query() {
	await queryBooks(form.queryString);
}

async function queryBooks(queryString: string = "") {
	return await bookApi.queryBooks(queryString).then(({ data: result }) => {
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
}

const borrowBookDialogForm = reactive({
	book_uuid: -1,
	visible: false,
	desc: "",
});
function showBorrowBookDialogForm(book_uuid: number) {
	borrowBookDialogForm.book_uuid = book_uuid;
	borrowBookDialogForm.visible = true;
}

async function applyBorrowBook() {
	bookApi
		.applyBorrowBook(borrowBookDialogForm.book_uuid, borrowBookDialogForm.desc)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				ElMessage.success("请求成功");
				borrowBookDialogForm.visible = false;
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("请求失败");
		});
}
onMounted(() => {
	queryBooks();
});
</script>

<style scoped>
.book-page {
	min-width: 900px;
	height: 2000px;
}
.book-operate {
	display: flex;
	justify-content: space-between;
	.operate_left {
		margin-left: 30px;
	}
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
			position: relative;
			margin-right: 20px;
			background: #cdcdcd;
			img {
				width: 100%;
				height: 100%;
			}
			.editIcon {
				width: fit-content;
				position: absolute;
				top: 6px;
				right: 6px;
				margin: 0;
				font-size: 26px;
				cursor: pointer;
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
				margin: 4px 0;
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
