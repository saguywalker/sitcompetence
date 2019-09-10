<template>
	<div class="give-badge-selection">
		<div class="box dropdown">
			<h2 class="box-header">
				Select badge to give
			</h2>
			<div class="box-content">
				<ul class="selected-student">
					<li
						v-for="(item, index) in selectedStudents"
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
							<p>{{ item.studentId }} {{ item.fullName }}</p>
							<icon-arrow-dropdown class="icon" />
						</a>
						<transition name="slide-down">
							<div
								v-if="item.show"
								class="badge-form"
							>
								<b-row>
									<b-col
										v-for="(option, id) in options"
										:key="`${option.value}${item.studentId}${id}`"
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:for="`${option.value}${item.studentId}${id}`"
											:class="[
												'badge-checkbox',
												item.badges.includes(option.value) ? 'is-select' : ''
											]"
										>
											<base-image size="90" />
											<input
												:id="`${option.value}${item.studentId}${id}`"
												v-model="item.badges"
												:value="`${option.value}`"
												type="checkbox"
												name="hi"
												@input="removeError(index)"
											>
											<p class="text">{{ option.text }}</p>
										</label>
									</b-col>
								</b-row>
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
import IconArrowDropdown from "@/components/icons/IconArrowDropdown.vue";
import { mapState } from "vuex";

export default {
	components: {
		IconArrowDropdown
	},
	data() {
		return {
			selectStudent: [],
			options: [
				{ text: "Team working", value: "teamWorking" },
				{ text: "Flexible", value: "flexible" },
				{ text: "Leadership", value: "leadership" },
				{ text: "Communication", value: "communication" }
			],
			errors: []
		};
	},
	computed: {
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
	created() {
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
		goBack() {
			this.$router.push({ name: "give-badge" });
		},
		removeError(index) {
			this.errors[index] = false;
		}
	}
};
</script>