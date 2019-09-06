<template>
	<div
		v-if="hide"
		class="base-page-step"
	>
		<button
			v-if="step.back"
			class="link back"
			@click="goBack"
		>
			<span class="back">
				<icon-arrow />
				Back
			</span>
			<p class="name">
				{{ step.back.name }}
			</p>
		</button>
		<button
			v-if="step.next"
			class="link next"
			@click="goNext"
		>
			<span class="next">
				Next
				<icon-arrow />
			</span>
			<p class="name">
				{{ step.next.name }}
			</p>
		</button>
	</div>
</template>
<style lang="scss">
@import "@/styles/components/base-page-step.scss";
</style>
<script>
import IconArrow from "@/components/icons/IconArrow.vue";

export default {
	components: {
		IconArrow
	},
	props: {
		step: {
			type: Object,
			default: () => {}
		},
		hide: {
			type: Boolean,
			default: true
		}
	},
	methods: {
		goNext() {
			this.$emit("next");
		},
		goBack() {
			this.$emit("back");
		}
	}
};
</script>