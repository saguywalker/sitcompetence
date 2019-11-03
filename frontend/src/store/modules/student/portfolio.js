import { Base, Verify, Portfolio } from "@/services";
import {
	LOAD_PORTFOLIO,
	UPDATE_PORTFOLIO,
	UPDATE_PORTFOLIO_LINK,
	UPDATE_VERIFY_DATA,
	CLEAR_VERIFY_DATA
} from "../../mutationTypes";

const state = {
	portfolios: [],
	link: "",
	verify: [],
	show: false
};

const mutations = {
	[LOAD_PORTFOLIO](stateData, data) {
		stateData.portfolios = data;
	},
	[UPDATE_PORTFOLIO](stateData, data) {
		stateData.portfolios = data;
	},
	[UPDATE_PORTFOLIO_LINK](stateData, data) {
		stateData.link = data;
	},
	[UPDATE_VERIFY_DATA](stateData, data) {
		stateData.verify.push(data);
		stateData.show = true;
	},
	[CLEAR_VERIFY_DATA](stateData) {
		stateData.verify = [];
		stateData.show = false;
	}
};

const actions = {
	async loadPortfolio({ commit }, id) {
		const	response = await Base.getBadgesByStudentId(id);
		if (response.status === 200) {
			commit(LOAD_PORTFOLIO, response.data);
		}

		return response;
	},
	async loadShareLink({ commit }) {
		const	response = await Portfolio.getShareLink();
		if (response.status === 200) {
			commit(UPDATE_PORTFOLIO_LINK, response.data);
		}

		return response;
	},
	async verifyTransaction({ commit }, data) {
		const	response = await Verify.postVerifyTransaction(data);

		if (response.status === 200) {
			commit(UPDATE_VERIFY_DATA, response.data);
		}

		return response.data;
	},
	clearVerify({ commit }) {
		commit(CLEAR_VERIFY_DATA);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};