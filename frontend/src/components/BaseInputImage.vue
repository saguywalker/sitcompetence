<template>
	<div class="input-image">
		<b-form-group
			:label="label"
			label-for="input-1"
		>
			<base-image
				:src="previewImageSrc"
				:size="size"
				class="my-3"
			/>
			<b-form-file
				id="input-1"
				:value="value"
				accept="image/jpg, image/jpeg, image/png"
				plain
				@input="handleUpload"
			/>
		</b-form-group>
	</div>
</template>
<script>
export default {
	props: {
		label: {
			type: String,
			default: ""
		},
		value: {
			type: File,
			default: null
		},
		size: {
			type: String,
			default: "100"
		}
	},
	data() {
		return {
			previewImageSrc: "https://via.placeholder.com/100"
		};
	},
	methods: {
		handleUpload(e) {
			// Validate file
			if (!e) {
				return;
			}
			this.previewImageSrc = URL.createObjectURL(e);
			this.$emit("input", e);
		}
	}
};
</script>