<template>
	<button
		:class="[
			'hamburger',
			isOpen ? 'active' : ''
		]"
		@click="toggleHam"
	>
		<span class="line" />
		<span class="line" />
		<span class="line" />
	</button>
</template>
<style lang="scss">
@import "@/styles/components/hamburger.scss";
</style>
<script>
export default {
	props: {
		isOpen: {
			type: Boolean,
			default: false
		}
	},
	methods: {
		toggleHam() {
			this.$emit("click");
		}
	}
};
</script>