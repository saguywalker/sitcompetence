<template>
	<div class="give-badge-main">
		<b-row>
			<b-col lg="8">
				<div class="box">
					<div class="box-header">
						Student list
					</div>
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
									Semester
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
					<div class="box-header">
						Selected student
					</div>
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
		<router-link :to="{ name: 'give-badge-selection'}">
			<b-button>
				Next
			</b-button>
		</router-link>
	</div>
</template>
<style lang="scss" scoped>
@import "@/styles/pages/give-badge-main.scss";
</style>
<script>
import IconCheck from "@/components/icons/IconCheck.vue";
import IconCrossCircle from "@/components/icons/IconCrossCircle.vue";

export default {
	components: {
		IconCheck,
		IconCrossCircle
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
			selectedItems: []
		};
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
		}
	}
};
</script>