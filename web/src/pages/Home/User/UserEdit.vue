<template>
	<UserFormVue
		title="修改资料"
		:shouldPassword="false"
		:visible="props.visible"
		:user-username="userInfo.username"
		:user-name="userInfo.name"
		:user-sex="userInfo.sex"
		:user-phone="userInfo.phone"
		:user-birth="userInfo.birth"
		:user-isAdmin="userInfo.isAdmin"
		@close="emits('close')"
		@submit="submit"
	/>
</template>

<script setup lang="ts">
import UserFormVue from "./UserForm.vue";
import { ref, reactive, watch } from "vue";
import { userApi } from "~/apis";
import * as utils from "~/utils";

const props = defineProps<{
	visible: boolean;
	userId: number;
}>();

const emits = defineEmits<{
	(e: "close"): void;
	(e: "editAfter"): void;
}>();

const dialogVisible = ref(props.visible);

let isLoading = ref(false);

const userInfo = reactive({
	username: "",
	name: "",
	sex: "0" as "0" | "1",
	phone: "",
	birth: "",
	isAdmin: false,
});

async function selectUser() {
	if (isLoading.value) return;
	isLoading.value = true;
	return await userApi
		.selectUser(props.userId)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				const user = result.data.user;
				userInfo.username = user.username;
				userInfo.name = user.name;
				userInfo.sex = user.sex;
				userInfo.phone = user.phone;
				userInfo.birth = utils.dayjsFormatDate(user.birth);
				userInfo.isAdmin = user.isAdmin;
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
	username: string;
	password: string;
	name: string;
	sex: "0" | "1";
	phone: string;
	birth: Date | null;
	isAdmin: boolean;
}) {
	if (isLoading.value) return;
	isLoading.value = true;
	return userApi
		.updateOhterUser(props.userId, form)
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
			selectUser();
		}
	},
);
</script>
