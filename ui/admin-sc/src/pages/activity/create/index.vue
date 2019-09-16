<template>
	<div>
		<div class="box">
			<h1 class="box-header">
				Insert the activity detail
			</h1>
		</div>
		<base-page-step
			:step="step"
			@next="submit"
		/>
	</div>
</template>
<script>
import { mapState } from "vuex";

export default {
	computed: {
		...mapState("createActivity", [
			"steps"
		]),
		step() {
			return this.$route.meta.step;
		}
	},
	methods: {
		async submit() {
			await this.$store.dispatch("createActivity/addStep", this.step.step);
			this.$router.push({ name: "create-activity-select-competence" });
		}
	}
};
</script>