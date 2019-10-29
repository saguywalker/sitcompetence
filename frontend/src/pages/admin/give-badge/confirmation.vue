<template>
	<div class="give-badge-selection">
		<div class="box dropdown">
			<h2 class="box-header">
				Confirmation Summary
			</h2>
			<div class="box-content">
				<ul class="selected-student">
					<li
						v-for="(item, index) in selectedStudents"
						:key="`${item}${index}`"
					>
						<a
							:class="[
								'dropdown',
								item.show ? 'active' : ''
							]"
							@click="item.show = !item.show"
						>
							<p>{{ item.student_id }} {{ item.fullName }}</p>
							<icon-arrow-dropdown class="icon" />
						</a>
						<transition name="slide-down">
							<div
								v-if="item.show"
								class="badge-form"
							>
								<div class="my-row">
									<b-col
										v-for="(badge, id) in item.badges"
										:key="`${badge}${item.student_id}${id}`"
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:class="[
												'badge-checkbox'
											]"
										>
											<base-image
												:src="getBadgeImgById(badge.competence_id)"
												size="90"
											/>
											<p class="text">{{ badge.competence_name }}</p>
										</label>
									</b-col>
								</div>
							</div>
						</transition>
					</li>
				</ul>
			</div>
		</div>
		<b-modal
			id="modal-prevent-closing"
			ref="modal"
			title="Submit Your Name"
			@show="resetModal"
			@hidden="resetModal"
			@ok="handleOk"
		>
			<form
				ref="form"
				@submit.stop.prevent="handleSubmit"
			>
				<b-form-group
					:state="skKeyState"
					label="Secret Key"
					label-for="sk-input"
					invalid-feedback="Secret key is required"
				>
					<b-form-input
						id="sk-input"
						v-model="skKey"
						:state="skKeyState"
						required
						@input="skKeyState = null"
					/>
				</b-form-group>
			</form>
		</b-modal>
		<base-page-step
			:step="step"
			@next="showModal"
			@back="goBack"
		/>
	</div>
</template>
<script>
import IconArrowDropdown from "@/components/icons/IconArrowDropdown.vue";
import { COMPETENCE } from "@/constants";
import loading from "@/plugin/loading";
import { getSemester, getLoginUser } from "@/helpers";
import { mapState } from "vuex";

export default {
	components: {
		IconArrowDropdown
	},
	data() {
		return {
			selectStudent: [],
			skKey: "",
			skKeyState: null
		};
	},
	computed: {
		...mapState("giveBadge", [
			"selectedStudents",
			"steps"
		]),
		...mapState("base", ["user"]),
		step() {
			return this.$route.meta.step;
		}
	},
	beforeRouteEnter(to, from, next) {
		next((vm) => {
			if (!vm.steps.includes("selection")) {
				vm.$router.replace({ name: "give-badge" });
			}
		});
	},
	created() {
		this.selectStudent = this.selectedStudents;
	},
	methods: {
		getBadgeImgById(id) {
			return COMPETENCE[id].img;
		},
		checkFormValidity() {
			const valid = this.$refs.form.checkValidity();
			this.skKeyState = !!valid;
			return valid;
		},
		showModal() {
			this.$refs.modal.show();
		},
		hideModal() {
			this.$refs.modal.hide();
		},
		resetModal() {
			this.skKey = "";
			this.skKeyState = null;
		},
		handleOk(bvModalEvt) {
			// Prevent modal from closing
			bvModalEvt.preventDefault();
			// Trigger submit handler
			this.handleSubmit();
		},
		handleSubmit() {
			// Exit when the form isn't valid
			if (!this.checkFormValidity()) {
				return;
			}
			// Submit
			this.submit();
			// Hide the modal manually
			this.$nextTick(() => {
				this.hideModal();
			});
		},
		async submit() {
			loading.start();

			try {
				await this.$store.dispatch("giveBadge/submitGiveBadge", {
					giver: getLoginUser(),
					semester: getSemester()
				});
				await this.$store.dispatch("giveBadge/addStep", this.step.step);
				this.$router.push({ name: "give-badge-success" });
			} catch (err) {
				this.$bvToast.toast(`There was a problem submitting data: ${err.message}`, {
					title: "Submit give badge",
					variant: "danger",
					autoHideDelay: 1500
				});
			} finally {
				loading.stop();
			}
		},
		async goBack() {
			await this.$store.dispatch("giveBadge/deleteStep", this.step.step);
			this.$router.push({ name: "give-badge-selection" });
		}
	}
};
</script>