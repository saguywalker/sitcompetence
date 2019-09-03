<template>
	<ol class="base-breadcrumb">
		<li
			v-for="(item, index) in items"
			:key="`${item}${index}`"
		>
			{{ item }}
		</li>
	</ol>
</template>
<style lang="scss" scoped>
@import "@/styles/components/base-breadcrumb.scss";
</style>
<script>
export default {
	props: {
		items: {
			type: Array,
			required: true
		}
	}
};
</script>