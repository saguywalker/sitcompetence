<template>
	<b-container>
		<code>
			{
			token: {{ token }}
			}
		</code>
	</b-container>
</template>
<script>
import { mapState } from "vuex";

export default {
	computed: mapState("base", ["token"])
};
</script>