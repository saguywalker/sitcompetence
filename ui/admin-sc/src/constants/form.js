export const CREATE_ACTIVITY_FORM = Object.freeze([
	{
		label: "Activity name",
		type: "text",
		model: "name"
	},
	{
		label: "Description",
		type: "textarea",
		model: "description"
	},
	{
		label: "Image",
		type: "file",
		model: "img"
	},
	{
		label: "Activity date",
		type: "date",
		model: "activityDate"
	}
]);