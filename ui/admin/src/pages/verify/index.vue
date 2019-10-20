<template>
	<div class="verify">
		<section class="section">
			<div class="box">
				<h1 class="box-header">
					Verify
				</h1>
				<div class="verify-content">
					<label
						for="hash-input"
						class="label"
					>
						Input the hash here:
					</label>
					<b-form-input
						id="hash-input"
						v-model="hash"
						class="input"
						placeholder="Search here"
						size="sm"
					/>
					<b-button
						variant="primary"
						size="sm"
						@click="verifyTransaction"
					>
						Search
					</b-button>
				</div>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/verify.scss";
</style>
<script>
import { mapState } from "vuex";

export default {
	data() {
		return {
			hash: ""
		};
	},
	computed: {
		...mapState("verify", [
			"hashId"
		])
	},
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			if (from.name === "give-badge-success") {
				vm.hash = vm.hashId;
			}
		});
	},
	methods: {
		verifyTransaction() {
			this.$store.dispatch("verify/verifyTransaction", this.hash);
		}
	}
};
</script>