<template>
	<b-container />
</template>
<script>
import loading from "@/plugin/loading";
import { mapState } from "vuex";

export default {
	computed: {
		...mapState("base", ["user"]),
		userData() {
			return sessionStorage.getItem("user");
		}
	},
	beforeCreate() {
		loading.start();

		if (this.userData) {
			this.$router.push({ name: "give-badge" });
		}
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