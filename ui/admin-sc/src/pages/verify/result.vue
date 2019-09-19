<template>
	<div class="box">
		<h1>I don't know what you will response</h1>
		<h1>But the verify is complete if you see this page.</h1>
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