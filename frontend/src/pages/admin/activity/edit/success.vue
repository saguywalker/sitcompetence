<template>
	<div class="create-activity-success">
		<div class="box">
			<div class="form">
				<success-logo text="Edit" />
				<router-link
					:to="{ name: 'create-activity' }"
					class="mt-3"
				>
					<b-button
						variant="admin-primary"
						size="sm"
					>
						Home
					</b-button>
				</router-link>
			</div>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/create-activity-success.scss";
</style>
<script>
import SuccessLogo from "@/components/admin/SuccessLogo.vue";
import { mapState } from "vuex";

export default {
	components: {
		SuccessLogo
	},
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			// TODO: GET the success data and show in the form
			if (!vm.steps.includes("summary")) {
				vm.$router.replace({ name: "edit-activity" });
			}
		});
	},
	beforeRouteLeave(to, from, next) {
		this.$store.dispatch("editActivity/clearStep");
		next();
	},
	computed: mapState("editActivity", ["steps"])
};
</script>