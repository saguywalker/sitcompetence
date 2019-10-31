<template>
	<div class="genkey wrapper">
		<section class="section">
			<div class="mb-4">
				<h1 class="mb-3">
					Welcome to SIT-Competence Admin
				</h1>
				<p>
					The SIT-Competence Admin is the application that can submit the badge to the blockchain
					, or we can say that it is one of the blockchain node, and it can make a transaction with this application.
				</p>
				<p class="mt-2">
					Therefore, we need you to generate a secure Public/Private key pair before entering to this application. This is for authencity and verifiable reasons.
				</p>
			</div>
			<div class="mb-4">
				<h2>We are using ED25519 keys</h2>
				<p>
					Following
					<a
						href="https://linux-audit.com/using-ed25519-openssh-keys-instead-of-dsa-rsa-ecdsa/"
						target="_blank"
					>
						best practices
					</a>
					, you should always favor ED25519 SSH keys, since they are more secure and have better performance over the other types.
				</p>
			</div>
			<div class="mb-4">
				<h2>Generating a new key pair</h2>
				<p>Creating a new SSH key pair:</p>
				<p>To generate a new ED25519 SSH key pair, click <strong>Random</strong> button to generate the keys until you are satisfied.</p>
				<b-button
					variant="primary"
					class="my-2"
					@click="randomKeyPair"
				>
					Random
				</b-button>
				<p><strong>Secret Key:</strong></p>
				<div class="code-block">
					<code class="code">{{ b64SecretKey }}</code>
				</div>
				<p><strong>Public Key:</strong></p>
				<div class="code-block">
					<code class="code">{{ b64PublicKey }}</code>
				</div>
			</div>
			<div class="mb-4">
				<h2>Confirmation</h2>
				<ol class="step-list">
					<li>
						Once the key pair is decided, you must keep the secret key with yourself.
						<span style="color: red">
							* Important ! Don't modify the key *
						</span>
						<p>To confirm, please copy the selected </p>
					</li>
					<li>
						(Optional) Keep the secret key in the browser local storage. It will be stored locally only on your device.
						<br><br>
						This is for the convenience. When you create a transaction in the blockchain (in our case, give badge),
						you will not have to submit the secret key every time before you make a transaction to the blockchain.
						You will see the example in next step.
					</li>
				</ol>
			</div>
			<div class="mb-4">
				<h2>Secret Key Use Case Example</h2>

				<p>Before submmitting the transaction. you have to insert your secret key.</p>
				<p>It will look like this:</p>
				<img
					class="img-example"
					src="https://i.imgur.com/QbnC1rJ.png"
					alt="submit-skkey-example"
				>
				<p>If you enable to keep the secret key in the local storage, you will not see this input prompt.</p>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/user/genkey.scss";
</style>
<script>
import { getLoginUser, getED25519KeyPair } from "@/helpers";
// import { Base } from "@/services";

export default {
	data() {
		return {
			keyPair: {},
			error: null
		};
	},
	computed: {
		user() {
			return getLoginUser();
		},
		b64PublicKey() {
			if (this.keyPair.publicKey) {
				return btoa(this.keyPair.publicKey);
			}

			return "";
		},
		b64SecretKey() {
			if (this.keyPair.secretKey) {
				return btoa(this.keyPair.secretKey);
			}

			return "";
		}
	},
	methods: {
		async submit() {
			// const response = await Base.postSetPublicKey(this.formattedPbKey);
			// if (response.status === 200) {
			// 	const notification = {
			// 		title: "Set Public key",
			// 		message: "Submit key successful",
			// 		variant: "success"
			// 	};
			// 	this.$store.dispatch("base/addNotification", notification);
			// 	this.$router.push({ name: "give-badge" });
			// } else {
			// 	this.$bvToast.toast("There was a problem submitting data", {
			// 		title: "Set Public key",
			// 		variant: "danger",
			// 		autoHideDelay: 1500
			// 	});
			// }
		},
		randomKeyPair() {
			this.keyPair = getED25519KeyPair();
		}
	}
};
</script>