<template>
	<div
		:class="[
			'input-text',
			inputSize,
			{focused : isFilled},
			{error: isHasErrorMessage}
		]"
	>
		<label>
			<input
				v-model="selfValue"
				:type="type"
				:name="name"
				:required="required"
				:autocomplete="autocomplete"
				:maxlength="maxlength"
				:placeHolder="placeHolder"
				class="input-text-input"
				@keydown="hideKeyboard"
				@focus="onFocus"
				@blur="onBlur"
				@change="checkOnChange"
			>
			<span
				v-show="label"
				class="input-text-label"
			>
				{{ label }}
			</span>
			<span
				v-if="isHasErrorMessage"
				class="input-text-error"
			>
				{{ isHasErrorMessage }}
			</span>
		</label>
	</div>
</template>

<style lang="scss">
@import "@/styles/components/input-text.scss";
</style>

<script>
export default {
	props: {
		value: {
			type: [String, Number],
			default: ""
		},
		id: {
			type: String,
			default: ""
		},
		type: {
			type: String,
			default: "text"
		},
		name: {
			type: String,
			default: ""
		},
		required: {
			type: Boolean,
			default: false
		},
		label: {
			type: String,
			default: ""
		},
		autocomplete: {
			type: String,
			default: ""
		},
		maxlength: {
			type: String,
			default: ""
		},
		max: {
			type: [Number],
			default: 99999
		},
		errorMessage: {
			type: String,
			default: ""
		},
		placeHolder: {
			type: [String],
			default: ""
		},
		size: {
			type: String,
			default: ""
		}
	},
	data() {
		return {
			selfValue: "",
			focused: false
		};
	},
	computed: {
		isError() {
			return this.errorMessage;
		},
		isFilled() {
			return this.selfValue || this.focused;
		},
		isHasErrorMessage() {
			return this.errorMessage ? this.errorMessage : false;
		},
		inputSize() {
			if (this.size === "lg") {
				return "lg";
			}

			return "";
		}
	},
	watch: {
		value(newValue) {
			if (this.selfValue !== newValue) {
				this.selfValue = newValue;
			}
		},
		selfValue(newValue) {
			this.$emit("input", newValue);
		}
	},
	created() {
		this.selfValue = this.value;
	},
	methods: {
		checkOnChange(e) {
			// autocomplete on iPhone doesn`t fire checkOnKeyup()
			// this function adds additional check of input value
			const value = e.target.value;
			this.$emit("input", value);
		},
		checkOnKeyup(e) {
			let value = e.target.value;
			if (e.key === "Backspace") {
				value = e.target.value;
			}
			this.$emit("input", value);
		},
		onFocus() {
			if (this.name !== "referenceCode") {
				this.focused = true;
			}

			if (this.errorMessage) {
				this.$emit("removeMessageError", this.name);
			}
		},
		onBlur(e) {
			if (!e.target.value) {
				this.focused = false;
			}

			this.$emit("blur", e);
		},
		hideKeyboard(e) {
			if (e.keyCode === 13) { // when press enter
				e.preventDefault();
				e.target.blur();
			}
		}
	}
};
</script>