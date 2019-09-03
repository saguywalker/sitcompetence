<template>
	<div class="give-badge">
		<div class="header">
			<h1 class="title">
				Give badge
			</h1>
			<base-breadcrumb
				:items="breadcrumbList"
				class="breadcrumb"
			/>
		</div>
		<router-view />
	</div>
</template>
<style lang="scss" scoped>
@import "@/styles/layouts/give-badge.scss";
</style>
<script>
import BaseBreadcrumb from "@/components/BaseBreadcrumb.vue";

export default {
	components: {
		BaseBreadcrumb
	},
	data() {
		return {
			breadcrumbList: []
		};
	},
	watch: {
		$route() {
			this.updateBreadcrumb();
		}
	},
	mounted() {
		this.updateBreadcrumb();
	},
	methods: {
		updateBreadcrumb() {
			this.breadcrumbList = this.$route.meta.breadcrumb;
		}
	}
};
</script>