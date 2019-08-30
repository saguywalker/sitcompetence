<template>
	<div
		:class="[
			'default',
			isSidebarOpen ? 'sidebar-open' : ''
		]"
	>
		<nav-bar @is-toggle="handleSidebar" />
		<side-bar />
		<nuxt class="content-wrapper" />
	</div>
</template>
<style lang="scss" scoped>
@import "@/styles/layouts/default.scss";
</style>
<script>
import Vue from "vue";
import SideBar from "@/components/SideBar.vue";
import NavBar from "@/components/NavBar.vue";

export default Vue.extend({
	head() {
		return {
			title: "SITCOM - Admin"
		};
	},
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
});
</script>