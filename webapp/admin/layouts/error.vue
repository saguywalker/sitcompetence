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
<script lang="ts">
import Vue, { PropOptions } from "vue";

interface Error {
	statusCode: number,
	message: string
}

export default Vue.extend({
	name: "NuxtError",
	props: {
		error: {
			type: Object,
			default: null
		} as PropOptions<Error>
	},
	head(): { title: string } {
		return {
			title: this.message
		};
	},
	computed: {
		statusCode(): number { // <--- Get the status code
			return (this.error && this.error.statusCode) || 500;
		},
		message(): string { // <--- Print the error
			return this.error.message;
		}
	}
});
</script>
