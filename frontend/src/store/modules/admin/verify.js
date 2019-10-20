import { Verify } from "@/services";
import {
	UPDATE_VERIFY_HASH,
	UPDATE_VERIFY_DATA
} from "../../mutationTypes";

const state = {
	hashId: "",
	verifyData: []
};

const mutations = {
	[UPDATE_VERIFY_HASH](stateData, data) {
		stateData.hashId = data;
	},
	[UPDATE_VERIFY_DATA](stateData, data) {
		stateData.verifyData = [
			...stateData.verifyData,
			data
		];
	}
};

const actions = {
	updateHashId({ commit }, data) {
		commit(UPDATE_VERIFY_HASH, data);
	},
	async verifyTransaction({ commit }, data) {
		const	response = await Verify.postVerifyTransaction(data);

		if (response.status === 200) {
			commit(UPDATE_VERIFY_DATA, response.data);
		}

		return response.data;
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};