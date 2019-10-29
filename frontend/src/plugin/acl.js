import Vue from "vue";
import { AclInstaller, AclCreate, AclRule } from "vue-acl";
import { getLoginUserRole } from "@/helpers";
import router from "@/router";

Vue.use(AclInstaller);

export default new AclCreate({
	notfound: {
		path: "unknown"
	},
	router,
	acceptLocalRules: true,
	globalRules: {
		isAdmin: new AclRule("inst_group").generate(),
		isStudent: new AclRule("st_group").generate(),
		isPublic: new AclRule("*").generate()
	},
	middleware: async (acl) => {
		const userRole = getLoginUserRole();
		await acl.change(userRole);
	}
});