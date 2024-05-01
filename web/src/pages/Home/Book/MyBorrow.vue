<template>
	<div class="my-borrow">
		<BorrowRecordTable audiences="user" :data="dataPage.list" @update="querySelfRecords" />
		<div style="height: 10px"></div>
		<el-pagination
			background
			layout="total, prev, pager, next, jumper"
			:total="dataPage.total"
			:page-size="dataPage.page_size"
			@update:page-size="onUpdatePageSize"
		/>
	</div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { BorrowRecordTable } from "~/components";
import { bookApi } from "~/apis";

const dataPage = reactive({
	list: [] as bookApi.BorrowRecord[],
	page_num: 1,
	page_size: 8,
	total: 0,
});

async function onUpdatePageSize(val: number) {
	dataPage.page_size = val;
	await querySelfRecords();
}

async function querySelfRecords() {
	await bookApi
		.querySelfRecords()
		.then(({ data: result }) => {
			if (result.code == 10200) {
				const data = result.data;
				dataPage.list = data.list;
				console.log("data.list", dataPage.list);
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
	querySelfRecords();
});
</script>

<style scoped>
.my-borrow {
	display: flex;
	flex-direction: column;
	align-items: center;
}
</style>
