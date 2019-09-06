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
			const { handler, exclude } = binding.value;
			let clickOnExcludeEl = false;

			exclude.forEach((refName) => {
				if (!clickOnExcludeEl) {
					const excludeEl = vnode.context.$refs[refName];
					clickOnExcludeEl = excludeEl.contains(event.target);
				}
			});

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