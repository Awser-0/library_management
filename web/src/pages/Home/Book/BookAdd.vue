<template>
	<BookFormVue title="添加书籍" :visible="props.visible" @close="emits('close')" @submit="submit" />
</template>

<script setup lang="ts">
import BookFormVue from "./BookForm.vue";
import { ref, watch } from "vue";
import { bookApi } from "~/apis";

const props = defineProps<{
	visible: boolean;
}>();

const emits = defineEmits<{
	(e: "close"): void;
	(e: "addAfter"): void;
}>();

const dialogVisible = ref(props.visible);

let isLoading = ref(false);

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
		.addBook(form)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				emits("close");
				ElMessage.success("添加成功");
				emits("addAfter");
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("添加失败");
		})
		.finally(() => {
			isLoading.value = false;
		});
}

watch(
	() => props.visible,
	(value) => {
		dialogVisible.value = value;
	},
);
</script>
