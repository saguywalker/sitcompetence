<template>
	<section class="section">
		<div class="page-header-admin">
			<h2 class="title">
				{{ currentPage }}
			</h2>
			<base-breadcrumb
				:items="breadcrumbList"
				class="breadcrumb"
			/>
		</div>
		<router-view />
	</section>
</template>
<script>
import BaseBreadcrumb from "@/components/admin/BaseBreadcrumb.vue";

export default {
	components: {
		BaseBreadcrumb
	},
	data() {
		return {
			breadcrumbList: [],
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