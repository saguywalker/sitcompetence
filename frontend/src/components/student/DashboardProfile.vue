<template>
	<div class="dashboard-profile">
		<b-container>
			<b-row class="justify-content-md-center">
				<b-col lg="7">
					<h1 class="title">
						Update Profile
					</h1>
					<div class="setting card">
						<base-input-image
							v-model="userData.profile_picture"
							:src="profilePic"
							:label="inputImageLabel"
						/>
						<b-form-group
							:label="inputDescLabel"
							class="mt-3"
							label-for="input-2"
						>
							<b-form-textarea
								id="input-2"
								v-model="userData.motto"
								placeholder="Describe yourself..."
								rows="5"
							/>
						</b-form-group>
					</div>
					<b-button
						:disabled="isError"
						variant="primary"
						size="sm"
						@click="submit"
					>
						Submit
					</b-button>
				</b-col>
			</b-row>
		</b-container>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/student/dashboard-profile.scss";
</style>
<script>
import BaseInputImage from "@/components/BaseInputImage.vue";
import { getLoginUser } from "@/helpers";
import { mapState } from "vuex";

export default {
	components: {
		BaseInputImage
	},
	data() {
		return {
			profilePic: "",
			userData: {
				profile_picture: null,
				motto: ""
			}
		};
	},
	computed: {
		...mapState("dashboard", ["profile"]),
		isNoData() {
			return !this.profile.profile_picture && this.profile.motto.length === 0;
		},
		isError() {
			return !this.userData.profile_picture && this.userData.motto.length === 0;
		},
		user() {
			return getLoginUser();
		},
		inputImageLabel() {
			return this.profilePic ? "Edit Profile Image" : "Upload Profile Image";
		},
		inputDescLabel() {
			return this.userData.motto.length === 0 ? "Enter your description" : "Edit your description";
		}
	},
	created() {
		this.profilePic = this.user.additional.profile_picture;
		this.userData.motto = this.user.additional.motto;
	},
	methods: {
		async submit() {
			this.$Progress.start();
			try {
				let payload = {};
				if (this.userData.profile_picture === null && this.profilePic.length !== 0) {
					payload = {
						profile_picture: this.profilePic,
						motto: this.userData.motto
					};
				} else {
					payload = {
						...this.userData
					};
				}

				await this.$store.dispatch("dashboard/updateProfile", payload);
				this.$bvToast.toast("Edit successfully", {
					title: "Edit Student Data",
					variant: "success",
					autoHideDelay: 1500
				});
			} catch (err) {
				this.$Progress.fail();
				this.$bvToast.toast(`Edit profile: ${err.message}`, {
					title: "Edit profile error",
					variant: "danger",
					autoHideDelay: 1500
				});
			} finally {
				this.userData.profile_picture = null;
				this.$Progress.finish();
			}
		},
		logout() {
			this.$store.dispatch("base/logout");
		}
	}
};
</script>