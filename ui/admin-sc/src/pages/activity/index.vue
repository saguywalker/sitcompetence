<template>
	<div class="activity-main">
		<section class="section">
			<div class="page-header">
				<h1 class="title">
					Activity
				</h1>
				<div class="tools">
					<router-link :to="{ name: 'create-activity' }">
						<b-button
							class="item"
							size="sm"
							variant="primary"
						>
							Create Activity
						</b-button>
					</router-link>
				</div>
			</div>
			<div class="box">
				<h1 class="box-header">
					Posted Activity
				</h1>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in postActivities"
						:key="`${activity.id}${index}`"
						:to="{
							name: 'activity_id',
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
			<div class="box">
				<h1 class="box-header">
					Saved Activity
				</h1>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in saveActivities"
						:key="`${activity.id}${index}`"
						:to="{
							name: 'activity_id',
							params: {
								id: activity.id
							}
						}"
						class="activity-card-link"
					>
						<activity-card
							:title="activity.title"
							:description="activity.description"
							date="19-08-2019"
						/>
					</router-link>
				</div>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/activity-main.scss";
</style>
<script>
import ActivityCard from "@/components/ActivityCard.vue";
import { mapState } from "vuex";
// TODO: GET list of activities
// Saved Act, Posted Act
export default {
	components: {
		ActivityCard
	},
	computed: {
		...mapState("activity", [
			"postActivities",
			"saveActivities"
		])
	}
};
</script>