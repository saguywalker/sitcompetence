<template>
	<div>
		<b-toast
			v-for="(notification, index) in notifications"
			id="my-toast"
			:key="`${notification}${index}`"
			:variant="notification.variant"
			solid
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
	computed: mapState("base", ["notifications"])
};
</script>