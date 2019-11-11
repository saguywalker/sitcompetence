<template>
	<div class="activity-detail">
		<section class="section">
			<div class="activity-detail-header">
				<router-link
					:to="{ name: 'admin-activity' }"
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
								id: activityId
							}
						}"
					>
						<b-button
							size="sm"
							variant="admin-primary"
						>
							Approve Student
						</b-button>
					</router-link>
					<router-link
						v-if="false"
						:to="{
							name: 'edit-activity',
							params: {
								id: activityId
							}
						}"
					>
						<b-button size="sm">
							Edit
						</b-button>
					</router-link>
					<b-button
						size="sm"
						variant="danger"
						@click="deleteActivity"
					>
						Delete
					</b-button>
				</div>
			</div>
			<div class="box">
				<div class="content">
					<div class="activity-header">
						<span class="eyebrow">{{ eyebrowContent }}</span>
						<h2 class="title">
							{{ activityDetail.activity_name }}
						</h2>
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

					<h2 class="attendees-title">
						Attendees
						<span class="badge">{{ activityAttendeeNumber }}</span>
					</h2>
					<ul
						v-if="activityDetail.attendees"
						class="attendees-list"
					>
						<li
							v-for="(item, index) in activityDetail.attendees"
							:key="`${item}${index}`"
							class="item"
						>
							{{ item.firstname }}
							{{ item.department }}
						</li>
					</ul>
					<p v-else>
						No attendees
					</p>
				</div>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/activity-detail.scss";
</style>
<script>
import IconArrow from "@/components/icons/IconArrow.vue";
import loading from "@/plugin/loading";
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
	data() {
		return {
			attendees: ["Yoda", "Luke", "Darth", "Ray", "Han", "C3PO", "R2D2"]
		};
	},
	computed: {
		...mapState("adminActivity", [
			"activities",
			"activity"
		]),
		...mapGetters("adminActivity", [
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
		activityAttendeeNumber() {
			if (this.activityDetail.attendees) {
				return this.activityDetail.attendees.length;
			}

			return 0;
		}
		// activityAttendees() {
		// 	return this.activityDetail.map((activity) => activity.attendees);
		// }
	},
	async created() {
		loading.start();

		if (this.activities.length !== 0) {
			loading.stop();
			return;
		}

		try {
			await this.$store.dispatch("adminActivity/loadActivityById", this.activityId);
		} catch (err) {
			this.$bvToast.toast(`Fetching data problem: ${err.message}`, {
				title: "Fetching activity error",
				variant: "danger",
				autoHideDelay: 1500
			});
		} finally {
			loading.stop();
		}
	},
	methods: {
		deleteActivity() {
			this.$bvModal.msgBoxConfirm("Are you sure you want to delete? This can't be undone.", {
				title: "Delete Confirmation",
				size: "sm",
				buttonSize: "sm",
				okVariant: "danger",
				okTitle: "Yes",
				cancelTitle: "No",
				footerClass: "p-2",
				hideHeaderClose: false,
				centered: true
			})
				.then((value) => {
					loading.start();

					if (value) {
						this.$store.dispatch("activity/deleteActivityById", this.activityId)
							.then(() => {
								this.$bvToast.toast("Success delete activity", {
									title: "Delete successful",
									variant: "success",
									autoHideDelay: 1500
								});
								this.$router.push({ name: "activity" });
							})
							.catch((err) => {
								this.$bvToast.toast("Error delete activity", {
									title: "Deleting activity error",
									message: `There was a problem deleting: ${err.message}`,
									variant: "danger",
									autoHideDelay: 1500
								});
							});
					}
				})
				.finally(() => {
					loading.stop();
				});
		}
	}
};
</script>