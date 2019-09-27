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
				<b-col lg="6">
					<activity-card />
				</b-col>
				<b-col lg="6">
					<activity-card />
				</b-col>
				<b-col lg="6">
					<activity-card />
				</b-col>
				<b-col lg="6">
					<activity-card />
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

export default {
	components: {
		ActivityCard
	}
};
</script>