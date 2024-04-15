<template>
	<div class="user-page">
		<el-table :data="users" style="width: fit-content">
			<el-table-column fixed prop="id" label="ID" width="150" />
			<el-table-column label="用户名" width="120">
				<template #default="scope">{{ scope.row["username"] }}</template>
			</el-table-column>
			<el-table-column prop="11" label="姓名" width="120" />
			<el-table-column prop="phone" label="电话" width="120" />
			<el-table-column prop="address" label="管理员" width="150">
				<template #default="scope">{{ scope.row["isAdmin"] ? "是" : "否" }}</template>
			</el-table-column>
			<el-table-column prop="create_time" label="创建时间" width="150"></el-table-column>
			<el-table-column fixed="right" label="Operations" width="120">
				<template #default>
					<el-button link type="primary" size="small">详细</el-button>
					<el-button link type="primary" size="small">修改</el-button>
				</template>
			</el-table-column>
		</el-table>
	</div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import dayjs from "dayjs";
import { userApi } from "~/apis";

const users = ref<userApi.UserInfo[]>([]);

async function queryUser(query_string: string = "") {
	await userApi
		.queryUser(query_string)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				console.log(result.data);
				users.value = result.data.users.map((item) => {
					return {
						...item,
						create_time: dayjs(item.create_time).format("YYYY-MM-DD"),
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
	font-size: inherit;
}
</style>
