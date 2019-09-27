<template>
	<div class="student-layout">
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
		</transition>
		<transition
			name="page"
			mode="out-in"
		>
			<router-view />
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
		}
	}
};
</script>