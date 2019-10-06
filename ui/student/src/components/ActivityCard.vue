<template>
	<div class="activity-card">
		<div class="heading">
			<h2 class="title">
				{{ activity.activity_name }}
			</h2>
		</div>
		<div class="competences">
			<h4 class="title">
				Competences
			</h4>
			<ul class="competences-list">
				<li
					v-for="(competence, index) in activity.competences"
					:key="`${competence}${index}`"
					class="item"
				>
					{{ competence.competence_name }}
				</li>
			</ul>
		</div>
		<ul class="info">
			<li class="info-item">
				<icon-calendar class="icon" />
				<span>{{ formatDate }}</span>
			</li>
			<li class="info-item">
				<icon-time class="icon" />
				<span>{{ activity.time }}</span>
			</li>
			<li class="info-item">
				<icon-location class="icon" />
				<span>{{ activity.location }}</span>
			</li>
		</ul>
		<p class="content">
			{{ activity.description }}
		</p>
	</div>
</template>
<style lang="scss">
@import "@/styles/components/activity-card.scss";
</style>
<script>
import IconCalendar from "@/components/icons/IconCalendar.vue";
import IconLocation from "@/components/icons/IconLocation.vue";
import IconTime from "@/components/icons/IconTime.vue";
import { getEditDateFormat } from "@/helpers";

export default {
	components: {
		IconCalendar,
		IconLocation,
		IconTime
	},
	props: {
		activity: {
			type: Object,
			default: () => {}
		}
	},
	computed: {
		formatDate() {
			return getEditDateFormat(this.activity.activity_date);
		}
	}
};
</script>