<template>
	<div class="activity-detail">
		<section class="section">
			<div class="activity-detail-header">
				<router-link
					:to="{ name: 'activity' }"
					class="back-btn"
				>
					<icon-arrow />
					Back
				</router-link>
				<div class="function">
					<router-link
						:to="{
							name: 'activity-approve',
							params: {
								id: $route.params.id
							}
						}"
					>
						<b-button
							size="sm"
							variant="primary"
						>
							Approve Student
						</b-button>
					</router-link>
					<b-button
						size="sm"
					>
						Edit
					</b-button>
					<b-button
						size="sm"
						variant="danger"
					>
						Delete
					</b-button>
				</div>
			</div>
			<div class="box">
				<div class="content">
					<div class="activity-header">
						<span class="eyebrow">@10.00 on Aug 28 2019</span>
						<h1 class="title">
							{{ activityDetail.title }}
						</h1>
						<h5>Organized by Nongtiny</h5>
						<h5>Category: Play</h5>
					</div>

					<h2>Location</h2>
					<p>CB4 KMUTT</p>

					<h2>Activity details</h2>
					<p>{{ activityDetail.description | noValue }}</p>

					<h2 class="attendees-title">
						Attendees
						<span class="badge">{{ attendees.length }}</span>
					</h2>
					<ul class="attendees-list">
						<li
							v-for="(item, index) in attendees"
							:key="`${item}${index}`"
							class="item"
						>
							{{ item }}
						</li>
					</ul>
				</div>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/activity-detail.scss";
</style>
<script>
import IconArrow from "@/components/icons/IconArrow.vue";
import { mapGetters } from "vuex";

export default {
	components: {
		IconArrow
	},
	filters: {
		noValue(value) {
			if (value === "") {
				return "-";
			}

			return value;
		}
	},
	data() {
		return {
			attendees: ["Yoda", "Luke", "Darth", "Ray", "Han", "C3PO", "R2D2"]
		};
	},
	computed: {
		...mapGetters("activity", [
			"getActivityById"
		]),
		activityDetail() {
			return this.getActivityById(this.$route.params.id);
		}
	}
};
</script>