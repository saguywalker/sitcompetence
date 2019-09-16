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
		label: "Open registration",
		type: "date",
		type2: "time",
		model: "openRegistDate",
		model2: "openRegistTime"
	},
	{
		label: "Close registration",
		type: "date",
		type2: "time",
		model: "closeRegistDate",
		model2: "closeRegistTime"
	},
	{
		label: "Activity start",
		type: "date",
		type2: "time",
		model: "actStartDate",
		model2: "actStartTime"
	},
	{
		label: "Activity end",
		type: "date",
		type2: "time",
		model: "actEndDate",
		model2: "actEndTime"
	}
]);