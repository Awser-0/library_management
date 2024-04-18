<template>
	<div class="user-page">
		<div class="user-operate">
			<el-button type="primary" size="default" @click="userAddForm.visible = true">添加</el-button>
		</div>
		<el-table :data="users" style="width: fit-content">
			<el-table-column fixed prop="id" label="ID" />
			<el-table-column label="用户名" width="160">
				<template #default="scope">{{ scope.row["username"] }}</template>
			</el-table-column>
			<el-table-column prop="name" label="姓名" width="160" />
			<el-table-column prop="address" label="性别">
				<template #default="scope">{{ scope.row["sex"] == "0" ? "男" : "女" }}</template>
			</el-table-column>
			<el-table-column prop="phone" label="电话" width="160" />
			<el-table-column prop="address" label="管理员">
				<template #default="scope">{{ scope.row["isAdmin"] ? "是" : "否" }}</template>
			</el-table-column>
			<el-table-column prop="create_time" label="创建时间" width="160"></el-table-column>
			<el-table-column fixed="right" label="Operations" width="180">
				<template #default="scope">
					<el-button link type="primary" @click="showUserEdit(scope.row['id'])">修改资料</el-button>
					<el-button link type="primary" @click="showPassReset(scope.row['id'])">
						重置密码
					</el-button>
				</template>
			</el-table-column>
		</el-table>
		<UserAddVue
			:visible="userAddForm.visible"
			@close="userAddForm.visible = false"
			@edit-after="queryUser"
		/>
		<UserEditVue
			:visible="userEditForm.visible"
			:user-id="userEditForm.userId"
			@close="userEditForm.visible = false"
			@edit-after="queryUser"
		/>
		<PassResetVue
			:visible="passResetForm.visible"
			:user-id="passResetForm.userId"
			@close="passResetForm.visible = false"
		/>
	</div>
</template>

<script setup lang="ts">
import UserAddVue from "./UserAdd.vue";
import UserEditVue from "./UserEdit.vue";
import PassResetVue from "./PassReset.vue";
import { ref, reactive, onMounted } from "vue";
import { userApi } from "~/apis";
import * as utils from "~/utils";

const users = ref<userApi.UserInfo[]>([]);
const userAddForm = reactive({
	visible: false,
});
const userEditForm = reactive({
	visible: false,
	userId: -1,
});
const passResetForm = reactive({
	visible: false,
	userId: -1,
});
function showUserEdit(userId: number) {
	userEditForm.userId = userId;
	userEditForm.visible = true;
}
function showPassReset(userId: number) {
	passResetForm.userId = userId;
	passResetForm.visible = true;
}

async function queryUser(query_string: string = "") {
	await userApi
		.queryUser(query_string)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				users.value = result.data.users.map((item) => {
					return {
						...item,
						create_time: utils.dayjsFormatDate(item.create_time),
					};
				});
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("请求失败");
		});
}
onMounted(() => {
	queryUser();
});
</script>

<style scoped>
.user-page {
	width: fit-content;
	margin: 0 auto;
	font-size: inherit;
}
.user-operate {
	margin-bottom: 20px;
	display: flex;
	flex-direction: row-reverse;
}
</style>
