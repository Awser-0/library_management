<template>
	<div class="user-page">this is user-page</div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { userApi } from "~/apis";
async function queryUser(query_string: string = "") {
	await userApi
		.queryUser(query_string)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				console.log(result.data);
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
