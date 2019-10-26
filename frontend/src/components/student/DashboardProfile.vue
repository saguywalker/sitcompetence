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
@import "@/styles/pages/student/dashboard-profile.scss";
</style>
<script>
import { getLoginUser } from "@/helpers";

export default {
	data() {
		return {
			userData: {}
		};
	},
	computed: {
		user() {
			return getLoginUser();
		}
	},
	async created() {
		this.userData = this.user;
	},
	methods: {
		logout() {
			this.$store.dispatch("base/logout");
		}
	}
};
</script>