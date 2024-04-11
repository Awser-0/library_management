<template>
	<div class="login">
		<div class="login_left"></div>
		<div class="login_right">
			<el-form
				class="login-form bounce_move_from_right"
				label-width="auto"
				:model="form"
				:rules="rules"
				ref="formRef"
			>
				<h2 class="form-title">登录</h2>
				<el-form-item label="账号" prop="username">
					<el-input placeholder="请输入账号" v-model="form.username" />
				</el-form-item>
				<el-form-item label="密码" prop="password">
					<el-input placeholder="请输入密码" v-model="form.password" />
				</el-form-item>
				<el-form-item>
					<el-button @click="submitForm(formRef)" class="form-login-button" type="primary">
						登录
					</el-button>
				</el-form-item>
			</el-form>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import type { FormInstance, FormRules } from "element-plus";
import * as stores from "~/stores";
import { userApi } from "~/apis";
import { router, RouterName } from "~/router";

const userStore = stores.userUserStore();

const formRef = ref<FormInstance>();

type Form = {
	username: string;
	password: string;
};

const form = reactive({
	username: "u1",
	password: "1234",
});

const rules = reactive<FormRules<Form>>({
	username: [
		{ required: true, message: "请输入账号", trigger: "blur" },
		// { min: 3, max: 5, message: "Length should be 3 to 5", trigger: "blur" },
	],
	password: [
		{ required: true, message: "请输入密码", trigger: "blur" },
		// { min: 3, max: 5, message: "Length should be 3 to 5", trigger: "blur" },
	],
});

const submitForm = async (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	await formEl.validate((valid, fields) => {
		if (!valid) return console.log("error submit!", fields);
		userApi.login(form.username, form.password).then(({ data: result }) => {
			if (result.code == 10200) {
				const { user, token } = result.data;
				userStore.signIn({ ...user }, token);
				router.push({ name: RouterName.Home });
			} else {
				console.log(result.msg);
			}
		});
	});
};

onMounted(() => {});
</script>

<style scoped>
.login {
	--login-left-bg-color: rgb(40, 49, 73);
	--login-right-bg-color: #ffffff;
	min-height: 100vh;
	background: orange;
	display: flex;
	.login_left {
		flex-grow: 1;
		background: var(--login-left-bg-color);
	}
	.login_right {
		flex-grow: 1;
		position: relative;
		background: var(--login-right-bg-color);
	}
	.login-form {
		max-width: 280px;
		height: fit-content;
		position: absolute;
		inset: 0;
		margin: auto;

		.form-title {
			font-size: 22px;
			text-align: center;
			line-height: 2em;
		}
		.form-login-button {
			margin: auto;
		}
	}
	/* .login-form :deep(.el-form-item) {
		margin-bottom: 1.5em;
	} */
}
</style>
