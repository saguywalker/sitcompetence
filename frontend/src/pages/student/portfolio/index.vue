<template>
	<div class="portfolio">
		<div class="page-header">
			<div class="wrapper">
				<h2 class="title">
					Portfolio
				</h2>
			</div>
		</div>
		<div
			v-if="false"
			class="portfolio-actions"
		>
			<div class="right">
				<b-button
					variant="primary"
					class="item"
				>
					Edit
				</b-button>
				<b-button
					variant="primary"
					class="item"
				>
					Share
				</b-button>
			</div>
		</div>
		<section class="wrapper">
			<div class="portfolio-section">
				<aside class="profile-bar">
					<div class="profile-header">
						<base-image :size="resizeImage" />
						<div class="name">
							<h5>{{ user.username }}</h5>
						</div>
					</div>
					<p class="profile-description">
						Lorem ipsum dolor, sit amet consectetur adipisicing fdfdelit. Voluptatum quidem hic sequi distinctio consectetur pariatur nostrum nesciunt, culpa quis quaerat saepe laborum officia! Impedit non suscipit consequuntur nostrum rerum temporibus?
					</p>
				</aside>
				<template v-if="statePortfolios">
					<div class="portfolio-content">
						<b-row>
							<b-col
								v-for="(com, index) in statePortfolios"
								:key="`${com.competence_id}${forceReRender}`"
								cols="6"
								class="competence-wrapper"
							>
								<base-image
									:src="getCompetenceImgById(com.competence_id)"
									:size="resizeImage"
									class="sitcom-badge"
								/>
								<p class="name">
									{{ getCompetenceNameById(com.competence_id) }}
								</p>
								<transition
									:key="`${com.competence_id}${index}`"
									name="fade"
									mode="out-in"
								>
									<p
										v-if="show"
										class="verify-status"
									>
										<span class="icon">
											<icon-check-circle v-if="verify[index]" />
											<icon-time-circle v-else />
										</span>
										{{ verifyText(index) }}
									</p>
								</transition>
							</b-col>
						</b-row>
						<div class="portfolio-footer">
							<b-button
								:block="resizeVerifyButton"
								variant="primary"
								size="sm"
								class="action-item"
								@click="testVerify"
							>
								Verify by Blockchain
							</b-button>
							<b-button
								v-if="show"
								variant="outline"
								size="sm"
								class="action-item"
								@click="clearVerify"
							>
								Clear verify
							</b-button>
						</div>
					</div>
				</template>
				<template v-else>
					<div class="portfolio-content">
						<h1>No data</h1>
					</div>
				</template>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/student/portfolio.scss";
</style>
<script>
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";
import IconTimeCircle from "@/components/icons/IconTimeCircle.vue";
import { widthSize } from "@/helpers/mixins";
import { getLoginUser } from "@/helpers";
import { COMPETENCE } from "@/constants";
import { mapState } from "vuex";

export default {
	components: {
		IconCheckCircle,
		IconTimeCircle
	},
	mixins: [widthSize],
	data() {
		return {
			fullName: "Tindanai Wongpipattanopas",
			forceReRender: 0,
			ports: []
		};
	},
	computed: {
		...mapState("portfolio", {
			statePortfolios: "portfolios",
			verify: "verify",
			show: "show"
		}),
		user() {
			return getLoginUser();
		},
		resizeVerifyButton() {
			if (this.windowWidth >= 768) {
				return false;
			}

			return true;
		},
		resizeImage() {
			if (this.windowWidth >= 768) {
				return "100";
			}

			return "70";
		}
	},
	async created() {
		this.$Progress.start();

		try {
			await this.$store.dispatch("portfolio/loadPortfolio", this.user.uid);
		} catch (err) {
			this.$Progress.fail();
			this.$bvToast.toast(`Fetching data problem: ${err.message}`, {
				title: "Fetching competences error",
				variant: "danger",
				autoHideDelay: 1500
			});
		} finally {
			this.$Progress.finish();
		}
	},
	methods: {
		getCompetenceNameById(id) {
			return COMPETENCE[id].name;
		},
		getCompetenceImgById(id) {
			return COMPETENCE[id].img;
		},
		verifyText(id) {
			return this.verify[id] ? "Verified" : "Unverified";
		},
		testVerify() {
			this.statePortfolios.reduce(async (previousPromise, competence) => {
				const payload = {
					competence_id: competence.competence_id,
					semester: competence.semester,
					student_id: competence.student_id
				};
				this.forceReRender++;
				await previousPromise;
				const result = this.$store.dispatch("portfolio/verifyTransaction", payload);
				return result;
			}, Promise.resolve());
		},
		clearVerify() {
			this.$store.dispatch("portfolio/clearVerify");
		}
	}
};
</script>