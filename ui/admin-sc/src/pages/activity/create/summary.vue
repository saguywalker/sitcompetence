<template>
	<div class="create-activity-summary">
		<div class="box">
			<div class="wrapper">
				<div class="header">
					<h2 class="title">
						Activity summary
					</h2>
					<h4 class="subtitle">
						Comfirm activity details
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
							Image:
						</label>
						<p class="data">
							<b-button
								class="preview-btn"
								@click.prevent="showPreview"
							>
								<icon-photo />
							</b-button>
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
						v-model="status"
						name="checkbox-1"
						value="accepted"
						unchecked-value="not_accepted"
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
		<b-modal
			id="preview-modal"
			content-class="preview-image-modal"
			centered
			hide-header
			hide-footer
		>
			<img
				:src="previewImageSrc"
				class="preview-image"
			>
		</b-modal>
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
import IconPhoto from "@/components/icons/IconPhoto.vue";
import IconPen from "@/components/icons/IconPen.vue";
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
		IconPhoto,
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
			status: "",
			summary: {}
		};
	},
	computed: {
		...mapState("createActivity", [
			"detailInput",
			"steps"
		]),
		previewImageSrc() {
			return URL.createObjectURL(this.summary.img);
		},
		step() {
			return this.$route.meta.step;
		}
	},
	created() {
		this.summary = this.detailInput;
	},
	methods: {
		async submit() {
			await this.$store.dispatch("createActivity/addStep", this.step.step);
			this.$router.push({ name: "create-activity-success" });
		},
		async goBack() {
			await this.$store.dispatch("createActivity/deleteStep", this.step.step);
			this.$router.push({ name: "create-activity-competence" });
		},
		showPreview() {
			this.$bvModal.show("preview-modal");
		}
	}
};
</script>