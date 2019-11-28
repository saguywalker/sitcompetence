<template>
	<div class="activity-main">
		<section class="section">
			<div class="page-header-admin">
				<h1 class="title">
					Activity
				</h1>
				<div class="tools">
					<router-link
						v-if="false"
						:to="{ name: 'past-activity' }"
					>
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
							variant="admin-primary"
						>
							Create Activity
						</b-button>
					</router-link>
				</div>
			</div>
			<div
				v-if="savedActivities.length !== 0"
				class="box"
			>
				<h2 class="box-header">
					Saved activity
				</h2>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in savedActivities"
						:key="`${activity.activity_id}${index}`"
						:to="{
							name: 'activity-detail',
							params: {
								id: activity.activity_id
							}
						}"
						class="activity-card-link"
					>
						<activity-card
							:title="activity.activity_name"
							:description="activity.description"
						/>
					</router-link>
				</div>
			</div>
			<div
				v-if="postedActivities.length !== 0"
				class="box"
			>
				<h2 class="box-header">
					On Student site activity
				</h2>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in postedActivities"
						:key="`${activity.activity_id}${index}`"
						:to="{
							name: 'activity-detail',
							params: {
								id: activity.activity_id
							}
						}"
						class="activity-card-link"
					>
						<activity-card
							:title="activity.activity_name"
							:description="activity.description"
						/>
					</router-link>
				</div>
			</div>
			<div
				v-if="approveActivities.length !== 0"
				class="box"
			>
				<h2 class="box-header">
					Approved activity
				</h2>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/activity-main.scss";
</style>
<script>
import ActivityCard from "@/components/admin/ActivityCard.vue";
import loading from "@/plugin/loading";
import { mapGetters } from "vuex";

export default {
	components: {
		ActivityCard
	},
	computed: {
		...mapGetters("adminActivity", [
			"postedActivities",
			"savedActivities",
			"approveActivities"
		])
	},
	async created() {
		loading.start();

		try {
			await this.$store.dispatch("adminActivity/loadActivity");
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