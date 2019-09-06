<template>
	<div>
		<h1>Event #{{ event.title }}</h1>
		<h1>Event #{{ event.time }}</h1>
		<h1>Event #{{ event.date }}</h1>
	</div>
</template>
<style lang="scss" scoped>

</style>
<script>
import { mapState } from "vuex";

export default {
	computed: {
		...mapState("events", ["event"])
	},
	async created() {
		await this.$store.dispatch("events/fetchEvent", this.$router.params.id);
	}
};
</script>