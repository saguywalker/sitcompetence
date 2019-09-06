<template>
	<header class="main-header">
		<router-link
			:to="{ name: 'admin' }"
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
		</router-link>
		<nav
			:class="[
				'nav-bar',
				isSidebarOpen ? 'sidebar-open' : ''
			]"
		>
			<button
				class="item"
				@click="toggleSidebar"
			>
				<icon-hamburger size="15" />
			</button>
			<button
				ref="dropdown-button"
				class="item profile"
				@click="toggleProfile"
			>
				<base-profile-image size="30" />
				<h4 class="profile-name">
					Tindanai Wongpipattanopas
				</h4>
			</button>
			<div
				v-click-outside="{
					exclude: ['dropdown-button'],
					handler: 'handleClickOutside'
				}"
				:class="[
					'profile-dropdown',
					isProfileOpen ? 'is-open' : ''
				]"
			>
				<div class="profile-dropdown-detail">
					<base-profile-image
						class="img"
						size="83"
					/>
					<h4 class="name">
						Tindanai Wongpipattanopas
					</h4>
					<p class="role">
						Staff
					</p>
				</div>
				<div class="profile-dropdown-footer">
					<b-button size="sm">
						Setting
					</b-button>
					<b-button size="sm">
						Sign out
					</b-button>
				</div>
			</div>
		</nav>
	</header>
</template>
<style lang="scss">
@import "@/styles/components/navbar.scss";
</style>
<script>
import BaseLogo from "@/components/BaseLogo.vue";
import BaseProfileImage from "@/components/BaseProfileImage.vue";
import IconHamburger from "@/components/icons/IconHamburger.vue";
import { clickOutside } from "@/helpers/directives/clickOutside";

export default {
	components: {
		BaseLogo,
		BaseProfileImage,
		IconHamburger
	},
	directives: {
		clickOutside
	},
	data() {
		return {
			isSidebarOpen: false,
			isProfileOpen: false
		};
	},
	methods: {
		toggleSidebar() {
			this.isSidebarOpen = !this.isSidebarOpen;
			this.$emit("is-toggle", this.isSidebarOpen);
		},
		toggleProfile() {
			this.isProfileOpen = !this.isProfileOpen;
		},
		handleClickOutside() {
			this.isProfileOpen = false;
		}
	}
};
</script>