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

export const CREATE_ACTIVITY_STEP = Object.freeze({
	detail: Object.freeze({
		next: Object.freeze({
			name: "Select competence",
			link: "create-activity-competence"
		}),
		back: Object.freeze({
			name: "Activity main",
			link: "activity"
		}),
		step: "detail"
	}),
	competence: Object.freeze({
		next: Object.freeze({
			name: "Confirmation",
			link: "create-activity-summary"
		}),
		back: Object.freeze({
			name: "Activity detail",
			link: "create-activity-detail"
		}),
		step: "select"
	}),
	summary: Object.freeze({
		next: Object.freeze({
			name: "Submit",
			link: "create-activity-success"
		}),
		back: Object.freeze({
			name: "Select Competence",
			link: "create-activity-competence"
		}),
		step: "summary"
	}),
	success: Object.freeze({
		step: "success"
	})
});