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
		isAdmin: new AclRule("admin").generate(),
		isStudent: new AclRule("st_group").generate(),
		isPublic: new AclRule("*").generate()
	},
	middleware: async (acl) => {
		const userRole = getLoginUserRole();
		if (userRole === "inst_group") {
			await acl.change("admin");
		} else {
			await acl.change(userRole);
		}
	}
});