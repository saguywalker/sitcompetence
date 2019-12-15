<template>
	<div class="give-badge-main">
		<v-select-page
			v-model="select"
			:data="items"
			:tb-columns="columns"
			:multiple="true"
			:width="tableSize"
			title="Select student to give"
			placeholder="Click here to select"
			key-field="student_id"
			show-field="student_id"
			@values="singleValues"
		/>
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
import { widthSize } from "@/helpers/mixins";
import { mapState } from "vuex";

export default {
	mixins: [widthSize],
	data() {
		return {
			columns: [
				{ title: "ID", data: "student_id" },
				{ title: "Firstname", data: "firstname" },
				{ title: "Department", data: "department" }
			],
			// Table
			items: [],
			select: "",
			selectedItems: []
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
			if (this.select.length > 0) {
				return true;
			}

			return false;
		},
		step() {
			return this.$route.meta.step;
		},
		tableSize() {
			if (this.windowWidth > 992) {
				return 700;
			}

			return 300;
		}
	},
	async created() {
		if (this.steps.length === 0) {
			this.isBusy = true;

			try {
				await this.$store.dispatch("base/loadStudentData");
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
	},
	methods: {
		async submit() {
			if (!this.hasSelectedItem) {
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
		},
		singleValues(data) {
			this.selectedItems = data;
		}
	}
};
</script>