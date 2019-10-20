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
						Back home
					</b-button>
				</router-link>
			</div>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/admin/create-activity-success.scss";
</style>
<script>
import SuccessLogo from "@/components/admin/SuccessLogo.vue";

export default {
	components: {
		SuccessLogo
	}
};
</script>