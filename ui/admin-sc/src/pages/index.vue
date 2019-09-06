<template>
	<b-container>
		<h1 class="title">
			{{ wtf }}
		</h1>
		<event-card
			v-for="(event, index) in events"
			:key="`${event}${index}`"
			:event="event"
			:data-index="index"
		/>
		<b-button
			variant="primary"
			@click="show"
		>
			Show toast
		</b-button>
		<b-button @click="modalShow = !modalShow">
			Open Modal
		</b-button>
		<b-modal v-model="modalShow">
			Hello From Modal!
		</b-modal>
	</b-container>
</template>
<script>
import EventCard from "@/components/EventCard.vue";
import { mapState } from "vuex";

export default {
	components: {
		EventCard
	},
	data() {
		return {
			message: "msddg",
			modalShow: false
		};
	},
	computed: {
		...mapState("events", [
			"events"
		]),
		wtf() {
			return `${this.message} Hi`;
		}
	},
	async created() {
		await this.$store.dispatch("events/fetchEvents");
	},
	methods: {
		show() {
			this.$bvToast.toast("Toast body content", {
				title: "Variant",
				variant: "success",
				solid: true
			});
		}
	}
};
</script>