<template>
	<div class="portfolio">
		<div class="page-header">
			<div class="wrapper">
				<h2 class="title">
					Portfolio
				</h2>
			</div>
		</div>
		{{ statePortfolios }}
		<div class="portfolio-actions">
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
							<h5
								v-for="(name, index) in splitedFullname"
								:key="`${name}${index}`"
							>
								{{ name }}
							</h5>
						</div>
					</div>
					<p class="profile-description">
						Lorem ipsum dolor, sit amet consectetur adipisicing elit. Voluptatum quidem hic sequi distinctio consectetur pariatur nostrum nesciunt, culpa quis quaerat saepe laborum officia! Impedit non suscipit consequuntur nostrum rerum temporibus?
					</p>
				</aside>
				<div class="portfolio-content">
					<b-row>
						<b-col
							v-for="(com, index) in portfolios"
							:key="`${com.competence_id}${forceReRender}`"
							cols="6"
							class="competence-wrapper"
						>
							<base-image
								:size="resizeImage"
								class="sitcom-badge"
							/>
							<p class="name">
								{{ com.competence_name }}
							</p>
							<transition
								:key="`${com.competence_id}${index}`"
								name="fade"
								mode="out-in"
							>
								<p
									v-if="com.verify_show"
									class="verify-status"
								>
									<span class="icon">
										<icon-check-circle v-if="com.verify_status" />
										<icon-time-circle v-else />
									</span>
									Verified
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
							v-if="portfolios[0].verify_show"
							size="sm"
							class="action-item"
							variant="outline-primary"
							@click="clearVerified"
						>
							Clear Verified
						</b-button>
					</div>
				</div>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/portfolio.scss";
</style>
<script>
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";
import IconTimeCircle from "@/components/icons/IconTimeCircle.vue";
import { widthSize } from "@/helpers/mixins";
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
			portfolios: [],
			verify: false
		};
	},
	computed: {
		...mapState("portfolio", {
			statePortfolios: "portfolios"
		}),
		userId() {
			return "59130500210";
		},
		splitedFullname() {
			return this.fullName.split(" ");
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
			await this.$store.dispatch("portfolio/loadPortfolio", this.userId);
			this.portfolios = this.statePortfolios;
		} catch (err) {
			this.$Progress.fail();
			this.$bvToast.toast(`Fetching data problem: ${err.message}`, {
				title: "Fetching competences error",
				variant: "danger",
				autoHideDelay: 1500
			});
		}
	},
	methods: {
		methodThatReturnsAPromise(id) {
			/* eslint-disable */
			return new Promise((resolve, reject) => {
				setTimeout(() => {
					console.log(`Resolve! ${id}`);
					this.setVerifyStatusById(id, false);
					++this.forceReRender;
					resolve();
				}, Math.random() * 1000);
			});
		},
		clearVerified() {
			this.portfolio.map((p) => {
				delete p.verify_show;
				delete p.verify_status;

				return p;
			});

			++this.forceReRender;
		},
		testVerify() {
			this.portfolio.reduce(async (previousPromise, port, index) => {
				console.log(`Loop: ${index}`);
				if (port.verify_show) {
					port.verify_status = null;
					port.verify_show = false;
				}
				await previousPromise;
				return this.methodThatReturnsAPromise(index);
			}, Promise.resolve());


			// try {

			// }
		},
		setVerifyStatusById(id, status) {
			this.portfolio[id].verify_status = status;
			this.portfolio[id].verify_show = true;
		}
	}
};
</script>