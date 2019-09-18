<template>
	<div
		:style="{
			width: imageWidth,
			height: height === '' ? 'auto' : `${height}px`
		}"
		class="base-image"
	>
		<img
			:style="{
				height: activityCard ? '140px' : '100%',
				borderRadius: imageRounded,
				borderTopLeftRadius: activityCard ? '5px' : 'none',
				borderTopRightRadius: activityCard ? '5px' : 'none'
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
		height: {
			type: String,
			default: ""
		},
		round: {
			type: Boolean,
			default: true
		},
		activityCard: {
			type: Boolean,
			default: false
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