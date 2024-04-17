<template>
	<el-dialog
		width="400"
		align-center
		v-model="dialogVisible"
		title="重置密码"
		:before-close="beforeClose"
	>
		<el-form class="form" label-width="auto" ref="formRef" :model="form" :rules="rules">
			<el-form-item label="密码" prop="password">
				<el-input v-model="form.password" placeholder="请输入账号" />
			</el-form-item>
		</el-form>
		<template #footer>
			<div class="dialog-footer">
				<el-button @click="emits('close')">取消</el-button>
				<el-button type="primary" @click="submitForm(formRef)">提交</el-button>
			</div>
		</template>
	</el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, watch } from "vue";
import { userApi } from "~/apis";
import type { FormInstance, FormRules } from "element-plus";

const props = defineProps<{
	visible: boolean;
	userId: number;
}>();
const emits = defineEmits<{
	(e: "close"): void;
}>();
const formRef = ref<FormInstance>();
const dialogVisible = ref(props.visible);

type Form = {
	password: string;
};
const form = reactive<Form>({
	password: "",
});
const formRules: FormRules<Form> = {
	password: [{ required: true, message: "请输入密码", trigger: "blur" }],
};
const rules = reactive(formRules);

async function submitForm(formEl: FormInstance | undefined) {
	if (!formEl) return;
	await formEl.validate((valid, fields) => {
		if (!valid) return console.log("error submit!", fields);
		resetPass(props.userId, form.password);
	});
}

async function resetPass(userId: number, password: string) {
	userApi
		.resetPassword(userId, password)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				ElMessage.success("修改成功");
				emits("close");
				form.password = "";
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("修改失败");
		});
}

watch(
	() => props.visible,
	(value) => {
		dialogVisible.value = value;
	},
);

const beforeClose = (_done: () => void) => {
	emits("close");
	form.password = "";
};
</script>

<style scoped></style>
