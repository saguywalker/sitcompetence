<template>
	<div class="success-logo">
		<div class="circle">
			<div class="content">
				<icon-check-circle
					size="30"
					fill="white"
				/>
				<h2 class="text">
					{{ text }}
					<br>
					Successfully
				</h2>
			</div>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/components/admin/success-logo.scss";
</style>
<script>
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";

export default {
	components: {
		IconCheckCircle
	},
	props: {
		text: {
			type: String,
			default: "Submit"
		}
	}
};
</script>