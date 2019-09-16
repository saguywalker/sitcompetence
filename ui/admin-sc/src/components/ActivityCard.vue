<template>
	<article class="activity-card">
		<base-image
			class="card-img"
			size="250"
			:round="false"
			:activity-card="true"
		/>
		<div class="card-body">
			<h4 class="title">
				{{ title }}
			</h4>
			<p
				v-if="description"
				class="description"
			>
				{{ description }}
			</p>
		</div>
	</article>
</template>
<style lang="scss">
@import "@/styles/components/activity-card.scss";
</style>
<script>
export default {
	props: {
		title: {
			type: String,
			required: true
		},
		description: {
			type: String,
			default: ""
		}
	}
};
</script>