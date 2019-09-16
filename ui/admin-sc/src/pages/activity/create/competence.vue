<template>
	<div class="create-activity-competence">
		<div class="box">
			<h2 class="box-header">
				Select badge to give
			</h2>
			<div class="box-content">
				<div class="my-row">
					<b-col
						v-for="(option, id) in options"
						:key="`${id}`"
						lg="3"
						cols="6"
						class="badge-wrapper"
					>
						<label
							:for="`${id}`"
							:class="[
								'badge-checkbox',
								'is-select'
							]"
						>
							<base-image size="90" />
							<input
								:id="`${id}`"
								:value="option"
								type="checkbox"
								@input="removeError(index)"
							>
							<p class="text">{{ option.name }}</p>
						</label>
					</b-col>
				</div>
			</div>
		</div>
		<base-page-step
			:step="step"
			@next="submit"
			@back="goBack"
		/>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/create-activity-competence.scss";
</style>
<script>
import { mapState } from "vuex";

export default {
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			if (!vm.steps.includes("detail")) {
				vm.$router.replace({ name: "create-activity" });
			}
		});
	},
	data() {
		return {
			errors: [],
			options: [ // TODO: Get all badge options from backend
				{
					id: "002",
					name: "Team working"
				},
				{
					id: "003",
					name: "Communication"
				},
				{
					id: "004",
					name: "Leadership"
				},
				{
					id: "005",
					name: "Flexible"
				}
			]
		};
	},
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
			this.$router.push({ name: "create-activity-confirmation" });
		},
		goBack() {
			this.$router.push({ name: "create-activity" });
		}
	}
};
</script>