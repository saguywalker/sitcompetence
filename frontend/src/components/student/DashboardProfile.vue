<template>
	<div class="dashboard-profile">
		<b-container>
			<b-row class="justify-content-md-center">
				<b-col lg="7">
					<h1 class="title">
						Update Profile
					</h1>
					<div class="setting card">
						<b-form-group
							label="Edit your username"
							label-for="input-1"
						>
							<b-form-input
								id="input-1"
								v-model="userData.username"
							/>
						</b-form-group>
					</div>
					<b-button
						variant="primary"
						@click="logout"
					>
						Sign out
					</b-button>
				</b-col>
			</b-row>
		</b-container>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/dashboard-profile.scss";
</style>
<script>
import { mapState } from "vuex";

export default {
	data() {
		return {
			userData: {}
		};
	},
	computed: {
		...mapState("base", [
			"user"
		])
	},
	async created() {
		this.$Progress.start();

		try {
			await this.$store.dispatch("base/loadUserDetail");
			this.userData = this.user;
		} catch (err) {
			this.$Progress.fail();
			this.$bvToast.toast(`Fetching user data problem: ${err.message}`, {
				title: "Fetching user error",
				variant: "danger",
				autoHideDelay: 1500
			});
		} finally {
			this.$Progress.finish();
		}
	},
	methods: {
		logout() {
			this.$store.dispatch("authentication/logout");
		}
	}
};
</script>