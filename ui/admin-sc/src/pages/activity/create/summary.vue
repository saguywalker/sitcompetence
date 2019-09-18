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
							Registration time:
						</label>
						<p class="data">
							{{ summary.openRegistTime }} - {{ summary.closeRegistTime }}
						</p>
					</div>
					<div class="detail-row">
						<label class="label">
							Registration date:
						</label>
						<p class="data">
							{{ summary.openRegistDate }}<br> to {{ summary.closeRegistDate }}
						</p>
					</div>
					<div class="detail-row">
						<label class="label">
							Activity time:
						</label>
						<p class="data">
							{{ summary.actStartTime }} - {{ summary.actEndTime }}
						</p>
					</div>
					<div class="detail-row">
						<label class="label">
							Activity date:
						</label>
						<p class="data">
							{{ summary.actStartDate }}
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
import { mapState } from "vuex";

export default {
	// beforeRouteEnter(to, from, next) {
	// 	next((vm) => {
	// 		if (!vm.steps.includes("select")) {
	// 			vm.$router.replace({ name: "create-activity" });
	// 		}
	// 	});
	// },
	components: {
		IconPhoto
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
		goBack() {
			this.$router.push({ name: "create-activity-competence" });
		},
		showPreview() {
			this.$bvModal.show("preview-modal");
		}
	}
};
</script>