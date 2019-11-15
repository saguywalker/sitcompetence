export const widthSize = {
	data() {
		return {
			windowWidth: 0
		};
	},
	created() {
		window.addEventListener("resize", this.handleResize);
		this.handleResize();
	},
	destroyed() {
		window.removeEventListener("resize", this.handleResize);
	},
	methods: {
		handleResize() {
			this.windowWidth = window.innerWidth;
		}
	}
};