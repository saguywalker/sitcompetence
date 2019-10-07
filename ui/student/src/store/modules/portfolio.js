import { Base } from "@/services";
import {
	LOAD_PORTFOLIO,
	UPDATE_PORTFOLIO,
	UPDATE_PORTFOLIO_LINK
} from "../mutationTypes";

const state = {
	portfolios: [],
	link: ""
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
	}
};

const actions = {
	async loadPortfolio({ commit }, id) {
		const	response = await Base.getBadgesByStudentId(id);
		if (response.status === 200) {
			commit(LOAD_PORTFOLIO, response.data);
		}

		return response;
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};