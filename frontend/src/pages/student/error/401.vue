<template>
	<div class="error-404">
		<div class="content">
			<h1 class="title">
				401 - Unauthorized
			</h1>
			<p class="description">
				Don't worry there is a way to go back home.
			</p>
			<router-link :to="{ name: 'admin' }">
				<b-button variant="primary">
					Back home
				</b-button>
			</router-link>
		</div>
	</div>
</template>
<style lang="scss">
@import "@/styles/pages/error404";
</style>