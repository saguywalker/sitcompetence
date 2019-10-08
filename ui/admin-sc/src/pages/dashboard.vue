<template>
	<b-container />
</template>
<script>
import loading from "@/plugin/loading";
import { mapState } from "vuex";

export default {
	computed: {
		...mapState("base", ["user"]),
		userDeta() {
			return sessionStorage.getItem("user");
		}
	},
	async beforeRouteEnter(to, from, next) {
		loading.start();
		next();
	},
	async created() {
		try {
			await this.$store.dispatch("base/loadUserDetail");
			location.href = "http://localhost:8080/admin/give-badge";
		} catch (err) {
			this.$bvToast.toast(`There was a problem loading user data: ${err.message}`, {
				title: "User data",
				variant: "danger",
				autoHideDelay: 1500
			});
		}
	}
};
</script>