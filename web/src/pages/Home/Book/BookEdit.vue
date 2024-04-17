<template>
	<BookFormVue
		title="修改书籍"
		:visible="props.visible"
		:book-title="bookInfo.title"
		:book-author="bookInfo.author"
		:book-cover="bookInfo.cover"
		:book-introduction="bookInfo.introduction"
		:book-publish-time="bookInfo.publish_time"
		:book-total="bookInfo.total"
		@close="emits('close')"
		@submit="submit"
	/>
</template>

<script setup lang="ts">
import BookFormVue from "./BookForm.vue";
import { ref, reactive, watch } from "vue";
import { bookApi } from "~/apis";

const props = defineProps<{
	visible: boolean;
	uuid: number;
}>();

const emits = defineEmits<{
	(e: "close"): void;
	(e: "editAfter"): void;
}>();

const dialogVisible = ref(props.visible);

let isLoading = ref(false);

const bookInfo = reactive({
	title: "",
	author: "",
	cover: "",
	introduction: "",
	publish_time: "",
	total: 0,
});

async function selectBook() {
	if (isLoading.value) return;
	isLoading.value = true;
	return await bookApi
		.selectBook(props.uuid)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				const book = result.data.book;
				const cover = book.cover != "" ? import.meta.env.VITE_IMAGE_BASE_URL + book.cover : "";
				bookInfo.title = book.title;
				bookInfo.author = book.author;
				bookInfo.cover = cover;
				bookInfo.introduction = book.introduction;
				bookInfo.publish_time = book.publish_time || "";
				bookInfo.total = book.total;
				return true;
			} else {
				ElMessage.error(result.msg);
				return false;
			}
		})
		.catch(() => {
			ElMessage.error("请求失败");
			return false;
		})
		.finally(() => {
			isLoading.value = false;
		});
}

async function submit(form: {
	title: string;
	author: string;
	cover: string;
	introduction: string;
	publish_time: Date | null;
	total: number;
}) {
	if (isLoading.value) return;
	isLoading.value = true;
	return bookApi
		.updateBook(props.uuid, form)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				emits("close");
				ElMessage.success("修改成功");
				emits("editAfter");
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("修改失败");
		})
		.finally(() => {
			isLoading.value = false;
		});
}

watch(
	() => props.visible,
	(value) => {
		dialogVisible.value = value;
		if (value == true) {
			selectBook();
		}
	},
);
</script>
