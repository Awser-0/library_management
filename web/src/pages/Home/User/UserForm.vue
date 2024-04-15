<template>
	<el-dialog
		width="700"
		align-center
		v-model="dialogVisible"
		:title="props.title"
		:before-close="beforeClose"
	>
		<el-form class="form" label-width="auto" ref="formRef" :model="form" :rules="rules">
			<el-form-item label="账号" prop="username">
				<el-input v-model="form.username" placeholder="请输入账号" />
			</el-form-item>
			<el-form-item label="密码" prop="password" v-if="props.shouldPassword">
				<el-input v-model="form.password" placeholder="请输入密码" />
			</el-form-item>
			<el-form-item label="姓名" prop="name">
				<el-input v-model="form.name" placeholder="请输入姓名" />
			</el-form-item>
			<el-form-item label="性别" prop="sex"></el-form-item>
			<el-form-item label="联系电话" prop="phone">
				<el-input v-model="form.phone" placeholder="请输入联系电话" />
			</el-form-item>
			<el-form-item label="管理员" prop="name"></el-form-item>
		</el-form>
		<template #footer>
			<div class="dialog-footer">
				<el-button @click="emits('close')">取消</el-button>
				<el-button type="primary" @click="submitForm(formRef)">提交</el-button>
			</div>
		</template>
	</el-dialog>
</template>

<script lang="ts" setup>
import { ref, reactive, watch } from "vue";
import { ElMessageBox } from "element-plus";
import { Plus } from "@element-plus/icons-vue";
import type { FormInstance, FormRules } from "element-plus";

const uploadImageRef = ref<HTMLInputElement>();

const props = defineProps<{
	title: string;
	visible: boolean;
	shouldPassword: boolean;
	userUsername?: string;
	userName?: string;
	userSex?: "0" | "1";
	userPhone?: "";
	userBirth?: string;
	userIsAdmin?: boolean;
}>();
const emits = defineEmits<{
	(e: "close"): void;
	(
		e: "submit",
		form: {
			username: string;
			password: string;
			name: string;
			sex: "0" | "1";
			phone: "";
			birth: Date | null;
			isAdmin: boolean;
		},
	): void;
}>();
// 图片点击
function uploadImageRefClick() {
	uploadImageRef.value?.click();
}

const formRef = ref<FormInstance>();

const dialogVisible = ref(props.visible);

// const handleChange = (value: number) => {
// 	console.log(value);
// };
type Form = {
	username: string;
	password: string;
	name: string;
	sex: "0" | "1";
	phone: "";
	birth: string;
	isAdmin: boolean;
};
// import.meta.env.VITE_IMAGE_BASE_URL + "1712754658244300.png"
const form = reactive<Form>({
	username: "",
	password: "",
	name: "",
	sex: "0",
	phone: "",
	birth: "",
	isAdmin: false,
});
const formRules: FormRules<Form> = {
	username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
};
if(props.shouldPassword) {
  formRules['password'] = [{ required: true, message: "请输入密码", trigger: "blur" }],
}
const rules = reactive(formRules);

const submitForm = async (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	await formEl.validate((valid, fields) => {
		if (!valid) return console.log("error submit!", fields);
		emits("submit", {
			username: form.username,
			password: form.password,
			name: form.name,
			sex: form.sex,
			phone: form.phone,
			birth: form.birth != "" ? new Date(form.birth + "UTC") : null,
			isAdmin: form.isAdmin,
		});
	});
};

watch(
	() => props.visible,
	(value) => {
		dialogVisible.value = value;
	},
);

watch(
	props,
	(value) => {
		form.username = value.userUsername ?? form.username;
		form.name = value.userName ?? form.name;
		form.sex = value.userSex ?? form.sex;
		form.phone = value.userPhone ?? form.phone;
		form.birth = value.userBirth ?? form.birth;
		form.isAdmin = value.userIsAdmin ?? form.isAdmin;
	},
	{ immediate: true },
);

const beforeClose = (_done: () => void) => {
	emits("close");
};
</script>

<style scoped></style>
