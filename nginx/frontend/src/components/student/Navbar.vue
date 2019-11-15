<template>
	<nav class="mynavbar wrapper">
		<router-link
			:to="{ name: 'dashboard' }"
			class="logo-wrapper"
		>
			<base-sitcom-logo
				:size="60"
				class="logo"
			/>
			<p class="text">
				SIT-Competence
			</p>
		</router-link>
		<div class="mynavbar-main">
			<div>
				<router-link
					:to="{ name: 'dashboard' }"
					class="item-link"
				>
					Dashboard
				</router-link>
				<router-link
					v-if="false"
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
			<div>
				<b-button
					variant="primary"
					class="mr-4"
					pill
					@click="logout"
				>
					Sign out
				</b-button>
				<router-link :to="{ name: 'dashboard' }">
					<icon-user fill="#3578ed" />
				</router-link>
			</div>
		</div>
		<hamburger
			:is-open="isNavMobileOpen"
			@click="handleNavMobileOpen"
		/>
	</nav>
</template>
<style lang="scss">
@import "@/styles/components/student/navbar.scss";
</style>
<script>
import Hamburger from "@/components/student/Hamburger.vue";
import IconUser from "@/components/icons/IconUser.vue";

export default {
	components: {
		Hamburger,
		IconUser
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
		},
		logout() {
			this.$store.dispatch("base/logout");
		}
	}
};
</script>