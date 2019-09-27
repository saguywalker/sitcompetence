<template>
	<div class="activity-main">
		<section class="section">
			<div class="page-header">
				<h1 class="title">
					Activity
				</h1>
				<div class="tools">
					<router-link :to="{ name: 'past-activity' }">
						<b-button
							class="item"
							size="sm"
						>
							Past Activity
						</b-button>
					</router-link>
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
					Now on student site
				</h1>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in postActivities"
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
			<div class="box">
				<h1 class="box-header">
					Saved activity
				</h1>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in saveActivities"
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
							date="19-08-2019"
						/>
					</router-link>
				</div>
			</div>
			<div class="box">
				<h1 class="box-header">
					From API activity
				</h1>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in acts"
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
import loading from "@/plugin/loading";
import { mapState } from "vuex";
// TODO: GET list of activities
// Saved Act, Posted Act
export default {
	components: {
		ActivityCard
	},
	data() {
		return {
			acts: []
		};
	},
	computed: {
		...mapState("activity", [
			"postActivities",
			"saveActivities",
			"activities"
		])
	},
	async created() {
		loading.start();

		try {
			await this.$store.dispatch("activity/loadActivity");
			this.acts = this.activities;
		} catch (err) {
			this.$bvToast.toast(`Fetching data problem: ${err.message}`, {
				title: "Fetching activity error",
				variant: "danger",
				autoHideDelay: 1500
			});
		} finally {
			loading.stop();
		}
	}
};
</script>