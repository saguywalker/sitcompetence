<template>
	<div
		:class="[
			'login',
			isShowDetailMobile ? 'mobile-active' : ''
		]"
	>
		<notification />

		<div
			class="detail"
			lg="8"
		>
			<div class="logo-wrapper">
				<p class="intro">
					Welcome to
				</p>
				<div class="logo">
					<base-sitcom-logo />
					<p class="text">
						SIT-Competence
					</p>
				</div>
			</div>
			<div class="content">
				<p>
					This project is a student competencies recording
					system for the School of Information Technology
					(SIT) which uses blockchain technology for
					permanent and decentalized data storage. This
					project also provides two web applications for staff
					and students.
				</p>
			</div>
		</div>
		<div class="login-section">
			<div class="login-wrapper">
				<h4 class="title">
					Login
				</h4>
				<form @submit.prevent="submit">
					<base-input-text
						v-model="userName"
						:error-message="error.userName ? 'Please enter a username' : ''"
						type="text"
						label="KMUTT ID"
						name="kmutt-id"
						@input="error.userName = false"
					/>
					<base-input-text
						v-model="passWord"
						:error-message="error.passWord ? 'Please enter a password' : ''"
						type="password"
						label="KMUTT Password"
						name="kmutt-password"
						@input="error.passWord = false"
					/>
					<b-button
						type="submit"
						variant="primary"
					>
						Login
					</b-button>
				</form>
			</div>
		</div>
		<button
			:class="[
				'mobile-activate',
				isShowDetailMobile ? 'active' : ''
			]"
			@click="handleShowDetailMobile"
		>
			{{ showDetailText }}
		</button>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/login.scss";
</style>
<script>
import Notification from "@/components/Notification.vue";
import loading from "@/plugin/loading";
import { getErrorStatusString } from "@/helpers";

export default {
	components: {
		Notification
	},
	data() {
		return {
			userName: "",
			passWord: "",
			error: {
				userName: false,
				passWord: false
			},
			isShowDetailMobile: false
		};
	},
	computed: {
		isError() {
			return this.error.userName || this.error.passWord;
		},
		showDetailText() {
			return this.isShowDetailMobile ? "Back to login" : "What is SIT-Competence?";
		}
	},
	methods: {
		getErrorStatusString,
		validateSubmit() {
			this.error.userName = this.userName === "";
			this.error.passWord = this.passWord === "";
		},
		async submit() {
			this.validateSubmit();

			if (!this.isError) {
				loading.start();

				try {
					await this.$store.dispatch("base/doLogin", {
						username: this.userName,
						password: this.passWord
					});
				} catch (err) {
					const status = getErrorStatusString(err.message);
					if (status === "401") {
						this.$bvToast.toast("Invalid username or password", {
							title: "Login error",
							variant: "danger",
							autoHideDelay: 1500
						});
					} else if (status === "403") {
						this.$bvToast.toast("You have been blocked for 3 minutes", {
							title: "Login Blocked",
							variant: "danger",
							autoHideDelay: 1500
						});
					} else if (status === "500") {
						this.$bvToast.toast("Something wrong", {
							title: "Internal Server Error",
							variant: "danger",
							autoHideDelay: 1500
						});
					}
				} finally {
					loading.stop();
				}
			}
		},
		handleShowDetailMobile() {
			this.isShowDetailMobile = !this.isShowDetailMobile;
		}
	}
};
</script>