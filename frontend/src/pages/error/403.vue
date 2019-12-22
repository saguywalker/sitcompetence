<template>
	<div class="error-404">
		<div class="content">
			<h1 class="title">
				403 - Forbidden
			</h1>
			<p class="description">
				If you got this status error, we are sorry that we going to force you to login again.
			</p>
			<b-button
				variant="primary"
				@click="backToLogin"
			>
				Back to login
			</b-button>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/error/404.scss";
</style>
<script>
import { clearLoginState } from "@/helpers";

export default {
	beforeRouteLeave(to, from, next) {
		this.clearLoginState();
		next();
	},
	methods: {
		clearLoginState,
		backToLogin() {
			this.clearLoginState();
			this.$router.push({ name: "login" });
		}
	}
};
</script>