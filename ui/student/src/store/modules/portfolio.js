import {
	LOAD_PORTFOLIO,
	UPDATE_PORTFOLIO_LINK
} from "../mutationTypes";

const state = {
	portfolio: {},
	link: ""
};

const mutations = {
	[LOAD_PORTFOLIO](stateData, data) {
		stateData.portfolio = {
			...data
		};
	},
	[UPDATE_PORTFOLIO_LINK](stateData, data) {
		stateData.link = data;
	}
};

const actions = {
	loadPortfolio({ commit }) {
		const response = "will do";
		commit(LOAD_PORTFOLIO, response);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};