<template>
	<div class="activity-approve">
		<section class="section">
			<v-select-page
				v-model="select"
				:data="items"
				:tb-columns="columns"
				:multiple="true"
				title="Select student to approve"
				placeholder="Click here to select"
				key-field="student_id"
				show-field="student_id"
				@values="singleValues"
			/>
			<b-modal
				id="modal-prevent-closing"
				ref="modal"
				title="Submit Your Secret Key"
				@show="resetModal"
				@hidden="resetModal"
				@ok="handleOk"
			>
				<form
					ref="form"
					@submit.stop.prevent="handleSubmit"
				>
					<b-form-group
						:state="skKeyState"
						label="Insert your Secret Key"
						label-for="sk-input"
						invalid-feedback="Secret key is required"
					>
						<b-form-textarea
							id="sk-input"
							v-model="skKey"
							:state="skKeyState"
							rows="6"
							required
							@input="skKeyState = null"
						/>
					</b-form-group>
				</form>
			</b-modal>
			<base-page-step
				:step="step"
				@next="showModal"
				@back="goBack"
			/>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/activity-approve.scss";
</style>
<script>
import loading from "@/plugin/loading";
import { mapState, mapGetters } from "vuex";

export default {
	data() {
		return {
			isBusy: false,
			currentPage: 1,
			perPage: 3,
			search: "",
			columns: [
				{ title: "ID", data: "student_id" },
				{ title: "Firstname", data: "firstname" },
				{ title: "Department", data: "department" }
			],
			// Table
			items: [],
			select: "",
			selectedItems: [],
			step: {
				next: Object.freeze({
					name: "Submit approve"
				}),
				back: Object.freeze({
					name: "Back"
				})
			},
			skKey: "",
			skKeyState: null
		};
	},
	computed: {
		...mapState("adminActivity", [
			"activity"
		]),
		...mapGetters("adminActivity", [
			"getActivityById"
		]),
		activityDetail() {
			return this.getActivityById(this.$route.params.id);
		},
		hasSelectedItem() {
			if (this.selectedItems.length > 0) {
				return true;
			}

			return false;
		}
	},
	async created() {
		this.isBusy = true;

		if (this.activityDetail) {
			this.isBusy = false;
			this.items = this.activityDetail.attendees;
			return;
		}

		try {
			await this.$store.dispatch("adminActivity/loadActivityById", this.$route.params.id);
		} catch (err) {
			this.$bvToast.toast(`Fetching data problem: ${err.message}`, {
				title: "Fetching activity error",
				variant: "danger",
				autoHideDelay: 1500
			});
		} finally {
			this.isBusy = false;
		}

		this.items = this.activity.attendees;
		this.rows = this.items.length;
	},
	methods: {
		singleValues(data) {
			this.selectedItems = data;
		},
		checkFormValidity() {
			const valid = this.$refs.form.checkValidity();
			this.skKeyState = !!valid;

			return valid;
		},
		showModal() {
			if (!this.hasSelectedItem) {
				this.$bvToast.toast("Please select at least one student", {
					title: "Select Student",
					variant: "danger",
					autoHideDelay: 1500
				});
				return;
			}

			if (this.secretKey) {
				this.submit(this.secretKey);
				return;
			}
			this.$refs.modal.show();
		},
		hideModal() {
			this.$refs.modal.hide();
		},
		resetModal() {
			this.skKey = "";
			this.skKeyState = null;
		},
		handleOk(bvModalEvt) {
			// Prevent modal from closing
			bvModalEvt.preventDefault();
			// Trigger submit handler
			this.handleSubmit();
		},
		handleSubmit() {
			// Exit when the form isn't valid
			if (!this.checkFormValidity()) {
				this.$bvToast.toast("Please insert the secret key before submitting", {
					title: "Secret key",
					variant: "danger",
					autoHideDelay: 1500
				});
				return;
			}
			// Submit
			this.submit(this.skKey);
			// Hide the modal manually
			this.$nextTick(() => {
				this.hideModal();
			});
		},
		async submit(skkey) {
			loading.start();
			try {
				await this.$store.dispatch("adminActivity/submitApprove", {
					sk: skkey,
					approvedStudents: this.selectedItems,
					activityId: this.$route.params.id
				});

				this.$router.push({
					name: "activity-approve-success",
					params: {
						id: this.$route.params.id
					}
				});
			} catch (err) {
				this.$bvToast.toast(`There was a problem submit approve: ${err.message}`, {
					title: "Submit approve activity",
					variant: "danger",
					autoHideDelay: 1500
				});
			} finally {
				loading.stop();
			}
		},
		goBack() {
			this.$router.push({
				name: "activity-detail",
				params: {
					id: this.$route.params.id
				}
			});
		}
	}
};
</script>