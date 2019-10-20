<template>
	<div class="toast">
		<b-toast
			v-for="(notification, index) in notifications"
			id="my-toast"
			:key="`${notification}${index}`"
			:variant="notification.variant"
			solid
			visible
			@hide="onClose(index)"
		>
			<template v-slot:toast-title>
				<strong>{{ notification.title }}</strong>
			</template>
			<p>{{ notification.message }}</p>
		</b-toast>
	</div>
</template>
<script>
import { mapState } from "vuex";

export default {
	computed: mapState("notification", ["notifications"]),
	methods: {
		onClose(index) {
			this.$store.dispatch("notification/remove", index);
		}
	}
};
</script>