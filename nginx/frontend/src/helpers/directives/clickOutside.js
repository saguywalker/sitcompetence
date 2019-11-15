import Vue from "vue";

let handleClickOutside;

/* eslint-disable import/prefer-default-export */
/**
 *
 * Click outside select element directive
 *
 * @example
 * v-click-outside={
 * 	handler: handle click outside function,
 * 	exclude: string[] of excluding element on click outside
 * }
 */
export const clickOutside = {
	bind(el, binding, vnode) {
		handleClickOutside = (event) => {
			event.stopPropagation();
			const { handler, exclude, sidebar } = binding.value;
			let clickOnExcludeEl = false;

			if (sidebar) {
				exclude.forEach((refName) => {
					if (!clickOnExcludeEl) {
						const excludeEl = vnode.context.$refs[refName];
						clickOnExcludeEl = excludeEl.contains(event.target);
					}
				});
				sidebar.forEach((refName) => {
					if (!clickOnExcludeEl) {
						const sidebarBtnToggle = vnode.context.$parent.$refs.navbar.$refs[refName];
						clickOnExcludeEl = sidebarBtnToggle.contains(event.target);
					}
				});
			} else {
				exclude.forEach((refName) => {
					if (!clickOnExcludeEl) {
						const excludeEl = vnode.context.$refs[refName];
						clickOnExcludeEl = excludeEl.contains(event.target);
					}
				});
			}

			if (!el.contains(event.target) && !clickOnExcludeEl) {
				vnode.context[handler]();
			}
		};

		document.addEventListener("click", handleClickOutside);
	},
	unbind() {
		document.removeEventListener("click", handleClickOutside);
	}
};

Vue.directive("click-outside", clickOutside);