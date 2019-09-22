<template>
	<div
		:class="[
			'login',
			isShowDetailMobile ? 'mobile-active' : ''
		]"
	>
		<div
			class="detail"
			lg="8"
		>
			<div class="logo-wrapper">
				<p class="intro">
					Welcome to
				</p>
				<div class="logo">
					<sitcom-logo />
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
				<div class="login-form">
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
						variant="primary"
						@click="submit"
					>
						Login
					</b-button>
				</div>
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
import SitcomLogo from "@/components/SitcomLogo.vue";

export default {
	components: {
		SitcomLogo
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
		validateSubmit() {
			this.error.userName = this.userName === "";
			this.error.passWord = this.passWord === "";
		},
		submit() {
			this.validateSubmit();

			if (!this.isError) {
				this.$router.push({ name: "main" });
			}
		},
		handleShowDetailMobile() {
			this.isShowDetailMobile = !this.isShowDetailMobile;
		}
	}
};
</script>