<template>
	<div class="give-badge-success">
		<div class="box">
			<div class="form">
				<success-logo />
				<div class="form-group">
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
				</div>
				<router-link :to="{ name: 'give-badge' }">
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
@import "@/styles/pages/give-badge-success";
</style>
<script>
import SuccessLogo from "@/components/SuccessLogo.vue";
import loading from "@/plugin/loading";
import { base64ToHex } from "@/helpers";
import { mapState } from "vuex";

export default {
	components: {
		SuccessLogo
	},
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			// TODO: GET the success data and show in the form
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
			"steps",
			"selectedStudents"
		]),
		hexTransactionId() {
			return base64ToHex(this.success.transaction_id);
		},
		hexMerkle() {
			return base64ToHex(this.success.merkleroot);
		}
	},
	methods: {
		async verifyHash() {
			loading.start();

			try {
				await this.$store.dispatch("verify/verifyTransaction", {
					data: this.success.data,
					transaction_id: this.success.transaction_id
				});

				this.$router.push({ name: "verify-result" });
			} catch (err) {
				const notification = {
					title: `Error verify give badge: ${err.status}`,
					message: `There was a problem verifying data: ${err.message}`,
					variant: "danger"
				};

				this.$store.dispatch("notification/add", notification);
			} finally {
				loading.stop();
			}
		}
	}
};
</script>