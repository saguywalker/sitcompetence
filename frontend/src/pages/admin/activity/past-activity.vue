<template>
	<div class="past-activity">
		<section class="section">
			<div class="box">
				<div class="box-header">
					<h1>Past Activity</h1>
					<b-form-select
						v-model="selected"
						:options="options"
						size="sm"
						class="semester-select"
					>
						<icon-calendar />
					</b-form-select>
				</div>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in pastActivities"
						:key="`${activity.id}${index}`"
						:to="{
							name: 'activity-detail',
							params: {
								id: activity.id
							}
						}"
						class="activity-card-link"
					>
						<activity-card
							:title="activity.title"
							:description="activity.description"
						/>
					</router-link>
				</div>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/past-activity.scss";
</style>
<script>
import ActivityCard from "@/components/admin/ActivityCard.vue";
import IconCalendar from "@/components/icons/IconCalendar.vue";
import { mapState, mapGetters } from "vuex";
// TODO: GET list of past activities
// Past Act
export default {
	components: {
		ActivityCard,
		IconCalendar
	},
	data() {
		return {
			selected: 12018,
			options: [
				{ value: 12018, text: "1/2561" },
				{ value: 22018, text: "2/2561" },
				{ value: 12019, text: "1/2562" },
				{ value: 22019, text: "1/2562" }
			],
			activites: []
		};
	},
	computed: {
		...mapState("activity", [
			"pastActivities"
		]),
		...mapGetters("activity", [
			"getApprovedActivitiesBySemester"
		])
	},
	created() {
		this.activites = this.pastActivities;
	},
	methods: {
		onSemesterChange(e) {
			this.activites = this.getApprovedActivitiesBySemester(e);
		}
	}
};
</script>