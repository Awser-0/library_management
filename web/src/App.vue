<template>
	<div class="app">
		<LoginPage />
	</div>
</template>

<script setup lang="ts">
import { onMounted } from "vue";
import { userApi } from "./apis";
import LoginPage from "~/page/LoginPage/LoginPage.vue";

let a: number = 10;

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
	background: blue;
}
</style>
