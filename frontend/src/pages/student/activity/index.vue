<template>
	<div class="activity">
		<div class="page-header">
			<div class="wrapper">
				<h2 class="title">
					Activity
				</h2>
			</div>
		</div>
		<div class="activity-body wrapper">
			<b-row>
				<b-col
					v-for="(activity, index) in activities"
					:key="`${activity.activity_id}${index}`"
					:class="[hasJoinActivityId(activity, loginUser.uid) ? 'disable' : '']"
					lg="6"
				>
					<router-link
						:to="{
							name: 'activity-detail_id',
							params: {
								id: activity.activity_id
							}
						}"
						class="link-detail"
					>
						<activity-card
							:activity="activity"
							:join="hasJoinActivityId(activity,loginUser.uid)"
						/>
					</router-link>
				</b-col>
			</b-row>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/student/activity.scss";
</style>
<script>
import ActivityCard from "@/components/student/ActivityCard.vue";
import { getLoginUser } from "@/helpers";
import { mapState } from "vuex";

export default {
	components: {
		ActivityCard
	},
	computed: {
		...mapState("activity", [
			"activities"
		]),
		loginUser() {
			return getLoginUser();
		}
	},
	methods: {
		hasJoinActivityId(activity, stdId) {
			if (activity.attendees) {
				return activity.attendees.some((student) => student.student_id === stdId);
			}

			return false;
		}
	}
};
</script>