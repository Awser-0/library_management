<template>
	<div class="my-borrow">
		<BorrowRecordTable audiences="user" :data="records" @update="querySelfRecords" />
	</div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { BorrowRecordTable } from "~/components";
import { bookApi } from "~/apis";

const records = ref<bookApi.BorrowRecord[]>([]);

async function querySelfRecords() {
	await bookApi
		.querySelfRecords()
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
	querySelfRecords();
});
</script>

<style scoped>
.my-borrow {
	font-size: inherit;
}
</style>
