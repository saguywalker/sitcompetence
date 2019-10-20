<template>
	<div class="box">
		<h1>If you see this page verify data is success</h1>
		<h2>Response Data:</h2>
		<br>
		{{ verifyData }}
	</div>
</template>
<script>
import { mapState } from "vuex";

export default {
	computed: {
		...mapState("verify", ["verifyData"])
	}
};
</script>