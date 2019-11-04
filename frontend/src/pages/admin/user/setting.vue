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
						Save secret key
					</b-button>
					<b-collapse
						id="saveSecret"
						class="mt-2"
					>
						<b-card>
							<p class="card-text">
								Collapse contents Here
							</p>
							<b-form-group
								description="For the security reasons the content of a key cannot be modified once added."
								label-for="secret-key"
								:invalid-feedback="errorMessage"
								class="mt-4"
							>
								<b-form-textarea
									id="secret-key"
									v-model="skKey"
									:state="errorStatus"
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