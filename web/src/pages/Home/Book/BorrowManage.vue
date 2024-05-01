<template>
	<div class="borrow-manage">
		<BorrowRecordTable audiences="admin" :data="dataPage.list" @update="queryRecords" />
		<div style="height: 10px"></div>
		<el-pagination
			background
			layout="total, prev, pager, next, jumper"
			:total="dataPage.total"
			:page-size="dataPage.page_size"
			@update:current-page="onUpdatePageNum"
		/>
	</div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { BorrowRecordTable } from "~/components";
import { bookApi } from "~/apis";
import zhCn from "element-plus/es/locale/lang/zh-cn";

const records = ref<bookApi.BorrowRecord[]>([]);

const dataPage = reactive({
	list: [] as bookApi.BorrowRecord[],
	page_num: 1,
	page_size: 8,
	total: 0,
});

async function onUpdatePageNum(val: number) {
	console.log("val", val);
	dataPage.page_num = val;
	await queryRecords();
}

async function queryRecords() {
	await bookApi
		.queryRecords({ page_num: dataPage.page_num, page_size: dataPage.page_size })
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
	queryRecords();
});
</script>

<style scoped>
.borrow-manage {
	display: flex;
	flex-direction: column;
	align-items: center;
}
</style>
