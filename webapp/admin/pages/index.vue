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
			@click="show"
			variant="primary"
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
<script lang="ts">
import Vue from "vue";
import EventCard from "@/components/EventCard.vue";
import { NuxtAppOptions } from "@nuxt/types/app";
import { mapState } from "vuex";

interface ServerData {
  data: object
}

export default Vue.extend({
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
		wtf(): string {
			return `${this.message} Hi`;
		}
	},
	async fetch({ store, error }: NuxtAppOptions) {
		try {
			await store!.dispatch("events/fetchEvents");
		} catch {
			error({
				statusCode: 503,
				message: "wow fucking shit"
			});
		}
	},
	methods: {
		show(): void {
			(this as any).$bvToast.toast("Toast body content", {
				title: "Variant",
				variant: "success",
				solid: true
			});
		}
	}
});
</script>