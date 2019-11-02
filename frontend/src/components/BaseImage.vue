<template>
	<div
		:style="{
			width: `${size}px`,
			height: `${size}px`,
			borderRadius: imageRounded
		}"
		class="base-image"
	>
		<img
			:style="{
				borderRadius: imageRounded
			}"
			:src="imageSrc"
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
		height: {
			type: String,
			default: ""
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
		},
		imageWidth() {
			if (this.size === "auto") {
				return "auto";
			}

			return `${this.size}px`;
		}
	}
};
</script>