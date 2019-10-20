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
					@click="testVerify"
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
		<div class="box">
			<h2 class="box-header">
				Verify Status
			</h2>
			<ul class="form">
				<li
					v-for="(s, index) in success"
					:key="`${s}${index}${rerender}`"
				>
					<h1>Student ID: {{ s.student_id }}</h1>
					<template v-if="verifyData[index]">
						<p>
							Recorded in the blockchain
						</p>
						<code>
							{{ verifyData[index] }}
						</code>
					</template>
					<template v-else>
						<p>
							Wait for verification
						</p>
						<code>
							{{ verifyData[index] }}
						</code>
					</template>
				</li>
			</ul>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/give-badge-success";
</style>
<script>
import SuccessLogo from "@/components/admin/SuccessLogo.vue";
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
		]),
		...mapState("adminVerify", [
			"verifyData"
		])
		// hexTransactionId() {
		// 	return base64ToHex(this.success.transaction_id);
		// },
		// hexMerkle() {
		// 	return base64ToHex(this.success.merkleroot);
		// }
	},
	methods: {
		testVerify() {
			this.success.reduce(async (previousPromise, student) => {
				const payload = {
					data: {
						...student
					}
				};
				this.rerender++;
				await previousPromise;
				return this.$store.dispatch("adminVerify/verifyTransaction", payload);
			}, Promise.resolve());
		}
	}
};
</script>