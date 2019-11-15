<template>
	<div class="activity-approve">
		<section class="section">
			<b-row>
				<b-col lg="9">
					<div
						:class="[
							'box',
							error.selectedItems ? 'error' : ''
						]"
					>
						<h1 class="box-header">
							Approve student
						</h1>
						<div class="table">
							<div class="header">
								<b-button-group class="search">
									<b-form-input
										v-model="search"
										placeholder="Search here"
										size="sm"
									/>
									<b-button
										variant="admin-primary"
										size="sm"
									>
										Search
									</b-button>
								</b-button-group>
							</div>
							<b-table
								ref="selectableTable"
								:items="items"
								:fields="fields"
								:busy="isBusy"
								selectable
								select-mode="multi"
								selected-variant="admin-primary"
								responsive="sm"
								@row-selected="onRowSelected"
							>
								<template v-slot:cell(selected)="{ rowSelected }">
									<template v-if="rowSelected">
										<span aria-hidden="true">
											<icon-check />
										</span>
										<span class="sr-only">Selected</span>
									</template>
									<template v-else>
										<span aria-hidden="true" />
										<span class="sr-only">Not selected</span>
									</template>
								</template>
							</b-table>
							<div class="footer">
								<div class="text-center">
									<b-button
										class="mr-2"
										variant="admin-primary"
										size="sm"
										@click="selectAllRows"
									>
										Select all
									</b-button>
									<b-button
										variant="outline-admin-primary"
										size="sm"
										@click="clearSelected"
									>
										Clear selected
									</b-button>
								</div>
								<div class="mt-3">
									<b-pagination
										v-model="currentPage"
										:total-rows="rows"
										:per-page="perPage"
										aria-controls="my-table"
										align="center"
										size="sm"
									/>
								</div>
							</div>
						</div>
					</div>
				</b-col>
				<b-col lg="3">
					<div class="box scrollable">
						<h2 class="box-header">
							Selected student
						</h2>
						<div class="box-content scrollable">
							<ul class="selected">
								<li
									v-for="(item, index) in selectedItems"
									:key="`${item}${index}`"
									class="item"
								>
									{{ item.student_id }}
									<button
										class="delete"
										@click="deleteSelectedRow(item.student_id)"
									>
										<icon-cross-circle />
									</button>
								</li>
							</ul>
						</div>
					</div>
				</b-col>
			</b-row>
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
import IconCheck from "@/components/icons/IconCheck.vue";
import IconCrossCircle from "@/components/icons/IconCrossCircle.vue";
import loading from "@/plugin/loading";
import { mapState, mapGetters } from "vuex";

export default {
	components: {
		IconCheck,
		IconCrossCircle
	},
	data() {
		return {
			isBusy: false,
			currentPage: 1,
			perPage: 3,
			search: "",
			fields: ["selected", "student_id", "firstname", "department"],
			items: [],
			selectedItems: [],
			rows: null,
			error: {
				selectedItems: false
			},
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
	watch: {
		selectedItems() {
			this.error.selectedItems = false;
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
		checkFormValidity() {
			const valid = this.$refs.form.checkValidity();
			this.skKeyState = !!valid;

			return valid;
		},
		showModal() {
			if (!this.hasSelectedItem) {
				this.error.selectedItems = true;
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
		onRowSelected(items) {
			this.selectedItems = items;
		},
		selectAllRows() {
			this.$refs.selectableTable.selectAllRows();
		},
		clearSelected() {
			this.$refs.selectableTable.clearSelected();
		},
		deleteSelectedRow(id) {
			const index = this.items.findIndex((item) => item.student_id === id);
			this.$refs.selectableTable.unselectRow(index);
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