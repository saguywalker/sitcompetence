<template>
	<div class="give-badge-success">
		<div class="box">
			<div class="form">
				<success-logo />
				<!-- <div class="form-group">
					<label class="label">Transaction ID:</label>
					<a
						class="hash"
						href="#"
						@click="verifyHash"
					>
						{{ hexTransactionId }}
					</a>
					<p class="description">
						Click the hash to verify
					</p>
					<label class="label">Merkle root:</label>
					<p class="hash">
						{{ hexMerkle }}
					</p>
				</div> -->
				<b-button
					variant="primary"
					class="mt-3"
					size="sm"
				>
					Verify
				</b-button>
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
@import "@/styles/pages/give-badge-success";
</style>
<script>
import SuccessLogo from "@/components/SuccessLogo.vue";
// import loading from "@/plugin/loading";
// import { base64ToHex } from "@/helpers";
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
	computed: {
		...mapState("giveBadge", [
			"success",
			"steps"
		])
		// hexTransactionId() {
		// 	return base64ToHex(this.success.transaction_id);
		// },
		// hexMerkle() {
		// 	return base64ToHex(this.success.merkleroot);
		// }
	}
};
</script>