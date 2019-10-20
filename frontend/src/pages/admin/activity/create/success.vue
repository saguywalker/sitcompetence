<template>
	<div class="create-activity-success">
		<div class="box">
			<div class="form">
				<success-logo text="Create" />
				<router-link
					:to="{ name: 'create-activity' }"
					class="mt-3"
				>
					<b-button
						variant="primary"
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
@import "@/styles/pages/create-activity-success.scss";
</style>
<script>
import SuccessLogo from "@/components/SuccessLogo.vue";
import { mapState } from "vuex";

export default {
	components: {
		SuccessLogo
	},
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			// TODO: GET the success data and show in the form
			if (!vm.steps.includes("summary")) {
				vm.$router.replace({ name: "create-activity" });
			}
		});
	},
	beforeRouteLeave(to, from, next) {
		this.$store.dispatch("createActivity/clearStep");
		next();
	},
	computed: mapState("createActivity", ["steps"])
};
</script>