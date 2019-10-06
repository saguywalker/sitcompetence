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
						<activity-card :activity="activity" />
					</router-link>
				</b-col>
			</b-row>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/activity.scss";
</style>
<script>
import ActivityCard from "@/components/ActivityCard.vue";
import { mapState } from "vuex";

export default {
	components: {
		ActivityCard
	},
	computed: {
		...mapState("activity", [
			"activities"
		])
	}
};
</script>