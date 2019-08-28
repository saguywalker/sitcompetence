<template>
	<header class="main-header">
		<nuxt-link
			:to="{ name: 'index' }"
			:class="[
				'logo',
				isSidebarOpen ? 'sidebar-open' : ''
			]"
		>
			<span class="detail">
				<base-logo
					:size="30"
					class="img"
				/>
				<p>
					<strong>Admin</strong> SC
				</p>
			</span>
		</nuxt-link>
		<nav
			:class="[
				'nav-bar',
				isSidebarOpen ? 'sidebar-open' : ''
			]"
		>
			<button
				@click="toggleSidebar"
				class="item"
			>
				<icon-hamburger size="15" />
			</button>
			<button class="item">
				<base-profile-image
					size="30"
				/>
			</button>
		</nav>
	</header>
</template>
<style lang="scss" scoped>
@import "@/styles/components/navbar.scss";
</style>
<script lang="ts">
import Vue from "vue";
import BaseLogo from "@/components/BaseLogo.vue";
import BaseProfileImage from "@/components/BaseProfileImage.vue";
import IconHamburger from "@/components/icons/IconHamburger.vue";

export default Vue.extend({
	components: {
		BaseLogo,
		BaseProfileImage,
		IconHamburger
	},
	data() {
		return {
			isSidebarOpen: false
		};
	},
	methods: {
		toggleSidebar() {
			this.isSidebarOpen = !this.isSidebarOpen;
			this.$emit("is-toggle", this.isSidebarOpen);
		}
	}
});
</script>