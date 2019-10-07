<template>
	<aside
		ref="sidebar"
		class="main-sidebar"
	>
		<section
			v-click-outside="{
				exclude: ['sidebar'],
				sidebar: ['toggleBtn'],
				handler: 'handleClickOutside'
			}"
			:class="[
				'sidebar',
				toggle ? 'desktop-hide' : ''
			]"
		>
			<div class="user-panel">
				<div class="user-panel-info">
					<p class="name">
						{{ loginUser }}
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
import IconGiveBadge from "@/components/icons/IconGiveBadge.vue";
import IconActivity from "@/components/icons/IconActivity.vue";
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";
import { widthSize } from "@/helpers/mixins";
import { clickOutside } from "@/helpers/directives/clickOutside";

export default {
	components: {
		IconGiveBadge,
		IconActivity,
		IconCheckCircle
	},
	directives: {
		clickOutside
	},
	filters: {
		longSurname(value) {
			if (value.length > 11) {
				return `${value.charAt(0)}.`;
			}
			return value;
		}
	},
	mixins: [widthSize],
	props: {
		toggle: {
			type: Boolean,
			required: true
		}
	},
	data() {
		return {
			windowWidth: 0
		};
	},
	computed: {
		loginUser() {
			return sessionStorage.getItem("user");
		}
	},
	methods: {
		handleClickOutside() {
			if (this.windowWidth < 992) {
				this.$emit("hide-mobile", false);
			}
		}
	}
};
</script>