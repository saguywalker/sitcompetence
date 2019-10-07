<template>
	<b-container>
		<code>
			{{ user }}
		</code>
	</b-container>
</template>
<script>
import { mapState } from "vuex";
import loading from "@/plugin/loading";

export default {
	computed: {
		...mapState("base", ["user"]),
		userDeta() {
			return sessionStorage.getItem("user");
		}
	},
	async created() {
		loading.start();

		try {
			await this.$store.dispatch("base/loadUserDetail");
		} catch (err) {
			this.$bvToast.toast(`There was a problem loading user data: ${err.message}`, {
				title: "User data",
				variant: "danger",
				autoHideDelay: 1500
			});
		} finally {
			loading.stop();
		}
	}
};
</script>