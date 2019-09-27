<template>
	<div class="create-activity-detail">
		<div class="box">
			<h1 class="box-header">
				Insert the activity detail
			</h1>
			<form @submit.prevent="submit">
				<fieldset class="container">
					<b-row
						v-for="(type, index) in fields"
						:key="`${type.label}${index}`"
						class="my-3"
					>
						<b-col lg="3">
							<label
								:for="`type-${type.type}`"
								class="label"
							>
								{{ type.label }}:
							</label>
						</b-col>
						<b-col lg="9">
							<b-form-input
								v-if="type.type !== 'textarea'"
								:id="`type-${type.type}`"
								v-model="input[type.model]"
								:state="error[type.model]"
								:type="type.type"
								size="sm"
								@input="error[type.model] = null"
							/>
							<b-form-textarea
								v-else
								:id="`type-${type.type}`"
								v-model="input[type.model]"
								class="t-area"
							/>
						</b-col>
					</b-row>
				</fieldset>
			</form>
		</div>
		<base-page-step
			:step="step"
			@next="submit"
			@back="goBack"
		/>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/create-activity-detail.scss";
</style>
<script>
import { CREATE_ACTIVITY_FORM } from "@/constants/form";
import { mapState } from "vuex";

export default {
	data() {
		return {
			fields: CREATE_ACTIVITY_FORM,
			input: {
				activity_name: "",
				description: "",
				activity_date: "",
				start_time: "",
				location: "",
				organizer: "",
				category: ""
			},
			error: {
				activity_name: "",
				activity_date: "",
				start_time: "",
				location: "",
				organizer: "",
				category: ""
			}
		};
	},
	computed: {
		...mapState("editActivity", [
			"detailInput",
			"steps"
		]),
		step() {
			return this.$route.meta.step;
		},
		hasError() {
			return Object.values(this.error).some((err) => err !== null);
		}
	},
	async created() {
		if (this.detailInput.activity_name) {
			try {
				await this.$store.dispatch("editActivity/loadActivityById", this.$route.params.id);
				this.input = this.detailInput;
			} catch (err) {
				this.$bvToast.toast(`Fetch activity id:${this.$route.params.id} error`, {
					title: `Error ${err.message}`,
					variant: "danger",
					autoHideDelay: 1500
				});
			}
		}
	},
	methods: {
		validateSubmit() {
			Object.keys(this.input).forEach((key) => {
				if (!this.input[key]) {
					this.error[key] = false;
				}
			});

			this.error.description = null;
		},
		goBack() {
			this.$router.push({
				name: "activity-detail",
				params: {
					id: this.$route.params.id
				}
			});
		},
		async submit() {
			this.validateSubmit();

			if (this.hasError) {
				this.$bvToast.toast("Please fill all the error field", {
					title: "No input error",
					variant: "danger",
					autoHideDelay: 1500
				});
				return;
			}

			await this.$store.dispatch("editActivity/setDetailInput", this.input);
			await this.$store.dispatch("editActivity/addStep", this.step.step);
			this.$router.push({ name: "edit-activity-competence" });
		}
	}
};
</script>