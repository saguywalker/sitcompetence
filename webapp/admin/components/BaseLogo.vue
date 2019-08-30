<template>
	<svg
		:width="logoSizeWidth"
		:height="logoSizeHeight"
		viewBox="0 0 76 113"
		fill="none"
		xmlns="http://www.w3.org/2000/svg"
	>
		<path
			d="M0.953171 32.2434L58.6619 50.2805V112.044L0.953171 93.8245V32.2434Z"
			fill="#396290"
		/>
		<path
			d="M20.3397 25.9315L74.8924 43.0515V101.674L20.3397 84.3814V25.9315Z"
			fill="#1E4571"
		/>
		<path
			d="M75.1179 43.5147L58.6619 50.0711V112.044L75.1179 101.674V43.5147Z"
			fill="#243E5C"
		/>
		<path
			d="M20.3397 25.9315L0.95318 32.2626V94.2352L20.3397 84.7674V25.9315Z"
			fill="#2C4C71"
		/>
		<path
			d="M75.5687 43.2892L20.1143 25.9316L0.953164 32.2434L58.8873 50.5028L75.5687 43.2892Z"
			fill="#243E5C"
		/>
		<g filter="url(#filter0_d)">
			<path
				d="M17.9959 67.4199L39.3932 72.8682V96.264L17.9959 90.1747V67.4199Z"
				fill="white"
			/>
			<path
				d="M39.3932 73.284L47.2049 69.7773V90.8969L39.3932 96.264V73.284Z"
				fill="#FBFBFB"
			/>
			<path
				d="M25.8076 64.5077L47.2049 69.7773L39.6706 73.6633L17.9959 67.4283L25.8076 64.5077Z"
				fill="#FBFBFB"
			/>
			<path
				d="M27.5775 74.2669L26.8596 71.336L29.9251 72.3709L27.5775 74.2669Z"
				fill="#072139"
			/>
			<path
				d="M27.8455 90.4901L26.8266 88.6221L28.8644 89.4712L27.8455 90.4901Z"
				fill="#072139"
			/>
			<path
				d="M28.9785 89.5013L26.7243 88.7324V73.2707L28.9785 73.9026V89.5013Z"
				fill="#072139"
			/>
		</g>
		<path
			d="M21.0527 50.073L35.9968 54.6921V74.5271L21.0527 69.3646V50.073Z"
			fill="#FFE0BD"
		/>
		<path
			d="M35.9968 55.2356L43.2991 51.1599V69.908L35.9968 74.5271V55.2356Z"
			fill="#FED9B0"
		/>
		<path
			d="M28.6946 46.6767L43.2991 51.4316L35.9968 56.5262L21.0527 50.0731L28.6946 46.6767Z"
			fill="#FED9B0"
		/>
		<path
			d="M36.1921 53.8318L43.4057 51.63V61.7584L36.1921 64.2538V53.8318Z"
			fill="black"
			fill-opacity="0.75"
		/>
		<path
			d="M36.1921 53.0362L43.4057 51.63V58.0985L37.1758 60.1961L36.1921 53.0362Z"
			fill="black"
			fill-opacity="0.91"
		/>
		<path
			d="M29.118 46.2198L43.4057 51.5946L38.8741 54.7859L29.118 53.7781L20.4124 49.5791L29.118 46.2198Z"
			fill="black"
		/>
		<path
			d="M28.097 45.3181L43.4057 49.462L39.3682 53.4333L29.5892 52.3973L20.8632 48.0807L28.097 45.3181Z"
			fill="black"
		/>
		<path
			d="M26.7276 45.3181L42.504 50.221L37.9725 53.4333L28.2163 52.4189L19.5107 48.1922L26.7276 45.3181Z"
			fill="black"
		/>
		<path
			d="M27.8335 46.6706L43.8565 51.1538L39.2694 54.3351L29.3935 53.3305L20.4124 48.8848L27.8335 46.6706Z"
			fill="black"
		/>
		<path
			d="M36.0239 64.2538L39.348 54.4966L27.626 52.9825L32.8747 58.3658L36.0239 64.2538Z"
			fill="black"
		/>
		<path
			d="M20.4124 61.9995V49.3757L29.8802 53.8719L23.7937 56.1199L20.4124 61.9995Z"
			fill="black"
		/>
		<defs>
			<filter
				id="filter0_d"
				x="17.9959"
				y="64.5077"
				width="37.209"
				height="39.7563"
				filterUnits="userSpaceOnUse"
				color-interpolation-filters="sRGB"
			>
				<feFlood
					flood-opacity="0"
					result="BackgroundImageFix"
				/>
				<feColorMatrix
					in="SourceAlpha"
					type="matrix"
					values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
				/>
				<feOffset
					dx="4"
					dy="4"
				/>
				<feGaussianBlur stdDeviation="2" />
				<feColorMatrix
					type="matrix"
					values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.18 0"
				/>
				<feBlend
					mode="normal"
					in2="BackgroundImageFix"
					result="effect1_dropShadow"
				/>
				<feBlend
					mode="normal"
					in="SourceGraphic"
					in2="effect1_dropShadow"
					result="shape"
				/>
			</filter>
		</defs>
	</svg>
</template>
<script>
import Vue from "vue";

export default Vue.extend({
	props: {
		size: {
			type: Number,
			default: 76
		}
	},
	computed: {
		logoSizeWidth() {
			return `${this.size}`;
		},
		logoSizeHeight() {
			return `${Math.ceil(this.size * 1.48)}`;
		}
	}
});
</script>