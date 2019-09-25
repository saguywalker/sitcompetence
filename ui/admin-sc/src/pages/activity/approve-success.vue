<template>
	<div class="create-activity-success">
		<div class="box">
			<div class="form">
				<success-logo text="Approve" />
				<router-link
					:to="{ name: 'activity' }"
					class="mt-3"
				>
					<b-button
						variant="primary"
						size="sm"
					>
						Home
					</b-button>
				</router-link>
			</div>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/create-activity-success.scss";
</style>
<script>
import SuccessLogo from "@/components/SuccessLogo.vue";

export default {
	components: {
		SuccessLogo
	}
};
</script>