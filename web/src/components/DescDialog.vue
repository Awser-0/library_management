<template>
	<el-dialog v-model="dialogVisible" :title="props.title" width="400" :before-close="beforeClose">
		<el-form>
			<el-form-item>
				<el-input v-model="desc" autocomplete="off" :placeholder="props.placeholder || ''" />
			</el-form-item>
		</el-form>
		<template #footer>
			<div class="dialog-footer">
				<el-button @click="emits('close')">取消</el-button>
				<el-button type="primary" @click="confirm">确定</el-button>
			</div>
		</template>
	</el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, watch } from "vue";

const props = defineProps<{
	visible: boolean;
	title: string;
	placeholder?: string;
}>();
const emits = defineEmits<{
	(e: "close"): void;
	(e: "confirm", desc: string): void;
}>();

const dialogVisible = ref(props.visible);
const desc = ref("");

function beforeClose(_done: () => void) {
	emits("close");
}
function confirm() {
	emits("confirm", desc.value);
}
watch(
	() => props.visible,
	(value) => {
		dialogVisible.value = value;
	},
);
</script>

<style scoped>
.root-page {
	font-size: inherit;
}
</style>
