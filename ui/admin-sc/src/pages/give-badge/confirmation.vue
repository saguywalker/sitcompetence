<template>
	<div class="give-badge-selection">
		<div class="box dropdown">
			<h2 class="box-header">
				Confirmation Summary
			</h2>
			<div class="box-content">
				<ul class="selected-student">
					<li
						v-for="(item, index) in selectedStudents"
						:key="`${item}${index}`"
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
								<div class="my-row">
									<b-col
										v-for="(badge, id) in item.badges"
										:key="`${badge}${item.studentId}${id}`"
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:class="[
												'badge-checkbox'
											]"
										>
											<base-image size="90" />
											<p class="text">{{ badge.name }}</p>
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
<script>
import IconArrowDropdown from "@/components/icons/IconArrowDropdown.vue";
import loading from "@/plugin/loading";
import { mapState } from "vuex";

export default {
	components: {
		IconArrowDropdown
	},
	data() {
		return {
			selectStudent: []
		};
	},
	computed: {
		...mapState("giveBadge", [
			"selectedStudents",
			"steps"
		]),
		step() {
			return this.$route.meta.step;
		}
	},
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			if (!vm.steps.includes("selection")) {
				vm.$router.replace({ name: "give-badge" });
			}
		});
	},
	created() {
		this.selectStudent = this.selectedStudents;
	},
	methods: {
		async submit() {
			// TODO: Set giver by login user
			loading.start();

			try {
				await this.$store.dispatch("giveBadge/submitGiveBadge", "tiny");
				await this.$store.dispatch("giveBadge/addStep", this.step.step);
				this.$router.push({ name: "give-badge-success" });
			} catch (err) {
				const notification = {
					title: "Submit give badge",
					message: `There was a problem submitting data: ${err.message}`,
					variant: "danger"
				};

				this.$store.dispatch("notification/add", notification);
			} finally {
				loading.stop();
			}
		},
		goBack() {
			this.$router.push({ name: "give-badge-selection" });
		}
	}
};
</script>