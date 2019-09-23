export const takeFileName = (path) => {
	if (!path) {
		return path;
	}

	// Take only last slash from full file
	const pathSplit = path.split("/");
	return pathSplit[pathSplit.length - 1];
};

export const base64ToHex = (str) => {
	const binary = window.atob(str);

	const hex = [];
	for (let i = 0; i < binary.length; i++) {
		let tmp = binary.charCodeAt(i).toString(16);
		if (tmp.length === 1) {
			tmp = `0${tmp}`;
		}

		hex[i] = tmp;
	}

	return hex.join("");
};

export const hexToBase64 = (str) => {
	return btoa(String.fromCharCode.apply(null, str.replace(/\r|\n/g, "").replace(/([\da-fA-F]{2}) ?/g, "0x$1 ").replace(/ +$/, "").split(" ")));
};

export const getSemester = () => {
	const today = new Date();
	const month = parseInt(today.getDate(), 10);
	const year = parseInt(today.getFullYear(), 10);

	if (month >= 8 && month <= 12) {
		return 10000 + year;
	}	else if (month === 1) {
		return 10000 + (year + 1);
	}

	return 20000 + year;
};