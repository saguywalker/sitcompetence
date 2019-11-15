<template>
	<article class="activity-card">
		<div class="card-body">
			<h4 class="title">
				{{ title }}
			</h4>
			<p
				v-if="description"
				class="description"
			>
				Description: {{ description }}
			</p>
			<p
				v-if="date"
				class="date"
			>
				Start on: {{ date }}
			</p>
		</div>
	</article>
</template>
<style lang="scss">
@import "@/styles/components/admin/activity-card.scss";
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
		},
		date: {
			type: String,
			default: ""
		}
	}
};
</script>