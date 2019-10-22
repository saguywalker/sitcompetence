import CryptoJS from "crypto-js";
import { MONTH_NAMES } from "@/constants";

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

export const getMonthNameByDateFormat = (format) => {
	const date = new Date(format);

	return MONTH_NAMES[date.getMonth()];
};

export const getYearByDateFormat = (format) => {
	const date = new Date(format);

	return date.getFullYear();
};

export const getEditDateFormat = (format) => {
	const date = new Date(format);
	const day = date.getDate();
	const computedDay = day > 9 ? day : `0${day}`;
	const month = date.getMonth();
	const computedMonth = month > 9 ? month : `0${month}`;
	const year = date.getFullYear();

	return `${year}-${computedMonth}-${computedDay}`;
};

// Decrypt
export const getPlainTextToken = (cipher) => {
	const reb64 = CryptoJS.enc.Hex.parse(cipher);
	const bytes = reb64.toString(CryptoJS.enc.Base64);
	const decrypt = CryptoJS.AES.decrypt(bytes, process.env.VUE_APP_USER_DATA_KEY);
	const plain = decrypt.toString(CryptoJS.enc.Utf8);
	return plain;
};

// Encrypt
export const getCiphertext = (message) => {
	const b64 = CryptoJS.AES.encrypt(message, process.env.VUE_APP_USER_DATA_KEY).toString();
	const e64 = CryptoJS.enc.Base64.parse(b64);
	const eHex = e64.toString(CryptoJS.enc.Hex);
	return eHex;
};

export const getLoginToken = () => {
	const loggedInData = localStorage.getItem("user");
	if (!loggedInData) {
		return null;
	}

	return JSON.parse(getPlainTextToken(loggedInData)).token;
};

export const getLoginUser = () => {
	const loggedInData = localStorage.getItem("user");
	if (!loggedInData) {
		return null;
	}

	return JSON.parse(getPlainTextToken(localStorage.getItem("user"))).user;
};

export const getLoginUserRole = () => {
	const loggedInData = localStorage.getItem("user");
	if (!loggedInData) {
		return null;
	}

	return JSON.parse(getPlainTextToken(localStorage.getItem("user"))).user.group;
};