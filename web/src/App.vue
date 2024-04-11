<template>
	<div class="app">
		<RouterView />
	</div>
</template>

<script setup lang="ts">
import { onMounted, onBeforeMount } from "vue";
import * as consts from "~/consts";
import * as stores from "~/stores";
import { userApi } from "~/apis";
import { router, RouterName } from "~/router";

const userStore = stores.userUserStore();

let a: number = 10;
onBeforeMount(() => {});
onMounted(() => {
	userApi
		.login("u1", "1234")
		.then(({ data: result }) => {
			console.log("result", result);
		})
		.catch((err) => {
			console.log("err", err);
		});
});
</script>

<style scoped>
.app {
	min-height: 100vh;
}
</style>
