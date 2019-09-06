/* eslint-disable import/prefer-default-export */
export const GIVE_BADGE_STEP = Object.freeze({
	main: Object.freeze({
		next: Object.freeze({
			name: "Badge Selection",
			link: "give-badge-selection"
		}),
		step: "main"
	}),
	selection: Object.freeze({
		next: Object.freeze({
			name: "Confirmation",
			link: "give-badge-confirmation"
		}),
		back: Object.freeze({
			name: "Select Student",
			link: "give-badge"
		}),
		step: "selection"
	}),
	confirmation: Object.freeze({
		next: Object.freeze({
			name: "Submit",
			link: "give-badge-success"
		}),
		back: Object.freeze({
			name: "Badge Selection",
			link: "give-badge-selection"
		}),
		step: "confirmation"
	}),
	success: Object.freeze({
		step: "success"
	})
});