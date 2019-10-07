<template>
	<div class="give-badge-main">
		<b-row>
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
							<b-button-group class="search">
								<b-form-input
									v-model="search"
									placeholder="Search here"
									size="sm"
								/>
								<b-button
									variant="primary"
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
							:per-page="perPage"
							selectable
							select-mode="multi"
							selected-variant="primary"
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
									variant="primary"
									size="sm"
									@click="selectAllRows"
								>
									Select all
								</b-button>
								<b-button
									variant="outline-primary"
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
		<base-page-step
			:step="step"
			@next="submit"
		/>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/give-badge-main.scss";
</style>
<script>
import loading from "@/plugin/loading";
import IconCheck from "@/components/icons/IconCheck.vue";
import IconCrossCircle from "@/components/icons/IconCrossCircle.vue";
import { mapState } from "vuex";

export default {
	components: {
		IconCheck,
		IconCrossCircle
	},
	data() {
		return {
			currentPage: 1,
			perPage: 10,
			search: "",
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
			if (this.selectedItems.length > 0) {
				return true;
			}

			return false;
		},
		step() {
			return this.$route.meta.step;
		},
		rows() {
			return this.items.length;
		}
	},
	watch: {
		selectedItems() {
			this.error.selectedItems = false;
		}
	},
	async created() {
		if (this.students.length === 0) {
			loading.start();

			try {
				await this.$store.dispatch("base/loadStudentData");
			} catch (err) {
				this.$bvToast.toast(`There was a problem loading student data: ${err.message}`, {
					title: "Student Table Error",
					variant: "danger",
					autoHideDelay: 1500
				});
			} finally {
				loading.stop();
			}
		}

		this.items = this.students;
		this.selectedItems = this.selectedStudents;
		if (this.steps.includes("selection")) {
			this.setUpSelectedItems();
		}
	},
	mounted() {
		this.selectedItems.forEach((item) => {
			const index = this.items.findIndex((i) => i.student_id === item.student_id);
			this.$refs.selectableTable.selectRow(index);
		});
	},
	methods: {
		setUpSelectedItems() {
			this.items.forEach((item, index) => {
				this.selectedStudents.forEach((student) => {
					if (item.student_id === student.student_id) {
						this.items[index] = student;
					}
				});
			});
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

			await this.$store.dispatch("giveBadge/updateSelectedStudents", this.selectedItems);
			await this.$store.dispatch("giveBadge/addStep", this.step.step);
			this.$router.push({ name: "give-badge-selection" });
		}
		// async tableItemProvider() {
		// 	let items;
		// 	try {
		// 		items = await this.$store.dispatch("base/loadStudentDataByPage", this.currentPage);
		// 	} catch (err) {
		// 		this.$bvToast.toast(`There was a problem fetchin student table: ${err.message}`, {
		// 			title: "Student Table Error",
		// 			variant: "danger",
		// 			autoHideDelay: 1500
		// 		});
		// 	} finally {
		// 		loading.stop();
		// 	}

		// 	return items;
		// }
	}
};
</script>