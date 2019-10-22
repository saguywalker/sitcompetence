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
				<base-sitcom-logo
					:size="40"
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
				ref="toggleBtn"
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
				<icon-user-cog />
				<p class="profile-name">
					{{ loginUser }}
				</p>
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
					<h4 class="name">
						{{ loginUser }}
					</h4>
					<p class="role">
						Staff
					</p>
				</div>
				<div class="profile-dropdown-footer">
					<router-link :to="{ name: 'user-setting' }">
						<b-button size="sm">
							Setting
						</b-button>
					</router-link>
					<b-button
						size="sm"
						@click="logout"
					>
						Sign out
					</b-button>
				</div>
			</div>
		</nav>
	</header>
</template>
<style lang="scss">
@import "@/styles/components/admin/navbar.scss";
</style>
<script>
import IconHamburger from "@/components/icons/IconHamburger.vue";
import IconUserCog from "@/components/icons/IconUserCog.vue";
import { clickOutside } from "@/helpers/directives/clickOutside";

export default {
	components: {
		IconHamburger,
		IconUserCog
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
	computed: {
		loginUser() {
			return sessionStorage.getItem("user");
		}
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
		},
		logout() {
			this.$store.dispatch("base/logout");
		}
	}
};
</script>