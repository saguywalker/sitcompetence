<template>
	<nav class="mynavbar">
		<router-link
			:to="{ name: 'dashboard' }"
			class="logo-wrapper"
		>
			<sitcom-logo
				:size="60"
				class="logo"
			/>
			<p class="text">
				SIT-Competence
			</p>
		</router-link>
		<div class="mynavbar-main">
			<router-link
				:to="{ name: 'dashboard' }"
				class="item-link"
			>
				Dashboard
			</router-link>
			<router-link
				:to="{ name: 'activity' }"
				class="item-link"
			>
				Activity
			</router-link>
			<router-link
				:to="{ name: 'portfolio' }"
				class="item-link"
			>
				Portfolio
			</router-link>
		</div>
		<hamburger
			:is-open="isNavMobileOpen"
			@click="handleNavMobileOpen"
		/>
	</nav>
</template>
<style lang="scss">
@import "@/styles/components/navbar.scss";
</style>
<script>
import SitcomLogo from "@/components/SitcomLogo.vue";
import Hamburger from "@/components/Hamburger.vue";
// In progress
export default {
	components: {
		SitcomLogo,
		Hamburger
	},
	props: {
		navClose: {
			type: Boolean,
			default: false
		}
	},
	data() {
		return {
			isNavMobileOpen: false
		};
	},
	watch: {
		$route() {
			this.isNavMobileOpen = false;
		}
	},
	methods: {
		handleNavMobileOpen() {
			if (this.navClose === false && this.isNavMobileOpen === true) {
				this.isNavMobileOpen = false;
			} else {
				this.isNavMobileOpen = !this.isNavMobileOpen;
			}

			this.$emit("nav-open", this.isNavMobileOpen);
		}
	}
};
</script>