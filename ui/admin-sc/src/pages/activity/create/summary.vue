<template>
	<div class="create-activity-summary">
		<div class="box">
			<div class="wrapper">
				<div class="header">
					<h2 class="title">
						Activity summary
					</h2>
					<h4 class="subtitle">
						Confirm activity details
					</h4>
				</div>
				<router-link
					:to="{ name: 'create-activity' }"
					class="edit"
				>
					<icon-pen /> Edit
				</router-link>
				<div class="detail-wrapper">
					<div class="detail-row">
						<label class="label">
							Activity name:
						</label>
						<p class="data">
							{{ summary.name }}
						</p>
					</div>
					<div class="detail-row">
						<label class="label">
							Activity date:
						</label>
						<p class="data">
							{{ summary.activityDate }}
						</p>
					</div>
					<div class="detail-row description">
						<label class="label">
							Description:
						</label>
						<p class="data">
							{{ summary.description | noValue }}
						</p>
					</div>
				</div>
				<div class="question">
					<b-form-checkbox
						id="checkbox-1"
						v-model="student_site"
						name="checkbox-1"
					>
						Post to student website
					</b-form-checkbox>
					<p class="descript">
						If you not select this activity, it still be shown in activity page.
						<br>
						You can post to student web later.
					</p>
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
@import "@/styles/pages/create-activity-summary.scss";
</style>
<script>
import IconPen from "@/components/icons/IconPen.vue";
import loading from "@/plugin/loading";
import { mapState } from "vuex";

export default {
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			if (!vm.steps.includes("select")) {
				vm.$router.replace({ name: "create-activity" });
			}
		});
	},
	components: {
		IconPen
	},
	filters: {
		noValue(value) {
			if (value === "") {
				return "-";
			}

			return value;
		}
	},
	data() {
		return {
			student_site: false,
			summary: {}
		};
	},
	computed: {
		...mapState("createActivity", [
			"detailInput",
			"steps"
		]),
		step() {
			return this.$route.meta.step;
		}
	},
	created() {
		this.summary = this.detailInput;
	},
	methods: {
		async submit() {
			loading.start();
			try {
				await this.$store.dispatch("createActivity/submitCreateActivity", {
					...this.summary,
					creator: "st01", // TODO: Get from login user
					student_site: this.student_site
				});
				await this.$store.dispatch("createActivity/addStep", this.step.step);
				this.$router.push({ name: "create-activity-success" });
			} catch (err) {
				const notification = {
					title: "Submit create activity",
					message: `There was a problem submitting data: ${err.message}`,
					variant: "danger"
				};

				this.$store.dispatch("notification/add", notification);
			} finally {
				loading.stop();
			}
		},
		async goBack() {
			await this.$store.dispatch("createActivity/deleteStep", this.step.step);
			this.$router.push({ name: "create-activity-competence" });
		}
	}
};
</script>