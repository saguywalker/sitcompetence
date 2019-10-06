<template>
	<div class="activity-detail wrapper">
		<div class="activity-header">
			<span class="eyebrow">eyebrowContent</span>
			<h1 class="title">
				activityDetail.activity_name
			</h1>
			<h5>Organized by activityDetail.organizer</h5>
			<h5>Category: activityDetail.category</h5>
		</div>

		<h2>Location</h2>
		<p>activityDetail.location</p>

		<h2>Activity details</h2>
		<p>activityDetail.description | noValue</p>

		<h2 class="attendees-title">
			Competence
		</h2>
		<!-- <ul class="attendees-list">
						<li
							v-for="(item, index) in activityDetail.competences"
							:key="`${item}${index}`"
							class="item"
						>
							item.competence_name
						</li>
					</ul> -->

		<h2 class="attendees-title">
			Attendees
			<span class="badge">attendees.length</span>
		</h2>
		<!-- <ul class="attendees-list">
						<li
							v-for="(item, index) in attendees"
							:key="`${item}${index}`"
							class="item"
						>
							item
						</li>
					</ul> -->
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/activity-detail.scss";
</style>