<template>
	<el-dialog
		width="700"
		align-center
		v-model="dialogVisible"
		:title="props.title"
		:before-close="beforeClose"
	>
		<el-form class="form" label-width="auto" ref="formRef" :model="form" :rules="rules">
			<div class="form_top">
				<div class="form_top_left">
					<el-form-item prop="cover">
						<input
							type="file"
							v-show="false"
							accept=".png,.jpg,.jpeg"
							ref="uploadImageRef"
							@change="uploadImage"
						/>
						<div class="cover" @click="uploadImageRefClick">
							<el-image :src="form.cover">
								<template #error>
									<div class="image-slot">
										<el-icon><Plus /></el-icon>
									</div>
								</template>
							</el-image>
						</div>
					</el-form-item>
				</div>
				<div class="form_top_right">
					<el-form-item label="书名" prop="title">
						<el-input v-model="form.title" placeholder="请输入书名" />
					</el-form-item>
					<el-form-item label="作者" prop="author">
						<el-input v-model="form.author" placeholder="请输入作者" />
					</el-form-item>
					<el-form-item label="发布日期">
						<el-date-picker v-model="form.publish_time" type="month" placeholder="选择发布日期" />
					</el-form-item>
					<el-form-item label="总数">
						<el-input-number v-model="form.total" />
					</el-form-item>
				</div>
			</div>
			<el-form-item label="简介">
				<el-input v-model="form.introduction" type="textarea" resize="none" :rows="4" />
			</el-form-item>
		</el-form>
		<template #footer>
			<div class="dialog-footer">
				<el-button>取消</el-button>
				<el-button type="primary" @click="submitForm(formRef)">提交</el-button>
			</div>
		</template>
	</el-dialog>
</template>

<script lang="ts" setup>
import { ref, reactive, watch } from "vue";
import { ElMessageBox } from "element-plus";
import { Plus } from "@element-plus/icons-vue";
import type { FormInstance, FormRules } from "element-plus";
import { fileApi } from "~/apis";

const uploadImageRef = ref<HTMLInputElement>();

function uploadImageRefClick() {
	uploadImageRef.value?.click();
}

const props = defineProps<{
	title?: string;
	visible: boolean;
}>();
const emits = defineEmits<{
	(e: "close"): void;
	(
		e: "submit",
		form: {
			title: string;
			author: string;
			cover: string;
			introduction: string;
			publish_time: Date | null;
			total: number;
		},
	): void;
}>();

const formRef = ref<FormInstance>();

const dialogVisible = ref(props.visible);

// const handleChange = (value: number) => {
// 	console.log(value);
// };
type Form = {
	title: string;
	author: string;
	cover: string;
	introduction: string;
	publish_time: string;
	total: number;
};
// import.meta.env.VITE_IMAGE_BASE_URL + "1712754658244300.png"
const form = reactive<Form>({
	title: "活着",
	author: "余华",
	cover: "",
	introduction: "简介 嘿嘿嘿嘿",
	publish_time: "",
	total: 30,
});

const rules = reactive<FormRules<Form>>({
	title: [{ required: true, message: "请输入书名", trigger: "blur" }],
	author: [{ required: true, message: "请输入作者", trigger: "blur" }],
});

const submitForm = async (formEl: FormInstance | undefined) => {
	if (!formEl) return;
	await formEl.validate((valid, fields) => {
		if (!valid) return console.log("error submit!", fields);
		emits("submit", {
			title: form.title,
			author: form.author,
			cover: form.cover,
			introduction: form.introduction,
			publish_time: form.publish_time != "" ? new Date(form.publish_time + "UTC") : null,
			total: form.total || 0,
		});
	});
};

async function uploadImage(evt: Event) {
	const file = (<any>evt).target.files[0] as File | null;
	if (file == null) return;
	await fileApi
		.uploadImage(file)
		.then(({ data: result }) => {
			if (result.code == 10200) {
				form.cover = import.meta.env.VITE_IMAGE_BASE_URL + result.data.name;
				ElMessage.error("上传成功");
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("上传失败");
		});
}

watch(
	() => props.visible,
	(value) => {
		dialogVisible.value = value;
	},
);

const beforeClose = (_done: () => void) => {
	emits("close");
};
</script>

<style scoped>
.form_top {
	display: flex;
	.form_top_left {
		flex-shrink: 0;
		margin-right: 20px;
	}
	.form_top_right {
		flex-grow: 1;
	}
}
.cover {
	width: 160px;
	height: 200px;
}
.el-image {
	padding: 0 5px;
	width: 100%;
	height: 100%;
}

.image-slot {
	display: flex;
	justify-content: center;
	align-items: center;
	width: 100%;
	height: 100%;
	background: var(--el-fill-color-light);
	color: var(--el-text-color-secondary);
	font-size: 30px;
}
.image-slot .el-icon {
	font-size: 30px;
}
</style>
