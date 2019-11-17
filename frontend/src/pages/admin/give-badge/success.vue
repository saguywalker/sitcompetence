<template>
	<div class="give-badge-success">
		<div class="box">
			<div class="form">
				<success-logo />
				<router-link :to="{ name: 'give-badge' }">
					<b-button
						class="mt-2"
						size="sm"
					>
						Back home
					</b-button>
				</router-link>
			</div>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/give-badge-success";
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
			if (!vm.steps.includes("confirmation")) {
				vm.$router.replace({ name: "give-badge" });
			}
		});
	},
	beforeRouteLeave(to, from, next) {
		this.$store.dispatch("giveBadge/clearStep");
		next();
	},
	data() {
		return {
			rerender: 0,
			status: []
		};
	},
	computed: {
		...mapState("giveBadge", [
			"success",
			"steps"
		])
	}
};
</script>