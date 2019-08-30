<template>
	<div>
		<div class="title">
			{{ message }}
		</div>
		<p v-if="statusCode === 404">
			<nuxt-link to="/">
				Return to fcking HOME
			</nuxt-link>
		</p>
	</div>
</template>
<script>
import Vue from "vue";

export default Vue.extend({
	name: "NuxtError",
	props: {
		error: {
			type: Object,
			default: null
		}
	},
	head() {
		return {
			title: this.message
		};
	},
	computed: {
		statusCode() { // <--- Get the status code
			return (this.error && this.error.statusCode) || 500;
		},
		message() { // <--- Print the error
			return this.error.message;
		}
	}
});
</script>
