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
<style lang="scss">
@import "@/styles/components/loading.scss";
</style>
<script>
export default {
	data() {
		return {
			loading: false
		};
	},
	methods: {
		start() {
			this.loading = true;
		},
		finish() {
			this.loading = false;
		},
		fail() {
			this.loading = false;
			this.$bvToast.toast("Toast body content", {
				title: "Something Error",
				variant: "danger",
				solid: true
			});
		}
	}
};
</script>