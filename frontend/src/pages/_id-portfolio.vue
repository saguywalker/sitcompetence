<template>
	<div
		:style="{
			backgroundImage: `url('${require('@/assets/images/port-bg.svg')}')`,
			backgroundSize: '320px',
			backgroundAttachment: 'fixed',
			backgroundPosition: 'center',
			backgroundRepeat: 'repeat',
		}"
		class="portfolio"
	>
		<header class="page-header port-id">
			<div class="title wrapper">
				<h2>Portfolio</h2>
			</div>
		</header>
		<section class="wrapper">
			<div class="portfolio-actions">
				<div class="right">
					<b-form-select
						v-model="selectedSemester"
						:options="semesters"
						class="select-semester"
						@input="filterPortfolio"
					/>
				</div>
			</div>
			<div class="portfolio-section">
				<aside class="profile-bar">
					<div class="profile-header">
						<base-image :size="resizeImage" />
						<div class="name">
							<h5>{{ user }}</h5>
						</div>
					</div>
					<p class="profile-description">
						Lorem ipsum dolor, sit amet consectetur adipisicing elit. Voluptatum quidem hic sequi distinctio consectetur pariatur nostrum nesciunt, culpa quis quaerat saepe laborum officia! Impedit non suscipit consequuntur nostrum rerum temporibus?
					</p>
				</aside>
				<div class="portfolio-content">
					<transition
						name="page"
						mode="out-in"
					>
						<b-row
							v-if="hasCompetence"
							class="portfolio-container"
						>
							<b-col
								v-for="(com, index) in portfolios"
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
					</transition>
					<transition
						name="fade"
						mode="out-in"
					>
						<h1 v-if="!hasCompetence">
							No competence record on this semester.
						</h1>
					</transition>
					<div
						v-if="hasCompetence"
						class="portfolio-footer"
					>
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
			</div>
		</section>
		<footer class="page-footer">
			<div class="wrapper">
				This is a footer.<br>Made with heart by SIT-competence team 💖.
				{{ eviden }}
			</div>
		</footer>
	</div>
</template>
<style lang="scss">
@import "@/styles/portfolio.scss";
</style>
<script>
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";
import IconTimeCircle from "@/components/icons/IconTimeCircle.vue";
import { widthSize } from "@/helpers/mixins";
import { filterBySemester } from "@/helpers";
import { COMPETENCE } from "@/constants";
import { Verify } from "@/services";
import { mapState } from "vuex";

export default {
	components: {
		IconCheckCircle,
		IconTimeCircle
	},
	mixins: [widthSize],
	data() {
		return {
			portfolios: [],
			verify: [],
			selectedSemester: 12019,
			semesters: [
				{ value: 12019, text: "1/2019" },
				{ value: 22018, text: "2/2018" },
				{ value: 12018, text: "1/2018" },
				{ value: 22017, text: "2/2017" },
				{ value: 12017, text: "1/2017" }
			],
			show: false,
			forceReRender: 0,
			hasCompetence: null
		};
	},
	computed: {
		...mapState("base", [
			"sharePortfolio"
		]),
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
		},
		statePortfolios() {
			return this.sharePortfolio.collected_competence;
		},
		user() {
			return this.sharePortfolio.firstname;
		},
		eviden() {
			return atob(this.sharePortfolio.evidence);
		}
	},
	methods: {
		filterBySemester,
		filterPortfolio() {
			this.portfolios = this.filterBySemester(this.statePortfolios.result, this.selectedSemester);
			this.hasCompetence = this.portfolios.length !== 0;
		},
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
			this.portfolios.reduce(async (previousPromise, competence) => {
				const payload = {
					competence_id: competence.competence_id,
					semester: competence.semester,
					student_id: competence.student_id
				};
				this.forceReRender++;
				await previousPromise;
				const response = await Verify.postVerifyTransaction(payload);
				this.verify.push(response.data);
				this.show = true;
				return response;
			}, Promise.resolve());
		},
		getBadgeImgById(id) {
			return COMPETENCE[id].img;
		},
		clearVerify() {
			this.verify = [];
			this.show = false;
		}
	}
};
</script>