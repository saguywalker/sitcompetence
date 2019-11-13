<template>
	<div class="give-badge-main">
		<b-row>
			<b-col lg="3">
				<div class="box scrollable">
					<h2 class="box-header">
						Selected student
					</h2>
					<div class="box-content scrollable">
						<ul class="selected">
							<li
								v-for="(item, index) in selectedStudents"
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
			<b-col lg="9">
				<div
					:class="[
						'box',
						error.selectedItems ? 'error' : ''
					]"
				>
					<h2 class="box-header">
						Student list
					</h2>
					<div class="table">
						<div class="header">
							<b-button-group class="search-header">
								<b-form-input
									v-model="searchValue"
									placeholder="Search here"
									size="sm"
								/>
								<b-button
									:disabled="!searchValue"
									size="sm"
									@click="searchValue = ''"
								>
									Clear
								</b-button>
							</b-button-group>
						</div>
						<b-table
							ref="selectableTable"
							:items="items"
							:fields="fields"
							:per-page="perPage"
							:current-page="currentPage"
							:busy="isBusy"
							:filter="searchValue"
							selectable
							select-mode="multi"
							selected-variant="admin-primary"
							responsive="sm"
							@row-selected="onRowSelected"
							@filtered="onSearched"
						>
							<template v-slot:table-busy>
								<div class="text-center text-danger my-2">
									<b-spinner
										variant="admin-primary"
										class="align-middle"
									/>
								</div>
							</template>
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
									class="admin-pagination"
									aria-controls="my-table"
									align="center"
									size="sm"
								/>
							</div>
						</div>
					</div>
				</div>
			</b-col>
		</b-row>
		<base-page-step
			:step="step"
			@next="submit"
		/>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/give-badge-main.scss";
</style>
<script>
import IconCrossCircle from "@/components/icons/IconCrossCircle.vue";
import IconCheck from "@/components/icons/IconCheck.vue";
import { getUnique } from "@/helpers";
import { mapState } from "vuex";

export default {
	components: {
		IconCrossCircle,
		IconCheck
	},
	data() {
		return {
			isBusy: false,
			searchValue: "",

			// Table
			currentPage: 1,
			perPage: 10,
			rows: null,
			fields: [
				{
					key: "selected"
				},
				{
					key: "student_id",
					sortable: true
				},
				{
					key: "firstname"
				},
				{
					key: "department"
				}
			],
			items: [],
			selectedItems: [],
			error: {
				selectedItems: false
			}
		};
	},
	computed: {
		...mapState("base", [
			"students"
		]),
		...mapState("giveBadge", [
			"selectedStudents",
			"steps"
		]),
		hasSelectedItem() {
			if (this.selectedStudents.length > 0) {
				return true;
			}

			return false;
		},
		step() {
			return this.$route.meta.step;
		}
	},
	watch: {
		selectedItems() {
			this.error.selectedItems = false;
		}
	},
	async created() {
		if (this.steps.length === 0) {
			this.isBusy = true;

			try {
				await this.$store.dispatch("base/loadStudentData");
				this.setupSelection();
			} catch (err) {
				this.$bvToast.toast(`There was a problem loading student data: ${err.message}`, {
					title: "Student Table Error",
					variant: "danger",
					autoHideDelay: 1500
				});
			} finally {
				this.isBusy = false;
			}
		}
		this.items = this.students;
		this.rows = this.items.length;
		this.selectedItems = this.selectedStudents;
	},
	mounted() {
		this.setupSelection();
	},
	updated() {
		this.setupUpdateSelection();
	},
	methods: {
		setupSelection() {
			this.selectedStudents.forEach((student) => {
				this.items.forEach((i, index) => {
					if (i.student_id === student.student_id) {
						this.$refs.selectableTable.selectRow(index);
					}
				});
			});
		},
		setupUpdateSelection() {
			this.selectedStudents.forEach((student) => {
				this.items.forEach((i, index) => {
					if (i.student_id === student.student_id) {
						const id = this.$refs.selectableTable.$refs.itemRows[index].$el.attributes[2].value;
						if (id) this.$refs.selectableTable.selectRow(id);
					}
				});
			});
		},
		async onRowSelected(items) {
			const arr = [
				...items,
				...this.selectedStudents
			];

			await this.$store.dispatch("giveBadge/updateSelectedStudents", getUnique(arr, "student_id"));
			this.selectedItems = this.selectedStudents;
		},
		selectAllRows() {
			this.$refs.selectableTable.selectAllRows();
		},
		clearSelected() {
			this.$refs.selectableTable.clearSelected();
		},
		handleSearch() {
			const searchOption = {
				key: this.selectedSearchBy,
				value: this.searchValue
			};

			this.$emit("searched", searchOption);
			this.currentPage = 1;
		},
		async deleteSelectedRow(stdId) {
			const index = this.items.findIndex((item) => item.student_id === stdId);
			const arr = this.selectedStudents.filter((value) => value.student_id !== stdId);
			await this.$store.dispatch("giveBadge/updateSelectedStudents", arr);
			this.selectedItems = this.selectedStudents;
			this.$refs.selectableTable.unselectRow(index);
		},
		async submit() {
			if (!this.hasSelectedItem) {
				this.error.selectedItems = true;
				this.$bvToast.toast("Please select at least one student to give", {
					title: "No student error",
					variant: "danger",
					autoHideDelay: 1500
				});
				return;
			}

			await this.$store.dispatch("giveBadge/addStep", this.step.step);
			this.$router.push({ name: "give-badge-selection" });
		},
		onSearched(e) {
			this.rows = e.length;
			this.currentPage = 1;
		}
	}
};
</script>