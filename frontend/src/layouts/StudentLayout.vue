<template>
	<div class="student-layout">
		<vue-progress-bar />
		<navbar
			:nav-close="mobileNavOpen"
			@nav-open="handleNavOpen"
		/>
		<transition
			name="fade"
			mode="out-in"
		>
			<div
				v-if="mobileNavOpen"
				class="nav-mobile"
			>
				<div class="main-navigation">
					<div class="link-group">
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
					<div class="navigation-secondary">
						<router-link :to="{ name: 'dashboard' }">
							<b-button variant="primary">
								Profile
							</b-button>
						</router-link>
						<b-button
							variant="primary"
							class="mt-3"
							@click="logout"
						>
							Sign out
						</b-button>
					</div>
				</div>
			</div>
		</transition>
		<transition
			name="page"
			mode="out-in"
		>
			<router-view class="main-layout" />
		</transition>
	</div>
</template>
<style lang="scss">
@import "@/styles/layouts/student-layout.scss";
</style>
<script>
import Navbar from "@/components/Navbar.vue";
import { widthSize } from "@/helpers/mixins";

export default {
	components: {
		Navbar
	},
	mixins: [widthSize],
	data() {
		return {
			isNavOpen: false
		};
	},
	computed: {
		mobileNavOpen() {
			if (this.windowWidth > 992) {
				return false;
			}

			return this.isNavOpen;
		}
	},
	watch: {
		$route() {
			this.isNavOpen = false;
		}
	},
	methods: {
		handleNavOpen(e) {
			this.isNavOpen = e;
		},
		logout() {
			this.$store.dispatch("authentication/logout");
		}
	}
};
</script>