<template>
	<aside class="main-sidebar">
		<section
			:class="[
				'sidebar',
				toggle ? 'desktop-hide' : ''
			]"
		>
			<div class="user-panel">
				<base-image
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
					<router-link
						:to="{ name: 'dashboard' }"
						class="item"
					>
						<icon-dashboard class="icon" />
						<span class="name">Dashboard</span>
					</router-link>
				</li>
				<li>
					<router-link
						:to="{ name: 'give-badge' }"
						class="item"
					>
						<icon-give-badge class="icon" />
						<span class="name">Give badge</span>
					</router-link>
				</li>
				<li>
					<router-link
						:to="{ name: 'activity' }"
						class="item"
					>
						<icon-activity class="icon" />
						<span class="name">Activity</span>
					</router-link>
				</li>
				<li>
					<router-link
						:to="{ name: 'badge-setting' }"
						class="item"
					>
						<icon-badge-setting class="icon" />
						<span class="name">Badge setting</span>
					</router-link>
				</li>
				<li>
					<router-link
						:to="{ name: 'verify' }"
						class="item"
					>
						<icon-check-circle
							class="icon"
							size="14"
						/>
						<span class="name">Verify</span>
					</router-link>
				</li>
			</ul>
		</section>
	</aside>
</template>
<style lang="scss">
@import "@/styles/components/sidebar.scss";
</style>
<script>
import IconDashboard from "@/components/icons/IconDashboard.vue";
import IconGiveBadge from "@/components/icons/IconGiveBadge.vue";
import IconActivity from "@/components/icons/IconActivity.vue";
import IconBadgeSetting from "@/components/icons/IconBadgeSetting.vue";
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";

export default {
	components: {
		IconDashboard,
		IconGiveBadge,
		IconActivity,
		IconBadgeSetting,
		IconCheckCircle
	},
	filters: {
		longSurname(value) {
			if (value.length > 11) {
				return `${value.charAt(0)}.`;
			}
			return value;
		}
	},
	props: {
		toggle: {
			type: Boolean,
			required: true
		}
	}
};
</script>