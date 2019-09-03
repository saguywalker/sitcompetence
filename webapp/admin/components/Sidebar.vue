<template>
	<aside class="main-sidebar">
		<section class="sidebar">
			<div class="user-panel">
				<base-profile-image
					class="image"
					size="50"
				/>
				<div class="user-panel-info">
					<p class="name">
						Tindanai <span>{{ "Wongpipattanopas" | longSurname }}</span>
					</p>
					<p class="role">
						Staff
					</p>
				</div>
			</div>
			<ul class="sidebar-menu">
				<li class="header">
					Admin
				</li>
				<li>
					<nuxt-link
						:to="{ name: 'dashboard' }"
						class="item"
					>
						<icon-dashboard class="icon" />
						<span class="name">Dashboard</span>
					</nuxt-link>
				</li>
				<li>
					<nuxt-link
						:to="{ name: 'give-badge' }"
						class="item"
					>
						<icon-give-badge class="icon" />
						<span class="name">Give badge</span>
					</nuxt-link>
				</li>
				<li>
					<nuxt-link
						:to="{ name: 'activity' }"
						class="item"
					>
						<icon-activity class="icon" />
						<span class="name">Activity</span>
					</nuxt-link>
				</li>
				<li>
					<nuxt-link
						:to="{ name: 'badge-setting' }"
						class="item"
					>
						<icon-badge-setting class="icon" />
						<span class="name">Badge setting</span>
					</nuxt-link>
				</li>
			</ul>
		</section>
	</aside>
</template>
<style lang="scss" scoped>
@import "@/styles/components/sidebar.scss";
</style>
<script>
import BaseProfileImage from "@/components/BaseProfileImage.vue";
import IconDashboard from "@/components/icons/IconDashboard.vue";
import IconGiveBadge from "@/components/icons/IconGiveBadge.vue";
import IconActivity from "@/components/icons/IconActivity.vue";
import IconBadgeSetting from "@/components/icons/IconBadgeSetting.vue";

export default {
	components: {
		BaseProfileImage,
		IconDashboard,
		IconGiveBadge,
		IconActivity,
		IconBadgeSetting
	},
	filters: {
		longSurname(value) {
			if (value.length > 11) {
				return `${value.charAt(0)}.`;
			}
			return value;
		}
	}
};
</script>