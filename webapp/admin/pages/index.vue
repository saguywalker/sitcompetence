<template>
	<div>
		<h1 class="title">
			{{ message }}
		</h1>
		<h1 class="title">
			{{ wtf }}
		</h1>
	</div>
</template>
<script lang="ts">
import Vue from "vue";

export default Vue.extend({
	data() {
		return {
			message: "msg"
		};
	},
	computed: {
		wtf(): String {
			return `${this.message} Hi`;
		}
	}
});
</script>
