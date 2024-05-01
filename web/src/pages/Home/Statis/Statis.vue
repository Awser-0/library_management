<template>
	<div class="myECharts" ref="myEChartsRef"></div>
</template>

<script setup lang="ts">
import { BorrowRecordTable } from "~/components";
import { ref, reactive, onMounted, onUnmounted } from "vue";
import * as echarts from "echarts";
import { bookApi } from "~/apis";

const myEChartsRef = ref<HTMLElement>();
let myChart: echarts.ECharts;

async function initECharts() {
	await bookApi
		.queryRecords({ page_num: 1, page_size: 0 })
		.then(({ data: result }) => {
			const { code, data } = result;
			console.log("data", data);
			if (code == 10200) {
				const groups = data.list.reduce(
					(groups, record) => {
						const key = String(record.book_uuid);
						if (!groups[key]) {
							groups[key] = [];
						}
						groups[key].push(record);
						return groups;
					},
					<{ [k in string]: bookApi.BorrowRecordDetail[] }>{},
				);
				const d = [] as { value: number; name: string }[];
				for (let key in groups) {
					console.log("key", key);
					groups[key];
					d.push({
						name: groups[key][0].book.title,
						value: groups[key].length,
					});
				}
				myChart.setOption({
					title: {
						text: "借阅记录统计",
						// subtext: "Fake Data",
						left: "center",
					},
					tooltip: {
						trigger: "item",
					},
					legend: {
						orient: "vertical",
						left: "left",
					},
					series: [
						{
							name: "借阅次数",
							type: "pie",
							radius: "50%",
							data: d,
							emphasis: {
								itemStyle: {
									shadowBlur: 10,
									shadowOffsetX: 0,
									shadowColor: "rgba(0, 0, 0, 0.5)",
								},
							},
						},
					],
				});
			} else {
				ElMessage.error(result.msg);
			}
		})
		.catch(() => {
			ElMessage.error("请求错误");
		});
}
function windowResize() {
	if (myChart) {
		myChart.resize();
	}
}
onMounted(() => {
	initECharts();
	myChart = echarts.init(myEChartsRef.value);
	window.addEventListener("resize", windowResize);
});
onUnmounted(() => {
	window.removeEventListener("resize", windowResize);
});
</script>

<style scope lang="less">
.myECharts {
	margin: 20px auto;
	height: 500px;
}
</style>
