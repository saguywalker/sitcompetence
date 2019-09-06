<template>
	<div>
		<base-page-step
			:step="step"
			@next="submit"
			@back="goBack"
		/>
	</div>
</template>
<script>
import BasePageStep from "@/components/BasePageStep.vue";

export default {
	components: {
		BasePageStep
	},
	computed: {
		step() {
			return this.$route.meta.step;
		}
	},
	methods: {
		submit() {
			this.$router.push({ name: "give-badge-success" });
		},
		goBack() {
			this.$router.push({ name: "give-badge-selection" });
		}
	}
};
</script>