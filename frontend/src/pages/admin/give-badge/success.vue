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
		<div class="box">
			<h2 class="box-header">
				Blockchain Status Log
			</h2>
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
		]),
		...mapState("adminVerify", [
			"verifyData"
		])
	},
	methods: {
		testVerify() {
			this.success.reduce(async (previousPromise, student) => {
				const payload = {
					competence_id: student.competence_id,
					semester: student.semester,
					student_id: student.student_id
				};
				this.rerender++;
				await previousPromise;
				return this.$store.dispatch("adminVerify/verifyTransaction", payload);
			}, Promise.resolve());
		}
	}
};
</script>