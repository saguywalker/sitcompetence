<template>
	<div class="give-badge-selection">
		<div class="box dropdown">
			<h2 class="box-header">
				Select badge to give
			</h2>
			<div class="box-content">
				<ul
					class="selected"
				>
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
							<p>{{ item.studentId }} {{ item.fullName }}</p>
							<icon-arrow-dropdown class="icon" />
						</a>
						<transition name="slide-down">
							<div
								v-if="item.show"
								class="badge-form"
							>
								<b-row>
									<b-col
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:for="`${item.studentId}${index}`"
											class="badge-checkbox"
										>
											<base-image size="90" />
											<input
												:id="`${item.studentId}${index}`"
												type="checkbox"
												name="hi"
											>
										</label>
										<p>Team-working</p>
									</b-col>
									<b-col
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:for="`${item.studentId}${index}`"
											class="badge-checkbox"
										>
											<base-image size="90" />
											<input
												:id="`${item.studentId}${index}`"
												type="checkbox"
												name="hi"
											>
										</label>
										<p>Leadership</p>
									</b-col>
									<b-col
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:for="`${item.studentId}${index}`"
											class="badge-checkbox"
										>
											<base-image size="90" />
											<input
												:id="`${item.studentId}${index}`"
												type="checkbox"
												name="hi"
											>
										</label>
										<p>Communication</p>
									</b-col>
									<b-col
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:for="`${item.studentId}${index}`"
											class="badge-checkbox"
										>
											<base-image size="90" />
											<input
												:id="`${item.studentId}${index}`"
												type="checkbox"
												name="hi"
											>
										</label>
										<p>Flexible</p>
									</b-col>
									<b-col
										lg="3"
										cols="6"
										class="badge-wrapper"
									>
										<label
											:for="`${item.studentId}${index}`"
											class="badge-checkbox"
										>
											<base-image size="90" />
											<input
												:id="`${item.studentId}${index}`"
												type="checkbox"
												name="hi"
											>
										</label>
										<p>Super Flexible</p>
									</b-col>
								</b-row>
							</div>
						</transition>
					</li>
				</ul>
			</div>
		</div>
		<base-page-step
			:step="step"
			@next="submit"
			@back="goBack"
		/>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/give-badge-selection.scss";
</style>
<script>
import BasePageStep from "@/components/BasePageStep.vue";
import BaseImage from "@/components/BaseImage.vue";
import IconArrowDropdown from "@/components/icons/IconArrowDropdown.vue";
import { mapState } from "vuex";

export default {
	components: {
		BasePageStep,
		BaseImage,
		IconArrowDropdown
	},
	data() {
		return {
			selectStudent: [],
			options: [
				{ text: "Orange", value: "orange" },
				{ text: "Apple", value: "apple" },
				{ text: "Pineapple", value: "pineapple" },
				{ text: "Grape", value: "grape" }
			]
		};
	},
	computed: {
		...mapState("giveBadge", {
			selectedStudents: "selectedStudents"
		}),
		step() {
			return this.$route.meta.step;
		}
	},
	created() {
		this.selectStudent = this.selectedStudents;
	},
	methods: {
		submit() {
			this.$router.push({ name: "give-badge-confirmation" });
		},
		goBack() {
			this.$router.push({ name: "give-badge" });
		},
		handleSelect(index) {
			this.selectStudent[index].show = true;
		}
	}
};
</script>