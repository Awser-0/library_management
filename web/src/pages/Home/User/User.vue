<template>
	<div class="user-page">
		<div class="user-operate">
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
			<el-button type="primary" size="default" @click="userAddForm.visible = true">添加</el-button>
		</div>
		<el-table :data="dataPage.list" style="width: fit-content">
			<el-table-column fixed prop="id" label="ID" />
			<el-table-column label="用户名" width="160">
				<template #default="scope">{{ scope.row["username"] }}</template>
			</el-table-column>
			<el-table-column prop="nickname" label="姓名" width="160" />
			<el-table-column prop="address" label="性别">
				<template #default="scope">{{ scope.row["sex"] == "0" ? "男" : "女" }}</template>
			</el-table-column>
			<el-table-column prop="phone" label="电话" width="160" />
			<el-table-column prop="address" label="管理员">
				<template #default="scope">{{ scope.row["is_admin"] ? "是" : "否" }}</template>
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
			<template #empty>
				<el-empty description="暂无记录" />
			</template>
		</el-table>
		<el-pagination
			class="pagination"
			background
			layout="total, sizes, prev, pager, next, jumper"
			v-model:current-page="dataPage.page_num"
			v-model:page-size="dataPage.page_size"
			:total="dataPage.total"
			:page-size="dataPage.page_size"
			:page-sizes="[5, 8, 10, 20]"
			@update:current-page="onUpdatePageNum"
			@update:page-size="onUpdatePageSize"
		/>
		<UserAddVue
			:visible="userAddForm.visible"
			@close="userAddForm.visible = false"
			@edit-after="queryUser"
			@add-after="query"
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
import { Search } from "@element-plus/icons-vue";

const dataPage = reactive({
	list: [] as userApi.UserInfo[],
	page_num: 1,
	page_size: 8,
	total: 0,
});

const form = reactive({
	queryString: "",
});

async function inputQuery() {
	dataPage.page_num = 1;
	await queryUser(form.queryString);
}

async function query() {
	await queryUser(form.queryString);
}

async function onUpdatePageNum(val: number) {
	query();
}

async function onUpdatePageSize(val: number) {
	query();
}

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
		.queryUser(query_string, { page_num: dataPage.page_num, page_size: dataPage.page_size })
		.then(({ data: result }) => {
			if (result.code == 10200) {
				const data = result.data;
				dataPage.list = data.list.map((item) => {
					return {
						...item,
						create_time: utils.dayjsFormatDate(item.create_time),
					};
				});
				dataPage.page_num = data.page_num;
				dataPage.page_size = data.page_size;
				dataPage.total = data.total;
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
	flex-direction: row;
	justify-content: space-between;
}
.pagination {
	width: fit-content;
	margin: 10px auto;
}
</style>
