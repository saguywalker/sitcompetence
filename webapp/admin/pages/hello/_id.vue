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
	async fetch({ store, error, params }) {
		try {
			await store.dispatch("events/fetchEvent", params.id);
		} catch {
			error({
				statusCode: 503,
				message: `wow nuxtjs eiei ${params.id}`
			});
		}
	}
};
</script>