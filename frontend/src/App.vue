<template>
	<div id="app">
		<router-view />
	</div>
</template>
<script>
export default {
	watch: {
		/* eslint-disable-next-line  */
		$route(to, from) {
			document.title = to.meta.title || "SIT-Competence";
		}
	}
};
</script>