export const takeFileName = (path) => {
	if (!path) {
		return path;
	}

	// Take only last slash from full file
	const pathSplit = path.split("/");
	return pathSplit[pathSplit.length - 1];
};