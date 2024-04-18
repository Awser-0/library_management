<template>
	<el-table :data="props.data" style="width: 100%">
		<el-table-column prop="book_uuid" label="书籍ID" width="180" />
		<el-table-column prop="user_id" label="Name" width="180" />
		<el-table-column prop="apply_description" label="申请理由" width="180" />
		<el-table-column prop="reply_description" label="回复理由" width="180" />
		<el-table-column label="归还时间" width="180">
			<template #default="scope">{{ utils.dayjsFormatDate(scope.row["return_time"]) }}</template>
		</el-table-column>
		<el-table-column label="操作" width="200">
			<template #default="scope">
				<div v-if="scope.row['state'] == BorrowRecordState.Apply">
					<template v-if="props.audiences == 'admin'">
						<el-button type="success" @click="showDescDialog(scope.row['id'], 'agree')">
							同意
						</el-button>
						<el-button type="danger" @click="showDescDialog(scope.row['id'], 'reject')">
							拒绝
						</el-button>
					</template>
					<template v-else>
						<el-button type="info" @click="cancel(scope.row['id'])">取消</el-button>
					</template>
				</div>
				<div v-if="scope.row['state'] == BorrowRecordState.Borrow">
					<template v-if="props.audiences == 'admin'">
						<el-button type="success" @click="returnBook(scope.row['id'])">归还</el-button>
					</template>
					<template v-else>借阅中</template>
				</div>
				<div v-if="scope.row['state'] == BorrowRecordState.Cancel">已取消</div>
				<div v-if="scope.row['state'] == BorrowRecordState.Reject">已拒绝</div>
				<div v-if="scope.row['state'] == BorrowRecordState.Return">已归还</div>
			</template>
		</el-table-column>
	</el-table>
	<DescDialog
		:visible="dialogVisible"
		title="回复"
		placeholder="请填写回复"
		@close="dialogVisible = false"
		@confirm="confirm"
	/>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { DescDialog } from "~/components";
import { bookApi } from "~/apis";
import { BorrowRecordState } from "~/consts";
import * as stores from "~/stores";
import * as utils from "~/utils";

const userStore = stores.useUserStore();

const props = defineProps<{
	audiences: "user" | "admin";
	data: bookApi.BorrowRecord[];
}>();
const emits = defineEmits<{
	(e: "update"): void;
}>();

const operateForm = reactive({
	record_id: -1,
	purpose: "" as "agree" | "reject",
});
const dialogVisible = ref(false);

function showDescDialog(record_id: number, purpose: "agree" | "reject") {
	operateForm.record_id = record_id;
	operateForm.purpose = purpose;
	dialogVisible.value = true;
}

async function cancel(record_id: number) {
	await bookApi
		.cancelBorrowBook(record_id)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				ElMessage.success("请求成功");
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("请求失败");
		});
	emits("update");
}

async function agree(record_id: number, desc: string) {
	await bookApi
		.agreeBorrowBook(record_id, desc)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				ElMessage.success("请求成功");
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("请求失败");
		});
	emits("update");
}

async function reject(record_id: number, desc: string) {
	await bookApi
		.rejectBorrowBook(record_id, desc)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				ElMessage.success("请求成功");
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("请求失败");
		});
	emits("update");
}
async function returnBook(record_id: number) {
	await bookApi
		.returnBorrowBook(record_id)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				ElMessage.success("请求成功");
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("请求失败");
		});
	emits("update");
}

async function confirm(desc: string) {
	if (operateForm.purpose == "agree") {
		await agree(operateForm.record_id, desc);
	} else {
		await reject(operateForm.record_id, desc);
	}
	dialogVisible.value = false;
}
</script>

<style scoped></style>
