<template lang="html">
	<div
		v-if="loading"
		class="loading"
	>
		<b-spinner
			class="spin"
			variant="primary"
			label="Spinning"
		/>
	</div>
</template>
<style lang="scss" scoped>
@import "@/styles/components/loading.scss";
</style>
<script lang="ts">
import Vue from "vue";

export default Vue.extend({
	data() {
		return {
			loading: false
		};
	},
	methods: {
		start(): void {
			this.loading = true;
		},
		finish(): void {
			this.loading = false;
		},
		fail(): void {
			this.loading = false;
			(this as any).$bvToast.toast("Toast body content", {
				title: "Something Error",
				variant: "danger",
				solid: true
			});
		}
	}
});
</script>