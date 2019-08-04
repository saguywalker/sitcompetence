import React from "react";

function BaseImagePic({imgSrc, size, round}) {
	function imageSizeClass(size) {
		switch (size) {
		case "32":
			return "is-32x32";
		case "64":
			return "is-64x64";
		default:
			return "is-128x128";
		}
	}

	function imageRoundClass(round) {
		if (round) {
			return "is-rounded";
		}
	}

	return (
		<figure className={`image ${imageSizeClass(size)}`}>
			<img
				className={`${imageRoundClass(round)}`}
				src={imgSrc || "https://bulma.io/images/placeholders/128x128.png"}
				alt="128x128"
			/>
		</figure>
	);
}

export default BaseImagePic;