<template>
	<div class="give-badge-selection">
		<div class="box dropdown">
			<h2 class="box-header">
				Select badge to give
			</h2>
			<div class="box-content">
				<ul class="selected-student">
					<li
						v-for="(item, index) in selectStudent"
						:key="`${item}${index}`"
						:class="{'error': errors[index]}"
					>
						<a
							:class="[
								'dropdown',
								item.show ? 'active' : ''
							]"
							@click="item.show = !item.show"
						>
							<p>{{ item.student_id }} {{ item.first_name }} {{ item.last_name }}</p>
							<icon-arrow-dropdown class="icon" />
						</a>
						<transition name="slide-down">
							<div
								v-if="item.show"
								class="badge-form"
							>
								<div class="my-row">
									<b-col
										v-for="(option, id) in options"
										:key="`${item.student_id}${id}`"
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:for="`${item.student_id}${id}`"
											:class="[
												'badge-checkbox',
												hasSelected(item.badges, option.competence_id) ? 'is-select' : '',
												hasCollected(item.collected_competence, option.competence_id) ? 'is-collected' : ''
											]"
										>
											<base-image size="90" />
											<input
												:id="`${item.student_id}${id}`"
												v-model="item.badges"
												:value="option"
												type="checkbox"
												@input="removeError(index)"
											>
											<p class="text">{{ option.competence_name }}</p>
										</label>
									</b-col>
								</div>
							</div>
						</transition>
					</li>
				</ul>
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
@import "@/styles/pages/give-badge-selection.scss";
</style>
<script>
import loading from "@/plugin/loading";
import IconArrowDropdown from "@/components/icons/IconArrowDropdown.vue";
import { mapState } from "vuex";

export default {
	components: {
		IconArrowDropdown
	},
	data() {
		return {
			selectStudent: [],
			errors: [],
			options: []
		};
	},
	computed: {
		...mapState("base", [
			"badges"
		]),
		...mapState("giveBadge", [
			"selectedStudents",
			"steps"
		]),
		step() {
			return this.$route.meta.step;
		},
		hasError() {
			return this.errors.some((error) => error);
		}
	},
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			if (!vm.steps.includes("main")) {
				vm.$router.replace({ name: "give-badge" });
			}
		});
	},
	async created() {
		if (this.badges.length === 0) {
			loading.start();
			try {
				await this.$store.dispatch("base/loadBadgeData");
			} catch (err) {
				this.$bvToast.toast(`There was a problem fetching badges data: ${err.message}`, {
					title: "Badge Data Error",
					variant: "danger",
					autoHideDelay: 1500
				});
			} finally {
				loading.stop();
			}
		}
		this.options = this.badges;
		this.selectStudent = this.selectedStudents;
	},
	methods: {
		validateSubmit() {
			this.errors = [];

			this.selectStudent.forEach((student) => {
				if (student.badges.length === 0) {
					this.errors.push(true);
				} else {
					this.errors.push(false);
				}
			});
		},
		async submit() {
			this.validateSubmit();

			if (this.hasError) {
				this.$bvToast.toast("Please select the badge to student", {
					title: "No badge error",
					variant: "danger",
					autoHideDelay: 1500
				});
				return;
			}

			await this.$store.dispatch("giveBadge/updateSelectedBadge", this.selectStudent);
			await this.$store.dispatch("giveBadge/addStep", this.step.step);
			this.$router.push({ name: "give-badge-confirmation" });
		},
		async goBack() {
			await this.$store.dispatch("giveBadge/deleteStep", this.step.step);
			this.$router.push({ name: "give-badge" });
		},
		removeError(index) {
			this.errors[index] = false;
		},
		hasSelected(badges, id) {
			return badges.some((badge) => badge.competence_id === id);
		},
		hasCollected(collected, id) {
			if (!collected) {
				return false;
			}
			return collected.some((collectedBadge) => collectedBadge.competence_id === id);
		}
	}
};
</script>