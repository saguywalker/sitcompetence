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
					<b-button
						class="item"
						size="sm"
						variant="primary"
					>
						Approve Attendance
					</b-button>
				</div>
			</div>
			<div class="box">
				<h1 class="box-header">
					Posted Activity
				</h1>
				<div class="my-row">
					<router-link
						v-for="(activity, index) in activities"
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
					<activity-card
						title="Hello"
						description="Lorem ipsum dolor sit amet consectetur adipisicing elit. Mollitia vero debitis accusantium fuga maxime quisquam ullam ratione eum quidem velit officiis, unde exercitationem nobis ipsam quam eaque sunt iste sequi."
					/>
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

// TODO: GET list of activities
// Saved Act, Posted Act
export default {
	components: {
		ActivityCard
	},
	data() {
		return {
			activities: [
				{
					id: "2093",
					title: "Run",
					description: "Bab P' Toon"
				},
				{
					id: "321",
					title: "Forrest Gump",
					description: ""
				},
				{
					id: "8632",
					title: "Pokemon",
					description: "The best Pokemon Trainer in the world is Sasuke wtf..."
				},
				{
					id: "4433",
					title: "Jeab",
					description: ""
				},
				{
					id: "4321",
					title: "Prep",
					description: "Cheapest Flight"
				},
				{
					id: "6332",
					title: "Potato Corner",
					description: "French Fries"
				}
			]
		};
	}
};
</script>