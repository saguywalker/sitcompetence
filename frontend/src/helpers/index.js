import CryptoJS from "crypto-js";
import SHA256 from "crypto-js/sha256";
import nacl from "tweetnacl";
import aesjs from "aes-js";
import Cookies from "js-cookie";
import { MONTH_NAMES } from "@/constants";

nacl.util = require("tweetnacl-util");

const encodeByteArrayToString = nacl.util.encodeBase64;
const decodeBase64ToByteArray = nacl.util.decodeBase64;

export function clearCookies() {
	const cookies = Cookies.get();

	Object.keys(cookies).forEach((item) => {
		Cookies.remove(item);
	});
}

export function clearLoginCookie() {
	Cookies.remove("x-session-token");
}

export function clearLoginState() {
	Cookies.remove("x-session-token", { path: "/" });
	localStorage.removeItem("user");
}

export function getCookie(name) {
	return Cookies.getJSON(name);
}

export function setCookie(name, value) {
	return Cookies.set(name, value);
}

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
export const getPlainTextToken = (cipher, key) => {
	const reb64 = CryptoJS.enc.Hex.parse(cipher);
	const bytes = reb64.toString(CryptoJS.enc.Base64);
	const decrypt = CryptoJS.AES.decrypt(bytes, key);
	const plain = decrypt.toString(CryptoJS.enc.Utf8);
	return plain;
};

// Encrypt
export const getCiphertext = (message, key) => {
	const base64 = CryptoJS.AES.encrypt(message, key).toString();
	const encode64 = CryptoJS.enc.Base64.parse(base64);
	const encodeHex = encode64.toString(CryptoJS.enc.Hex);
	return encodeHex;
};

export const getCiphertextBase64 = (message, key) => {
	const base64 = CryptoJS.AES.encrypt(message, key).toString();
	return base64;
};

export const getSecretBase64 = (sk) => {
	return btoa(sk);
};

export const getSHA256Message = (message) => SHA256(message).toString();

export const getLoginToken = () => {
	const loggedInData = localStorage.getItem("user");
	if (!loggedInData) {
		return null;
	}

	return JSON.parse(getPlainTextToken(loggedInData, process.env.VUE_APP_USER_DATA_KEY)).token;
};

export const getLoginUser = () => {
	const loggedInData = localStorage.getItem("user");
	if (!loggedInData) {
		return null;
	}
	return JSON.parse(getPlainTextToken(loggedInData, process.env.VUE_APP_USER_DATA_KEY));
};

export const getLoginUserRole = () => {
	const loggedInData = localStorage.getItem("user");
	if (!loggedInData) {
		return null;
	}

	return JSON.parse(getPlainTextToken(loggedInData, process.env.VUE_APP_USER_DATA_KEY)).group;
};

export const getSecretKey = () => {
	const sk = localStorage.getItem("sck");
	if (!sk) {
		return null;
	}

	return sk;
};

export const base64ToByteArray = (base64) => {
	const binaryString = window.atob(base64);
	const len = binaryString.length;
	const bytes = new Uint8Array(len);
	for (let i = 0; i < len; i++) {
		bytes[i] = binaryString.charCodeAt(i);
	}

	return bytes;
};

export const getED25519KeyPair = () => {
	const key = nacl.sign.keyPair();
	return {
		public: encodeByteArrayToString(key.publicKey),
		secret: encodeByteArrayToString(key.secretKey)
	};
};

export const parseHexString = (text) => {
	const result = [];
	let str = text;
	while (str.length >= 2) {
		result.push(parseInt(str.substring(0, 2), 16));
		str = str.substring(2, str.length);
	}

	return result;
};

export const getEncryptedHex = (text) => {
	const iv = parseHexString("00112233445566778899aabbccddeeff");
	const key = parseHexString(getSHA256Message(process.env.VUE_APP_AES));
	/* eslint-disable new-cap */
	const aesCbc = new aesjs.ModeOfOperation.cbc(key, iv);
	const encryptedBytes = aesCbc.encrypt(decodeBase64ToByteArray(text));
	const encryptedHex = aesjs.utils.hex.fromBytes(encryptedBytes);
	return encryptedHex;
};

export function getUnique(arr, comp) {
	const unique = arr.map((e) => e[comp])
		.map((e, i, final) => final.indexOf(e) === i && i)
		.filter((e) => arr[e]).map((e) => arr[e]);

	return unique;
}

export const isLoggedIn = () => {
	return !!Cookies.get("x-session-token") && !!localStorage.getItem("user");
};

export const filterBySemester = (badges, semester) => {
	return badges.filter((badge) => badge.semester === semester);
};

export function signMessage(msg, secret) {
	const signedMsg = nacl.sign.detached(decodeBase64ToByteArray(getSHA256Message(msg)), decodeBase64ToByteArray(secret));
	console.log(signedMsg);
	return encodeByteArrayToString(signedMsg);
}
