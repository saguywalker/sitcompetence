<template>
	<div class="activity-card">
		<div class="heading">
			<h2 class="title">
				Essential
			</h2>
		</div>
		<div class="competences">
			<h4 class="title">
				Competences
			</h4>
			<ul class="competences-list">
				<li class="item">
					Hi 10 pts.
				</li>
				<li class="item">
					Hi
				</li>
				<li class="item">
					Hi 20 pts.
				</li>
				<li class="item">
					Hi 10 pts.
				</li>
				<li class="item">
					Hi
				</li>
				<li class="item">
					Hi
				</li>
				<li class="item">
					Hi
				</li>
				<li class="item">
					Hi
				</li>
				<li class="item">
					Hi
				</li>
			</ul>
		</div>
		<ul class="info">
			<li class="info-item">
				<icon-calendar class="icon" />
				<span>11-12-61</span>
			</li>
			<li class="info-item">
				<icon-time class="icon" />
				<span>09:10</span>
			</li>
			<li class="info-item">
				<icon-location class="icon" />
				<span>CB2 KMUTT</span>
			</li>
		</ul>
		<p class="content">
			Lorem ipsum dolor, sit amet consectetur adipisicing elit. Fuga, perspiciatis, praesentium earum tenetur laudantium distinctio, non illum neque magnam odit ratione harum ducimus. Modi cupiditate magni minima, consectetur saepe distinctio.
		</p>
	</div>
</template>
<style lang="scss">
@import "@/styles/components/activity-card.scss";
</style>
<script>
import IconCalendar from "@/components/icons/IconCalendar.vue";
import IconLocation from "@/components/icons/IconLocation.vue";
import IconTime from "@/components/icons/IconTime.vue";

export default {
	components: {
		IconCalendar,
		IconLocation,
		IconTime
	}
};
</script>