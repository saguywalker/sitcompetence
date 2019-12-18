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
					<li id="confirm-pk">
						Once the key pair is decided, you must <strong><u>keep the secret key with yourself.</u></strong>
						<span style="color: red">
							* Important ! Don't modify the key *
						</span>
						<br><br>
						<p>To confirm, please copy the <strong>Public Key</strong> into this section</p>
						<b-form-group
							description="For the security reasons the content of a key cannot be modified once added."
							label="Enter your copied public key"
							label-for="public-key"
							:invalid-feedback="errorMessage"
							class="mt-4"
						>
							<b-form-textarea
								id="public-key"
								v-model="pbKey"
								:state="errorStatus"
								placeholder="Enter your public key here ..."
								class="mt-2"
								rows="5"
							/>
						</b-form-group>
					</li>
					<li>
						(Optional) Keep the secret key in the browser local storage. It will be stored locally only on your device.
						<b-form-checkbox
							v-model="saveSkLocal"
							:disabled="!errorStatus"
							name="check-button"
							class="my-3"
							switch
							@input="setSkLocal"
						>
							Save <strong>Secret Key</strong> on your local device <b>({{ yesNo }})</b>
						</b-form-checkbox>
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
					:src="require('../../../assets/images/genkeyex.png')"
					alt="submit-skkey-example"
				>
				<p>If you enable to keep the secret key in the local storage, you will not see this input prompt.</p>
			</div>
			<b-button
				:disabled="!errorStatus"
				variant="primary"
				@click="submit"
			>
				Proceed to the website
			</b-button>
			<br>
			<a
				v-if="!errorStatus"
				href="#confirm-pk"
				class="description"
			>
				Please confirm the public key before proceed to the website
			</a>
			<br>
			<b-button
				class="mt-2"
				size="sm"
				variant="danger"
				@click="logout"
			>
				Logout
			</b-button>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/user/genkey.scss";
</style>
<script>
import { getLoginUser, getED25519KeyPair, clearLoginState } from "@/helpers";
import { Base } from "@/services";
import loading from "@/plugin/loading";

export default {
	data() {
		return {
			pbKey: "",
			keyPair: {},
			saveSkLocal: false
		};
	},
	computed: {
		user() {
			return getLoginUser();
		},
		b64PublicKey() {
			return this.keyPair.public;
		},
		b64SecretKey() {
			return this.keyPair.secret;
		},
		errorMessage() {
			if (this.pbKey.length === 0) {
				return "Please enter your public key";
			}

			return "The generated public key and input public key are not the same";
		},
		errorStatus() {
			if (this.pbKey.length === 0 || this.b64PublicKey !== this.pbKey) {
				return false;
			}
			return true;
		},
		yesNo() {
			if (this.saveSkLocal) {
				return "Yes";
			}
			return "No";
		}
	},
	watch: {
		b64SecretKey() {
			this.saveSkLocal = false;
			localStorage.removeItem("sck");
		}
	},
	methods: {
		async submit() {
			loading.start();
			const response = await Base.postSetPublicKey(this.b64PublicKey);
			if (response.status === 200) {
				const notification = {
					title: "Set Public key",
					message: "Submit key successful",
					variant: "success"
				};
				this.$store.dispatch("base/addNotification", notification);
				this.$router.push({ name: "give-badge" });
			} else {
				this.$bvToast.toast("There was a problem submitting data", {
					title: "Set Public key",
					variant: "danger",
					autoHideDelay: 1500
				});
			}
			loading.stop();
		},
		randomKeyPair() {
			this.keyPair = getED25519KeyPair();
		},
		setSkLocal() {
			if (this.saveSkLocal) {
				localStorage.setItem("sck", this.b64SecretKey);
				return;
			}

			localStorage.removeItem("sck");
		},
		logout() {
			clearLoginState();
			this.$router.push({ name: "login" });
		}
	}
};
</script>