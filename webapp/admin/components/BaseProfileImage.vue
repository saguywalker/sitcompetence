<template>
	<div
		:style="{
			width: `${size}px`,
			height: `${size}px`
		}"
		class="base-profile-image"
	>
		<img
			:src="imageSrc"
			alt="Profile Image"
		>
	</div>
</template>
<style lang="scss" scoped>
@import "@/styles/components/base-profile-image.scss";
</style>
<script>
import Vue from "vue";

export default Vue.extend({
	props: {
		src: {
			type: String,
			default: ""
		},
		size: {
			type: String,
			default: "150"
		}
	},
	computed: {
		imageSrc() {
			return this.src || `https://via.placeholder.com/${this.size}`;
		}
	}
});
</script>