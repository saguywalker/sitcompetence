import React from "react";

function MainContent() {
	const numbers = [1, 2, 3, 4, 5];
	const listItems = numbers.map((number, key) =>
		<div
			key={key}
		>
			Lorem ipsum dolor sit amet consectetur adipisicing elit. Praesentium inventore provident exercitationem, dolores repudiandae fuga totam tenetur quam ut illum amet architecto quaerat soluta fugiat ullam odio eveniet? Cumque, incidunt.
			Natus quidem voluptates minima expedita, magnam officia nostrum necessitatibus officiis sed adipisci. Repellat facilis, ipsa eligendi ea nisi natus tenetur. Inventore sunt esse consequuntur optio sed tenetur, praesentium molestias maxime.
			Ad perspiciatis autem, eos vel eligendi numquam et sapiente aliquam culpa ex ut sequi facere deleniti ipsam ipsa consequatur, quasi consectetur, est voluptates saepe debitis corporis magni amet iusto? Officiis.
			Lorem ipsum dolor sit amet consectetur adipisicing elit. Iste cupiditate minus sunt molestias sed necessitatibus quis, consequatur autem expedita consectetur ullam tenetur eum nobis ab totam facilis corrupti. Veritatis, recusandae.
			Ut asperiores rerum quis molestiae nobis cupiditate, at cumque, nulla molestias laudantium odio dignissimos deserunt maxime laborum dolores ad suscipit quas. Quae rerum repudiandae ipsa. Ex magnam ipsa dicta nam!
			Magnam explicabo odio iusto! Fuga quam debitis voluptatum, beatae quae dolores, harum similique in placeat minima explicabo nihil ipsam nobis voluptates, reprehenderit delectus unde repellendus impedit culpa est mollitia nemo?
			Ducimus inventore doloremque sed sequi ut illo mollitia, iste maiores fugit, excepturi sunt esse voluptatem laboriosam earum doloribus! Incidunt eveniet quis facere, quia quasi dolores impedit ut in quisquam nihil!
			Perferendis laudantium eveniet deserunt quos nulla beatae odio at numquam quo cum recusandae, fugiat accusamus minus perspiciatis exercitationem ullam quaerat? Praesentium explicabo hic labore dolorem ipsa in velit soluta recusandae.
			Culpa, impedit numquam ipsum saepe autem in eos, molestiae ut esse odit aspernatur nisi ratione neque odio? Cum doloribus et officia exercitationem in, odio labore minus, praesentium id hic quia.
			Maxime quaerat in, rem vel maiores qui repudiandae, quidem voluptatem aliquam laboriosam placeat! Enim nam hic iure quidem accusamus, expedita, doloribus laboriosam dolorem praesentium suscipit ut in qui vel rem!
			Nobis beatae blanditiis repellendus itaque dicta fuga, aut dolorum at necessitatibus id deserunt mollitia magnam officiis impedit quas doloremque! Odit culpa ex nihil voluptate voluptatibus perferendis dolore natus at ut!
			Itaque consectetur ipsa atque quasi quos eos, sapiente iusto recusandae officiis asperiores qui rerum delectus soluta tempora distinctio corrupti, saepe consequatur mollitia! Veniam maiores odio quae a cum? Doloribus, facere.
			Ducimus dolorum similique corrupti voluptatem temporibus, suscipit repudiandae provident quibusdam esse laboriosam minima dignissimos, quisquam libero ea, voluptatum necessitatibus inventore sit dicta quidem cum commodi nobis repellat. Nostrum, odio magni.
		</div>
	);
	return (
		<div className="columns is-mobile">
			<div className="column test is-1"/>
			<div className="column test">
				<div className="container">
					<div className="section">
						{ listItems }
					</div>
				</div>
			</div>
		</div>
	);
}

export default MainContent;