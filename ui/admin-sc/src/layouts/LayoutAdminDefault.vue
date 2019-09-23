<template>
	<div
		:class="[
			'default',
			isSidebarOpen ? 'sidebar-open' : ''
		]"
	>
		<nav-bar
			ref="navbar"
			@is-toggle="handleSidebar"
		/>
		<side-bar
			:toggle="isSidebarOpen"
			@hide-mobile="hideMobile"
		/>
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
import { widthSize } from "@/helpers/mixins";

export default {
	components: {
		NavBar,
		SideBar,
		Notification
	},
	mixins: [widthSize],
	data() {
		return {
			isSidebarOpen: false
		};
	},
	methods: {
		handleSidebar(e) {
			if (this.windowWidth < 992 && e === this.isSidebarOpen) {
				this.isSidebarOpen = !this.isSidebarOpen;
			} else {
				this.isSidebarOpen = e;
			}
		},
		hideMobile(e) {
			this.isSidebarOpen = e;
		}
	}
};
</script>