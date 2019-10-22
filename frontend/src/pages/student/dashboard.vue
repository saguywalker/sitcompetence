<template>
	<div class="dashboard">
		<div class="page-header">
			<h1 class="wrapper title">
				Welcome {{ loginUser }} {{ envi }}
			</h1>
		</div>
		<section class="wrapper">
			<b-tabs
				v-model="selectedTab"
				content-class="mt-3"
			>
				<b-tab title="Dashboard" />
				<b-tab title="Profile" />
			</b-tabs>
			<keep-alive>
				<component :is="currentTabComponent" />
			</keep-alive>
		</section>
	</div>
</template>
<script>
import DashboardContent from "@/components/student/DashboardContent.vue";
import DashboardProfile from "@/components/student/DashboardProfile.vue";
import { getLoginUser } from "@/helpers";

export default {
	components: {
		DashboardContent,
		DashboardProfile
	},
	data() {
		return {
			selectedTab: 0,
			tabs: ["content", "profile"]
		};
	},
	computed: {
		currentTabComponent() {
			return `dashboard-${this.tabs[this.selectedTab]}`;
		},
		loginUser() {
			return getLoginUser();
		}
	}
};
</script>