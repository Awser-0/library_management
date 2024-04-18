<template>
	<div class="borrow-manage">
		<BorrowRecordTable audiences="admin" :data="records" @update="queryRecords" />
	</div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { BorrowRecordTable } from "~/components";
import { bookApi } from "~/apis";

const records = ref<bookApi.BorrowRecord[]>([]);

async function queryRecords() {
	await bookApi
		.queryRecords()
		.then(({ data: result }) => {
			if (result.code == 10200) {
				records.value = result.data.records;
				ElMessage.success("请求成功");
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("请求失败");
		});
}

onMounted(() => {
	queryRecords();
});
</script>

<style scoped>
.borrow-manage {
	font-size: inherit;
}
</style>
