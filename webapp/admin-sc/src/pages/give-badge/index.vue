<template>
	<div class="give-badge-main">
		<b-row>
			<b-col lg="8">
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
							<b-button-group>
								<b-button
									variant="outline-primary"
									size="sm"
								>
									Month
								</b-button>
								<b-button
									variant="outline-primary"
									size="sm"
								>
									Year
								</b-button>
								<b-button
									variant="outline-primary"
									size="sm"
								>
									Semesterr
								</b-button>
							</b-button-group>
						</div>
						<b-table
							ref="selectableTable"
							:items="items"
							:fields="fields"
							selectable
							select-mode="multi"
							selected-variant="primary"
							responsive="sm"
							@row-selected="onRowSelected"
						>
							<template
								slot="[more_details]"
								slot-scope="row"
							>
								<b-button
									size="sm"
									class="mr-2"
									@click="row.toggleDetails"
								>
									{{ row.detailsShowing ? 'Hide' : 'Show' }}
								</b-button>
							</template>
							<template
								slot="row-details"
								slot-scope="row"
							>
								<b-card>
									<b-row class="mb-2">
										<b-col
											sm="3"
											class="text-sm-right"
										>
											<b>Id:</b>
										</b-col>
										<b-col>{{ row.item.student_id }}</b-col>
									</b-row>

									<b-row class="mb-2">
										<b-col
											sm="3"
											class="text-sm-right"
										>
											<b>Name:</b>
										</b-col>
										<b-col>{{ row.item.first_name }}</b-col>
									</b-row>

									<b-button
										size="sm"
										@click="row.toggleDetails"
									>
										Hide
									</b-button>
								</b-card>
							</template>
							<template
								slot="[selected]"
								slot-scope="{ rowSelected }"
							>
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
					</div>
				</div>
			</b-col>
			<b-col lg="4">
				<div class="box">
					<h2 class="box-header">
						Selected student
					</h2>
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
			</b-col>
		</b-row>
		<base-page-step
			ref="stepRef"
			:step="step"
			@next="submit"
		/>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/give-badge-main.scss";
</style>
<script>
import IconCheck from "@/components/icons/IconCheck.vue";
import IconCrossCircle from "@/components/icons/IconCrossCircle.vue";
import BasePageStep from "@/components/BasePageStep.vue";
import { mapState } from "vuex";

export default {
	components: {
		IconCheck,
		IconCrossCircle,
		BasePageStep
	},
	data() {
		return {
			search: "",
			fields: ["selected", "student_id", "first_name", "last_name", "more_details"],
			items: [
				{ student_id: 59130500210, first_name: "Dickerson", last_name: "Macdonald" },
				{ student_id: 59130500445, first_name: "Larsen", last_name: "Shaw" },
				{ student_id: 59130500222, first_name: "Geneva", last_name: "Wilson" },
				{ student_id: 59130522033, first_name: "Jami", last_name: "Carney" }
			],
			selectMode: "multi",
			selectedItems: [],
			error: {
				selectedItems: false
			}
		};
	},
	computed: {
		...mapState("giveBadge", {
			selectedStudents: "selectedStudents"
		}),
		hasSelectedItem() {
			if (this.selectedItems.length > 0) {
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
	created() {
		this.selectedItems = this.selectedStudents;
	},
	mounted() {
		this.selectedItems.forEach((item) => {
			const index = this.items.findIndex((i) => i.student_id === item.student_id);
			this.$refs.selectableTable.selectRow(index);
		});
	},
	methods: {
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
		submit() {
			if (this.hasSelectedItem) {
				this.$store.dispatch("giveBadge/updateSelectedStudents", this.selectedItems);
				this.$router.push({ name: "give-badge-selection" });
			} else {
				this.error.selectedItems = true;
				this.$bvToast.toast("Please select at least one student to give", {
					title: "No student error",
					variant: "danger",
					autoHideDelay: 1500
				});
			}
		}
	}
};
</script>