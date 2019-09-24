<template>
	<div class="create-activity-competence">
		<div
			:class="[
				'box',
				hasError ? 'error' : ''
			]"
		>
			<h2 class="box-header">
				Select badge to give
			</h2>
			<div class="my-row">
				<b-col
					v-for="(option, id) in options"
					:key="`${id}`"
					lg="3"
					cols="6"
					class="badge-wrapper"
				>
					<label
						:for="`${option.name}${id}`"
						:class="[
							'badge-checkbox',
							hasSelected(option.id) ? 'is-select' : ''
						]"
					>
						<base-image size="90" />
						<input
							:id="`${option.name}${id}`"
							v-model="selects"
							:value="option"
							type="checkbox"
							@input="hasError = false"
						>
						<p class="text">{{ option.name }}</p>
					</label>
				</b-col>
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
	data() {
		return {
			hasError: false,
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
			],
			selects: []
		};
	},
	computed: {
		...mapState("createActivity", [
			"competences",
			"steps"
		]),
		step() {
			return this.$route.meta.step;
		}
	},
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			if (!vm.steps.includes("detail")) {
				vm.$router.replace({ name: "create-activity" });
			}
		});
	},
	created() {
		this.selects = this.competences;
	},
	methods: {
		hasSelected(id) {
			return this.selects.some((select) => select.id === id);
		},
		async submit() {
			this.validateSubmit();

			if (this.hasError) {
				this.$bvToast.toast("Please select the badge for the activity", {
					title: "No badge error",
					variant: "danger",
					autoHideDelay: 1500
				});
				return;
			}

			await this.$store.dispatch("createActivity/setCompetenceInput", this.selects);
			await this.$store.dispatch("createActivity/addStep", this.step.step);
			this.$router.push({ name: "create-activity-summary" });
		},
		async goBack() {
			await this.$store.dispatch("createActivity/deleteStep", this.step.step);
			this.$router.push({ name: this.step.back.link });
		},
		validateSubmit() {
			this.hasError = this.selects.length === 0;
		}
	}
};
</script>