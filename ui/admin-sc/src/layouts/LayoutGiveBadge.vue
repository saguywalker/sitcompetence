<template>
	<div class="give-badge">
		<section class="section">
			<div class="header">
				<h1 class="title">
					{{ currentPage }}
				</h1>
				<base-breadcrumb
					:items="breadcrumbList"
					class="breadcrumb"
				/>
			</div>
			<router-view />
		</section>
	</div>
</template>
<style lang="scss">
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
			breadcrumbList: [],
			step: {},
			isHideStep: false
		};
	},
	computed: {
		currentPage() {
			return this.breadcrumbList[this.breadcrumbList.length - 1];
		}
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