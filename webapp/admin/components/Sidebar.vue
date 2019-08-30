<template>
	<aside class="main-sidebar">
		<section class="sidebar">
			<div class="user-panel">
				<base-profile-image size="50" />
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
			</ul>
		</section>
	</aside>
</template>
<style lang="scss" scoped>
@import "@/styles/components/sidebar.scss";
</style>
<script>
import BaseProfileImage from "@/components/BaseProfileImage.vue";

export default {
	components: {
		BaseProfileImage
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