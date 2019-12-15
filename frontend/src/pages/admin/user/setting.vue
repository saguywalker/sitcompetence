<template>
	<div class="user-setting">
		<section class="section">
			<div class="box">
				<h1 class="box-header">
					User setting
				</h1>
				<section class="section">
					<b-button
						v-b-toggle.saveSecret
						variant="admin-primary"
						size="sm"
					>
						Save secret key in your local device
					</b-button>
					<b-collapse
						id="saveSecret"
						class="mt-2"
					>
						<b-card>
							<h4>Save secret key in your device</h4>
							<b-form-group
								description="This input has an auto save feature."
								label-for="secret-key"
								class="mt-4"
							>
								<b-form-textarea
									id="secret-key"
									v-model="skKey"
									placeholder="Enter your secret key here ..."
									class="mt-2"
									rows="5"
								/>
							</b-form-group>
						</b-card>
					</b-collapse>
				</section>
			</div>
		</section>
	</div>
</template>
<script>
import { getSecretKey } from "@/helpers";

export default {
	data() {
		return {
			skKey: ""
		};
	},
	watch: {
		skKey() {
			localStorage.setItem("sck", this.skKey);
		}
	},
	created() {
		this.skKey = this.getSecretKey();
	},
	methods: {
		getSecretKey
	}
};
</script>