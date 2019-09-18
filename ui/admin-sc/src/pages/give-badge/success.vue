<template>
	<div class="give-badge-success">
		<div class="box">
			<div class="form">
				<success-logo />
				<div class="form-group">
					<label class="label">Hash:</label>
					<a
						class="hash"
						href="#"
						@click="verifyHash"
					>
						{{ data.hash }}
					</a>
					<p class="description">
						Click the hash to verify
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
	data() {
		return {
			data: {
				hash: "Az2t2dSv2dSfk2CcFdf"
			}
		};
	},
	computed: {
		...mapState("giveBadge", ["selectedStudents", "steps"])
	},
	methods: {
		verifyHash() {
			this.$store.dispatch("verify/updateHashId", this.hash);
			this.$router.push({ name: "verify" });
		}
	}
};
</script>