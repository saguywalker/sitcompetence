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
						:for="`${option.competence_name}${id}`"
						:class="[
							'badge-checkbox',
							hasSelected(option.competence_id) ? 'is-select' : ''
						]"
					>
						<base-image size="90" />
						<input
							:id="`${option.competence_name}${id}`"
							v-model="selects"
							:value="option"
							type="checkbox"
							@input="hasError = false"
						>
						<p class="text">{{ option.competence_name }}</p>
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
			options: [],
			selects: []
		};
	},
	computed: {
		...mapState("base", [
			"badges"
		]),
		...mapState("editActivity", [
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
				vm.$router.replace({ name: "edit-activity" });
			}
		});
	},
	async created() {
		if (this.badges.length === 0) {
			await this.$store.dispatch("base/loadBadgeData");
		}

		this.options = this.badges;
		this.selects = this.competences;
	},
	methods: {
		hasSelected(id) {
			return this.selects.some((select) => select.competence_id === id);
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

			await this.$store.dispatch("editActivity/setCompetenceInput", this.selects);
			await this.$store.dispatch("editActivity/addStep", this.step.step);
			this.$router.push({ name: this.step.next.link });
		},
		async goBack() {
			await this.$store.dispatch("editActivity/deleteStep", this.step.step);
			this.$router.push({ name: this.step.back.link });
		},
		validateSubmit() {
			this.hasError = this.selects.length === 0;
		}
	}
};
</script>