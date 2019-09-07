<template>
	<div
		:style="{
			width: `${size}px`
		}"
		class="base-image"
	>
		<img
			:style="{
				borderRadius: imageRounded
			}"
			:src="imageSrc"
			alt="Image"
		>
	</div>
</template>
<style lang="scss">
@import "@/styles/components/base-image.scss";
</style>
<script>
export default {
	props: {
		src: {
			type: String,
			default: ""
		},
		size: {
			type: String,
			default: "150"
		},
		round: {
			type: Boolean,
			default: true
		}
	},
	computed: {
		imageSrc() {
			return this.src || `https://via.placeholder.com/${this.size}`;
		},
		imageRounded() {
			return this.round ? "50%" : "0";
		}
	}
};
</script>