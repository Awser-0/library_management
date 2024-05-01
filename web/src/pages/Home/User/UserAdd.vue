<template>
	<UserFormVue
		title="添加用户"
		:visible="props.visible"
		:should-password="true"
		@close="emits('close')"
		@submit="submit"
	/>
</template>

<script setup lang="ts">
import UserFormVue from "./UserForm.vue";
import { ref, watch } from "vue";
import { userApi } from "~/apis";

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
	username: string;
	password: string;
	nickname: string;
	sex: "0" | "1";
	phone: string;
	birth: Date | null;
	isAdmin: boolean;
}) {
	if (isLoading.value) return;
	isLoading.value = true;
	return userApi
		.register(form)
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
