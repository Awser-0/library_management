<template>
	<el-table :data="props.data" style="width: 100%">
		<el-table-column prop="book_uuid" label="书籍ID" width="180" />
		<el-table-column prop="user_id" label="Name" width="180" />
		<el-table-column label="操作">
			<template #default="scope">
				<div v-if="scope.row['state'] == BorrowRecordState.Apply">
					<template v-if="props.audiences == 'admin'">
						<el-button type="success">同意</el-button>
						<el-button type="danger">拒绝</el-button>
					</template>
					<template v-else>
						<el-button type="info">取消</el-button>
					</template>
				</div>
				<div v-if="scope.row['state'] == BorrowRecordState.Borrow">
					<template v-if="props.audiences == 'admin'">
						<el-button type="success">归还</el-button>
					</template>
					<template v-else>
						<el-button type="success">借阅中</el-button>
					</template>
				</div>
				<div v-if="scope.row['state'] == BorrowRecordState.Cancel">已取消</div>
				<div v-if="scope.row['state'] == BorrowRecordState.Reject">已拒绝</div>
				<div v-if="scope.row['state'] == BorrowRecordState.Return">已归还</div>
			</template>
		</el-table-column>
	</el-table>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { bookApi } from "~/apis";
import { BorrowRecordState } from "~/consts";
import * as stores from "~/stores";

const userStore = stores.useUserStore();

const props = defineProps<{
	audiences: "user" | "admin";
	data: bookApi.BorrowRecord[];
}>();
</script>

<style scoped></style>
