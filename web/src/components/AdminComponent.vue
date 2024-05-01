<template>
	<slot v-if="userStore.is_admin"></slot>
</template>

<script setup lang="ts">
import * as stores from "~/stores";
const userStore = stores.useUserStore();
</script>
