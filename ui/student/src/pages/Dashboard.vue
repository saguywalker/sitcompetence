<template>
	<div class="dashboard">
		<div class="page-header">
			<h1 class="wrapper title">
				Welcome
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
import DashboardContent from "@/components/DashboardContent.vue";
import DashboardProfile from "@/components/DashboardProfile.vue";

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
		}
	}
};
</script>