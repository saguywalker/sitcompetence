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
					Therefore, we need you to generate a secure shell (SSH) Public/Private key pair before enteringto this application. This is for authencity and verifiable.
				</p>
				<p class="mt-2">
					Several tools exist to generate SSH public/private key pairs. The following sections show how to generate an SSH key pair on UNIX, UNIX-like and Windows platforms.
				</p>
			</div>
			<div class="mb-4">
				<h2>Requirements</h2>
				<p>
					The only requirement is to have the OpenSSH client installed on your system. This comes pre-installed on GNU/Linux and macOS, but not on Windows
				</p>
				<p class="mt-2">
					Depending on your Windows version, there are different methods to work with SSH keys.
				</p>
			</div>
			<div class="mb-4">
				<h3>Windows 10: Windows Subsystem for Linux</h3>
				<p>
					Starting with Windows 10, you can
					<a
						href="https://docs.microsoft.com/en-us/windows/wsl/install-win10"
						target="_blank"
					>
						install the Windows Subsystem for Linux (WSL)
					</a>
					where you can run Linux distributions directly on Windows, without the overhead of a virtual machine. Once installed and set up, you’ll have the Git and SSH clients at your disposal.
				</p>
			</div>
			<div class="mb-4">
				<h3>Windows 10, 8.1, and 7: Git for Windows</h3>
				<p>
					The easiest way to install the SSH client on Windows 8.1 and Windows 7 is
					<a
						href="https://gitforwindows.org/"
						target="_blank"
					>
						Git for Windows
					</a>
					. It provides a Bash emulation (Git Bash) used for running Git from the command line and the <code>ssh-keygen</code> command that is useful to create SSH keys as you’ll learn below.
				</p>
			</div>
			<div class="mb-4">
				<h3>ED25519 SSH keys</h3>
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
				<h2>Generating a new SSH key pair</h2>
				<p>To create a new SSH key pair:</p>
				<ol class="step-list">
					<li>Open a terminal on Linux or macOS, or Git Bash / WSL on Windows.</li>
					<li>
						Generate a new ED25519 SSH key pair:
						<div class="code-block">
							<code class="code">ssh-keygen -t ed25519</code>
						</div>
					</li>
					<li>
						Next, you will be prompted to input a file path to save your SSH key pair to. If you don’t already have an SSH key pair and aren’t generating a deploy key, use the suggested path by pressing <strong>Enter</strong>. Using the suggested path will normally allow your SSH client to automatically use the SSH key pair with no additional configuration.
					</li>
					<li>
						Once the path is decided, you will be prompted to input a password to secure your new SSH key pair. It’s a best practice to use a password, but it’s not required and you can skip creating it by pressing <strong>Enter</strong> twice.
						If, in any case, you want to add or change the password of your SSH key pair, you can use the -p flag:
						<div class="code-block">
							<code class="code">ssh-keygen -p -f &lt;keyname&gt;</code>
						</div>
					</li>
				</ol>
			</div>
			<div class="mb-4">
				<h2>Add the public key to SIT-Competence Admin</h2>
				<ol class="step-list">
					<li class="mb-3">
						Copy your public SSH key to the clipboard by using one of the commands below depending on your Operating System:
						<p><strong>macOS:</strong></p>
						<div class="code-block">
							<code class="code">pbcopy &lt; ~/.ssh/id_ed25519.pub</code>
						</div>
						<p><strong>WSL / GNU/Linux (requires the xclip package):</strong></p>
						<div class="code-block">
							<code class="code">xclip -sel clip &lt; ~/.ssh/id_ed25519.pub</code>
						</div>
						<p><strong>Git Bash on Windows:</strong></p>
						<div class="code-block">
							<code class="code">cat ~/.ssh/id_ed25519.pub | clip</code>
						</div>
						<p>You can also open the key in a graphical editor and copy it from there, but be careful not to accidentally change anything.</p>
					</li>
					<li>
						Add your public SSH key to your SIT-Competence account:
						<b-form-group
							description="For the security reasons the content of an SSH key cannot be modified once added."
							label="Enter your public key"
							label-for="public-key"
							invalid-feedback="Please enter your public key"
							class="mt-4"
						>
							<b-form-textarea
								id="public-key"
								v-model="pbKey"
								:state="error"
								placeholder="Enter your public key here ..."
								class="mt-2"
								@input="error = null"
							/>
						</b-form-group>
						<b-button
							variant="primary"
							size="sm"
							@click="submit"
						>
							Save
						</b-button>
					</li>
				</ol>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/user/genkey.scss";
</style>
<script>
import { getLoginUser } from "@/helpers";

export default {
	data() {
		return {
			pbKey: "",
			error: null
		};
	},
	computed: {
		user() {
			return getLoginUser();
		}
	},
	methods: {
		submit() {
			if (this.pbKey.length === 0) {
				this.error = false;
				return;
			}

			this.$router.push({ name: "give-badge" });
		}
	}
};
</script>