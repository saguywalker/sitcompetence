import React from "react";
import BaseImagePic from "components/BaseImagePic";

function GiveBadgeStudentDescription() {
	return (
		<div className="gb-student-description">
			<BaseImagePic />
			<div className="gb-student-description-content">
				<h1 className="title is-5">
					Nongtiny
				</h1>
				<p className="subtitle is-6">
					Lorem ipsum dolor sit amet consectetur adipisicing elit. Illo deleniti, ipsum blanditiis accusamus facere facilis eveniet ad id unde fugiat? Vero quae temporibus quaerat nulla inventore minus ex dolorem omnis!
				</p>
			</div>
		</div>
	);
}

export default GiveBadgeStudentDescription;