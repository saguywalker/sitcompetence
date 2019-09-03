<template>
	<div
		:class="[
			'default',
			isSidebarOpen ? 'sidebar-open' : ''
		]"
	>
		<nav-bar @is-toggle="handleSidebar" />
		<side-bar :toggle="isSidebarOpen" />
		<router-view class="content-wrapper" />
	</div>
</template>
<style lang="scss" scoped>
@import "@/styles/layouts/default.scss";
</style>
<script>
import SideBar from "@/components/SideBar.vue";
import NavBar from "@/components/NavBar.vue";

export default {
	components: {
		NavBar,
		SideBar
	},
	data() {
		return {
			isSidebarOpen: false
		};
	},
	methods: {
		handleSidebar(e) {
			this.isSidebarOpen = e;
		}
	}
};
</script>