<template>
	<div class="portfolio">
		<div class="page-header">
			<div class="wrapper">
				<h2 class="title">
					Portfolio
				</h2>
			</div>
		</div>
		<section class="wrapper">
			<div class="portfolio-actions">
				<div class="right">
					<b-button
						id="popover-share-port"
						ref="shareButton"
						:disabled="!hasCompetence"
						variant="primary"
						class="item"
						@click="sharePortfolio"
					>
						Share
					</b-button>
				</div>
				<b-popover
					ref="portPopover"
					target="popover-share-port"
					triggers="manual"
					:show.sync="popoverShow"
					@hidden="onHidden"
				>
					<template v-slot:title>
						<button
							:style="{
								position: 'absolute',
								top: '7px',
								right: '10px'
							}"
							class="close-popover"
							aria-label="Close"
							@click="onClose"
						>
							<icon-cross-circle />
						</button>
						Share portfolio link
					</template>
					<b-button-group class="my-2">
						<b-input
							id="shareUrl"
							ref="shareUrlInput"
							size="sm"
							:value="`http://localhost:8080/viewProfile/${link}`"
						/>
						<b-button
							ref="copyBtn"
							class="copy-btn"
							size="sm"
							variant="primary"
							@click="onCopyUrl"
						>
							Copy
						</b-button>
					</b-button-group>
				</b-popover>
			</div>
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
				<template v-if="hasCompetence">
					<div class="portfolio-content">
						<b-row class="portfolio-container">
							<b-col
								v-for="(com, index) in statePortfolios.result"
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
						<h1>No competence record.</h1>
					</div>
				</template>
			</div>
		</section>
	</div>
</template>
<style lang="scss">
@import "@/styles/portfolio.scss";
</style>
<script>
import IconCheckCircle from "@/components/icons/IconCheckCircle.vue";
import IconTimeCircle from "@/components/icons/IconTimeCircle.vue";
import IconCrossCircle from "@/components/icons/IconCrossCircle.vue";
import loading from "@/plugin/loading";
import { widthSize } from "@/helpers/mixins";
import { getLoginUser } from "@/helpers";
import { COMPETENCE } from "@/constants";
import { mapState } from "vuex";

export default {
	components: {
		IconCheckCircle,
		IconTimeCircle,
		IconCrossCircle
	},
	mixins: [widthSize],
	data() {
		return {
			popoverShow: false,
			forceReRender: 0,
			hasCompetence: null
		};
	},
	computed: {
		...mapState("portfolio", {
			statePortfolios: "portfolios",
			verify: "verify",
			show: "show",
			link: "link"
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
			this.hasCompetence = this.statePortfolios.result.length !== 0;
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
		onOpen() {
			this.$refs.portPopover.$emit("open");
		},
		onClose() {
			this.popoverShow = false;
		},
		onHidden() {
			// Called just after the popover has finished hiding
			// Bring focus back to the button
			this.focusRef(this.$refs.shareButton);
		},
		onCopyUrl() {
			const testingCodeToCopy = document.querySelector("#shareUrl");
			testingCodeToCopy.select();

			try {
				document.execCommand("copy");

				this.$bvToast.toast("Copy successful", {
					title: "Share portfolio link",
					variant: "success",
					autoHideDelay: 1500
				});
			} catch (err) {
				this.$bvToast.toast("Oops, cannot copy link", {
					title: "Share portfolio link",
					variant: "danger",
					autoHideDelay: 1500
				});
			}

			window.getSelection().removeAllRanges();
		},
		focusRef(ref) {
			// Some references may be a component, functional component, or plain element
			// This handles that check before focusing, assuming a `focus()` method exists
			// We do this in a double `$nextTick()` to ensure components have
			// updated & popover positioned first
			this.$nextTick(() => {
				this.$nextTick(() => {
					(ref.$el || ref).focus();
				});
			});
		},
		async sharePortfolio() {
			loading.start();

			try {
				await this.$store.dispatch("portfolio/loadShareLink");
				this.onOpen();
			} catch (err) {
				this.$bvToast.toast(`Get share portfolio link error problem: ${err.message}`, {
					title: "Share portfolio link",
					variant: "danger",
					autoHideDelay: 1500
				});
			} finally {
				loading.stop();
			}
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
			this.statePortfolios.result.reduce(async (previousPromise, competence) => {
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