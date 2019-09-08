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
		<notification />
	</div>
</template>
<style lang="scss">
@import "@/styles/layouts/default.scss";
</style>
<script>
import SideBar from "@/components/SideBar.vue";
import NavBar from "@/components/NavBar.vue";
import Notification from "@/components/Notification.vue";

export default {
	components: {
		NavBar,
		SideBar,
		Notification
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