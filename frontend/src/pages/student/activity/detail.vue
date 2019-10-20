<template>
	<div class="activity-detail wrapper">
		<div class="activity-action-header">
			<router-link
				:to="{ name: 'activity' }"
				class="back"
			>
				<icon-arrow />
				Back
			</router-link>
		</div>
		<div class="activity-header">
			<span class="eyebrow">{{ eyebrowContent }}</span>
			<h1 class="title">
				{{ activityDetail.activity_name }}
			</h1>
			<h5>Organized by {{ activityDetail.organizer }}</h5>
			<h5>Category: {{ activityDetail.category }}</h5>
		</div>

		<h2>Location</h2>
		<p>{{ activityDetail.location }}</p>

		<h2>Activity details</h2>
		<p>{{ activityDetail.description | noValue }}</p>

		<h2 class="attendees-title">
			Competence
		</h2>
		<ul class="attendees-list">
			<li
				v-for="(item, index) in activityDetail.competences"
				:key="`${item}${index}`"
				class="item"
			>
				{{ item.competence_name }}
			</li>
		</ul>
		<b-button
			variant="primary"
			@click="joinActivity"
		>
			Join activity
		</b-button>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/student/activity-detail.scss";
</style>
<script>
import IconArrow from "@/components/icons/IconArrow.vue";
import { getMonthNameByDateFormat, getYearByDateFormat } from "@/helpers";
import { mapState, mapGetters } from "vuex";

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
	computed: {
		...mapState("activity", [
			"activities",
			"activity"
		]),
		...mapGetters("activity", [
			"getActivityById"
		]),
		eyebrowContent() {
			return `@${this.activityDetail.time} on ${getMonthNameByDateFormat(this.activityDetail.activity_date)} ${getYearByDateFormat(this.activityDetail.activity_date)}`;
		},
		activityId() {
			return this.$route.params.id;
		},
		activityDetail() {
			if (this.activities.length !== 0) {
				return this.getActivityById(this.activityId);
			}

			return this.activity;
		},
		loginUser() {
			return sessionStorage.getItem("user");
		}
	},
	async created() {
		this.$Progress.start();

		if (this.activities.length !== 0) {
			this.$Progress.finish();
			return;
		}

		try {
			await this.$store.dispatch("activity/loadActivityById", this.activityId);
		} catch (err) {
			this.$bvToast.toast(`Fetching data problem: ${err.message}`, {
				title: "Fetching activity error",
				variant: "danger",
				autoHideDelay: 1500
			});
			this.$Progress.fail();
		} finally {
			this.$Progress.finish();
		}
	},
	methods: {
		async joinActivity() {
			this.$Progress.start();

			try {
				await this.$store.dispatch("activity/joinActivity", {
					activity_id: this.activityId,
					student_id: this.loginUser
				});
				this.$bvToast.toast("Join activity success", {
					title: "Join activity successful",
					variant: "success",
					autoHideDelay: 1500
				});
				this.$router.push({ name: "activity" });
			} catch (err) {
				this.$bvToast.toast(`Join activity problem: ${err.message}`, {
					title: "Joining activity error",
					variant: "danger",
					autoHideDelay: 1500
				});
				this.$Progress.fail();
			} finally {
				this.$Progress.finish();
			}
		}
	}
};
</script>